package main

import (
	"flag"
	"log"
	"time"
	"io"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	pb "github.com/lethain/dfsmr/dfsmr"
)

var (
	addr = flag.String("addr", "localhost:5003", "Address to connect to")
)

func client() (*grpc.ClientConn, pb.DistributedFSMRunnerClient, error) {
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	fcli := pb.NewDistributedFSMRunnerClient(conn)
	return conn, fcli, err
}

func main() {
	flag.Parse()
	args := flag.Args()

	grpcConn, c, err := client()
	if err != nil {
		log.Fatalf("%v", err)
	}

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
				grpcConn.Close()
				grpcConn, c, _ = client()
				if status.Code(err) == codes.Unavailable {
					time.Sleep(1.0)
					continue
				} else {
					log.Fatalf("%v.Changes() = %v", c, err)
				}
			}
			log.Printf("%v %v", change.Client, change.Command)
		}
	}





}
