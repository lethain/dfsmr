package main

import (
	"flag"
	"log"
	"time"
	"io"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	pb "github.com/lethain/dfsmr/dfsmr"
)

var (
	addr = flag.String("addr", "localhost:5003", "Address to connect to")
)

func client() pb.DistributedFSMRunnerClient {
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return pb.NewDistributedFSMRunnerClient(conn)
}

func main() {
	flag.Parse()
	args := flag.Args()

	c := client()
	//defer conn.Close()

	
	if len(args) == 0 {		
		log.Fatalf("must specify at least one parameters, specified %v", len(args))
	}
	switch args[0] {
	case "start":
		name := "test-client"
		if len(args) > 1 {
			name = args[1]
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()		
		r, err := c.Start(ctx, &pb.StartRequest{Name: name})
		if err != nil {
			log.Fatalf("could not start: %v", err)
		}
		log.Printf("Started %v", r)		
	case "changes":
		ctx := context.Background()
		stream, err := c.Changes(ctx, &pb.ChangesRequest{})
		if err != nil {
			log.Fatalf("could not start: %v", err)
		}
		for {
			change, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("status: %v", status.Code(err))
				
				log.Fatalf("%v.Changes() = %v", c, err)
			}
			log.Printf("%v %v", change.Client, change.Command)
		}
	}
	
	
	


}
