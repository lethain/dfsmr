package main

import (
	"fmt"
	"flag"
	"log"
	"net"
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	pb "github.com/lethain/dfsmr/dfsmr"
	"google.golang.org/grpc/reflection"
)

var (
	addr = flag.String("addr", ":5003", "Port to bind on")
)


type server struct{
	listeners []chan *pb.ChangesReply
	machines []*pb.DefineRequest
	machinesMutex *sync.RWMutex
	changesMutex *sync.RWMutex
}

func MakeServer() *server {
	srv := &server{}
	srv.machinesMutex = &sync.RWMutex{}
	srv.changesMutex = &sync.RWMutex{}
	return srv
}

func (s *server) RegisterMachine(newMachine *pb.DefineRequest) error {
	s.machinesMutex.Lock()
	defer s.machinesMutex.Unlock()
	for _, machine := range s.machines {
		if newMachine.Name == machine.Name {
			return fmt.Errorf("Machine %v is already registered", newMachine.Name)
		}
	}
	s.machines = append(s.machines, newMachine)
	return nil
}

func (s *server) Machines(ctx context.Context, mr *pb.MachinesRequest) (*pb.MachinesReply, error) {
	s.machinesMutex.RLock()
	defer s.machinesMutex.RUnlock()
	return &pb.MachinesReply{Machines: s.machines}, nil
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
			s.listeners[i] = s.listeners[len(s.listeners) - 1]
			s.listeners = s.listeners[:len(s.listeners)-1]
			break
		}
	}
	s.changesMutex.Unlock()
	close(c)
}

func (s *server) record(ctx context.Context, op string, cmd interface{}) error {
	client := "unknown"
	p, ok := peer.FromContext(ctx)
	if ok {
		client = p.Addr.Network() + "://" + p.Addr.String()
	}
	cmdStr := fmt.Sprintf("%s %+v", op, cmd)
	cr := &pb.ChangesReply{Command: cmdStr, Client: client}
	s.changesMutex.RLock()
	defer s.changesMutex.RUnlock()
	for _, lc := range s.listeners {
		lc <- cr
	}
	return nil
}

func (s *server) Start(ctx context.Context, in *pb.StartRequest) (*pb.AckReply, error) {
	s.machinesMutex.RLock()
	defer s.machinesMutex.RUnlock()

	for _, m := range s.machines {
		if m.Name == in.Name {
			if err := s.record(ctx, "Start", in); err != nil {
				return nil, err
			}
			return &pb.AckReply{true, "Success", ""}, nil
		}
	}
	return nil, fmt.Errorf("No machine registered for %v", in.Name)
	
}

func (s *server) Define(ctx context.Context, machine *pb.DefineRequest) (*pb.DefineReply, error) {
	name := machine.Name
	if err := s.record(ctx, "Define", machine); err != nil {
		return nil, err
	}
	err := s.RegisterMachine(machine)
	if err != nil {
		return nil, err
	}
	return &pb.DefineReply{true, name, "Created machine.", ""}, nil
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
