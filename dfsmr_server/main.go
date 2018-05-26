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
	changes chan *pb.ChangesReply
	machines []*pb.DefineRequest
	machinesMutex *sync.RWMutex
}

func MakeServer() *server {
	srv := &server{}
	srv.changes = make(chan *pb.ChangesReply)
	srv.machinesMutex = &sync.RWMutex{}
	return srv
}

func (s *server) RegisterMachine(machine *pb.DefineRequest) {
	s.machinesMutex.Lock()
	s.machines = append(s.machines, machine)
	s.machinesMutex.Unlock()
}

func (s *server) Machines() []*pb.DefineRequest {
	s.machinesMutex.RLock()
	defer s.machinesMutex.RUnlock()
	return s.machines
}


func (s *server) record(ctx context.Context, op string, cmd interface{}) error {
	client := "unknown"
	p, ok := peer.FromContext(ctx)
	if ok {
		client = p.Addr.Network() + "://" + p.Addr.String()
	}
	cmdStr := fmt.Sprintf("%s %+v", op, cmd)
	s.changes <- &pb.ChangesReply{Command: cmdStr, Client: client}
	return nil
}

func (s *server) Start(ctx context.Context, in *pb.StartRequest) (*pb.AckReply, error) {
	if err := s.record(ctx, "Start", in); err != nil {
		return nil, err
	}
	return &pb.AckReply{true, "Success", ""}, nil
}

func (s *server) Define(ctx context.Context, machine *pb.DefineRequest) (*pb.DefineReply, error) {
	name := machine.Name
	if err := s.record(ctx, "Define", machine); err != nil {
		return nil, err
	}
	s.RegisterMachine(machine)
	return &pb.DefineReply{true, name, "Created machine.", ""}, nil
}


func (s *server) Changes(in *pb.ChangesRequest, stream pb.DistributedFSMRunner_ChangesServer) error {
	for msg := range s.changes {
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
