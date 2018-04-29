package main

import (
	"flag"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"	
	pb "github.com/lethain/dfsmr/dfsmr"
	"google.golang.org/grpc/reflection"	
)	

var (
	addr = flag.String("addr", ":5003", "Port to bind on")
)

type server struct{}

func (s *server) Start(ctx context.Context, in *pb.StartRequest) (*pb.AckReply, error) {
	return &pb.AckReply{true, "Success", ""}, nil
}

func (s *server) Changes(in *pb.ChangesRequest, stream pb.DistributedFSMRunner_ChangesServer) error {
	msgs := []string{"a", "b", "c", "d", "e"}
	for _, msg := range msgs {
		packed := &pb.ChangesReply{Command: msg, Client: "who-knows"}
		if err := stream.Send(packed); err != nil {
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
	pb.RegisterDistributedFSMRunnerServer(s, &server{})
	reflection.Register(s)
	log.Printf("Starting server on %v", *addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
