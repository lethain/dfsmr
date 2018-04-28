package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/lethain/dfsmr/dfsmr"
)

const (
	address     = "localhost:5003"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)	
	}
	defer conn.Close()
	c := pb.NewDistributedFSMRunnerClient(conn)
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Start(ctx, &pb.StartRequest{Name: "hi"})
	if err != nil {
		log.Fatalf("could not start: %v", err)		
	}
	log.Printf("Started %v", r)
	

}
