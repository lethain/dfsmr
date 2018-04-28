package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"	
	pb "github.com/lethain/dfsmr/dfsmr"
	"google.golang.org/grpc/reflection"	
)	

const (
	port = ":5003"
)

type server struct{}

func (s *server) Start(ctx context.Context, in *pb.StartRequest) (*pb.AckReply, error) {
	return &pb.AckReply{true, "Success", ""}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDistributedFSMRunnerServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
