package main

import (
	"flag"
	"log"
	"net"

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
}

func MakeServer() *server {
	srv := &server{}
	srv.changes = make(chan *pb.ChangesReply)
	return srv
}


func (s *server) record(ctx context.Context, cmd string) error {
	client := "unknown"
	p, ok := peer.FromContext(ctx)
	if ok {
		client = p.Addr.Network() + "://" + p.Addr.String()
	}
	s.changes <- &pb.ChangesReply{Command: cmd, Client: client}
	return nil
}

func (s *server) Start(ctx context.Context, in *pb.StartRequest) (*pb.AckReply, error) {
	if err := s.record(ctx, in.Name); err != nil {
		return nil, err
	}
	return &pb.AckReply{true, "Success", ""}, nil
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
