package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/lethain/dfsmr/dfsmr"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
)

var (
	addr = flag.String("addr", ":5003", "Port to bind on")
)

type server struct {
	listeners      []chan *pb.ChangesReply
	changesMutex   *sync.RWMutex
	machines       []*pb.DefineRequest
	machinesMutex  *sync.RWMutex
	instances      []*pb.TaskMessage
	instancesMutex *sync.RWMutex
}

func MakeServer() *server {
	srv := &server{}
	srv.machinesMutex = &sync.RWMutex{}
	srv.changesMutex = &sync.RWMutex{}
	srv.instancesMutex = &sync.RWMutex{}
	return srv
}

func (s *server) RegisterMachine(newMachine *pb.DefineRequest) error {
	s.machinesMutex.Lock()
	s.machines = append(s.machines, newMachine)
	s.machinesMutex.Unlock()
	return nil
}

func (s *server) Machines(ctx context.Context, mr *pb.MachinesRequest) (*pb.MachinesReply, error) {
	s.machinesMutex.RLock()
	ms := s.machines[:]
	s.machinesMutex.RUnlock()
	return &pb.MachinesReply{Machines: ms}, nil
}

func (s *server) Instances(ctx context.Context, ir *pb.InstancesRequest) (*pb.InstancesReply, error) {
	s.instancesMutex.RLock()
	is := s.instances[:]
	s.instancesMutex.RUnlock()
	return &pb.InstancesReply{Instances: is}, nil
}

func (s *server) changeListener() chan *pb.ChangesReply {
	c := make(chan *pb.ChangesReply)
	s.changesMutex.Lock()
	s.listeners = append(s.listeners, c)
	s.changesMutex.Unlock()
	return c
}

func (s *server) closeChangeListener(c chan *pb.ChangesReply) {
	s.changesMutex.Lock()
	for i, cl := range s.listeners {
		if cl == c {
			s.listeners[i] = s.listeners[len(s.listeners)-1]
			s.listeners = s.listeners[:len(s.listeners)-1]
			break
		}
	}
	s.changesMutex.Unlock()
	close(c)
}

func determineClient(ctx context.Context, supplied string) string {
	if supplied != "" {
		return supplied
	}
	p, ok := peer.FromContext(ctx)
	if ok {
		return p.Addr.Network() + "://" + p.Addr.String()
	}
	uid, err := uuid.NewV4()
	if err != nil {
		return "invalid-client"
	}
	return uid.String()
}

func (s *server) record(ctx context.Context, op string, cmd interface{}) error {
	client := determineClient(ctx, "")
	cmdStr := fmt.Sprintf("%s %+v", op, cmd)
	cr := &pb.ChangesReply{Command: cmdStr, Client: client}
	s.changesMutex.RLock()
	defer s.changesMutex.RUnlock()
	for _, lc := range s.listeners {
		lc <- cr
	}
	return nil
}

func (s *server) getMachine(id string) *pb.DefineRequest {
	s.machinesMutex.RLock()
	defer s.machinesMutex.RUnlock()
	for _, m := range s.machines {
		if id == m.Id {
			return m
		}
	}
	return nil
}

// Determine the starting node for a given state machine
func startNode(dr *pb.DefineRequest) (string, error) {
	for _, node := range dr.Nodes {
		if node.Start == true {
			return node.Id, nil
		}
	}
	return "", fmt.Errorf("machine %v does not specify a start node: %v", dr.Id, dr.Nodes)
}

func legalTransition(curr string, transition string, m *pb.DefineRequest) (bool, string) {
	for _, node := range m.Nodes {
		if node.Id == curr {
			for _, t := range node.Transitions {
				if t.Id == transition {
					return true, t.Node
				}
			}
		}
	}
	return false, ""
}

func (s *server) Start(ctx context.Context, in *pb.TaskMessage) (*pb.TaskMessage, error) {
	m := s.getMachine(in.Machine)
	if m == nil {
		return nil, fmt.Errorf("No machine registered for %v", in.Machine)
	}

	if in.Id == "" {
		uid, err := uuid.NewV4()
		if err != nil {
			return nil, err
		}
		in.Id = uid.String()
	}
	in.Node, _ = startNode(m)

	s.instancesMutex.Lock()
	s.instances = append(s.instances, in)
	s.instancesMutex.Unlock()
	if err := s.record(ctx, "Start", in); err != nil {
		return nil, err
	}
	return in, nil
}

func (s *server) Relinquish(ctx context.Context, rr *pb.RelinquishRequest) (*pb.RelinquishReply, error) {
	client := determineClient(ctx, "")
	log.Printf("client %v relinquishing %v", client, rr.Instance)
	s.instancesMutex.RLock()
	defer s.instancesMutex.RUnlock()
	for _, instance := range s.instances {
		if instance.Id == rr.Instance {
			instance.Owner = ""
			return &pb.RelinquishReply{instance.Id, instance.Node}, nil
		}
	}
	return nil, fmt.Errorf("instance with id %v does not exist", rr.Instance)
}

func (s *server) Transition(ctx context.Context, tr *pb.TransitionRequest) (*pb.TransitionReply, error) {
	client := determineClient(ctx, "")
	log.Printf("client %v relinquishing %v", client, tr.Instance)
	s.instancesMutex.RLock()
	defer s.instancesMutex.RUnlock()
	for _, instance := range s.instances {
		if instance.Id == tr.Instance {
			m := s.getMachine(instance.Machine)
			ok, nextNode := legalTransition(instance.Node, tr.Transition, m)
			if !ok {
				return nil, fmt.Errorf("attempted illegal transition %v from %v in %v", tr.Transition, instance.Node, m.Id)
			}
			prevNode := instance.Node
			instance.Node = nextNode
			instance.NodeParams = tr.NodeParams
			return &pb.TransitionReply{Instance: tr.Instance, Node: nextNode, PrevNode: prevNode}, nil
		}
	}
	return nil, fmt.Errorf("instance with id %v does not exist", tr.Instance)
}

func (s *server) Ready(ctx context.Context, rr *pb.ReadyRequest) (*pb.TaskMessage, error) {
	// scan instance for a ready instance meeting criteria, then return it.
	// you would want a more sophisticated scheduling algorithm than this,
	// maybe a priority queue based on start time and penalizing retries
	rr.Client = determineClient(ctx, rr.Client)
	machine := rr.Machine
	if machine == "" {
		machine = "any machine"
	}
	node := rr.Node
	if node == "" {
		node = "any node"
	}
	log.Printf("client %v ready for work on %v, filtering node by %v", rr.Client, machine, node)

	s.instancesMutex.RLock()
	defer s.instancesMutex.RUnlock()
	for _, instance := range s.instances {
		if rr.Machine != "" && rr.Machine != instance.Machine {
			continue
		}
		if rr.Node != "" && rr.Node != instance.Node {
			continue
		}
		if instance.Owner == "" {
			instance.Owner = rr.Client
			return instance, nil
		}
	}
	return nil, fmt.Errorf("no available work")
}

func (s *server) Define(ctx context.Context, machine *pb.DefineRequest) (*pb.DefineReply, error) {
	id := machine.Id
	if s.getMachine(machine.Id) != nil {
		return nil, fmt.Errorf("Machine %v is already registered", machine.Id)
	}

	_, err := startNode(machine)
	if err != nil {
		return nil, err
	}
	if err := s.record(ctx, "Define", machine); err != nil {
		return nil, err
	}
	err = s.RegisterMachine(machine)
	if err != nil {
		return nil, err
	}
	return &pb.DefineReply{true, id, "Created machine.", ""}, nil
}

func (s *server) Changes(in *pb.ChangesRequest, stream pb.DistributedFSMRunner_ChangesServer) error {
	c := s.changeListener()
	defer s.closeChangeListener(c)
	for msg := range c {
		if err := stream.Send(msg); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()

	log.Printf("Readying server on %v", *addr)
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	ds := MakeServer()

	pb.RegisterDistributedFSMRunnerServer(s, ds)
	reflection.Register(s)
	log.Printf("Starting server on %v", *addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
