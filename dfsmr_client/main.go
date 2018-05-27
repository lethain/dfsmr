package main

import (
	"flag"
	"log"
	"time"
	"io"
	"os"
	"fmt"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"

	"github.com/lethain/dfsmr/machines"
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

func start(args []string, c pb.DistributedFSMRunnerClient) {
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
}

func getMachines(args []string, c pb.DistributedFSMRunnerClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Machines(ctx, &pb.MachinesRequest{})
	if err != nil {
		log.Fatalf("could not retrieve machines: %v", err)
	}
	for _, m := range r.Machines {
		fmt.Printf("%v:\t%v", m.Name, m.Nodes)
	}
}

func getInstances(args []string, c pb.DistributedFSMRunnerClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Instances(ctx, &pb.InstancesRequest{})
	if err != nil {
		log.Fatalf("could not retrieve instances: %v", err)
	}
	for _, m := range r.Instances {
		fmt.Printf("%v:\t%#v", m.Name, m)
	}
}


func define(args []string, c pb.DistributedFSMRunnerClient) {
	if len(args) < 2 {
		log.Fatalf("must specify a YAML file to load: %v", args)
	}
	path := args[1]
	machine, err := machines.FromFile(path)
	if err != nil {
		log.Fatalf("Failed to load machine at %v: %v", path, err)
	}
	dr := machines.AsDefineRequest(machine)
	log.Printf("%+v", dr)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Define(ctx, dr)
	if err != nil {
		log.Fatalf("could not define: %v", err)
	}
	log.Printf("Defined %v", r)
}

func changes(grpcConn *grpc.ClientConn, c pb.DistributedFSMRunnerClient) {
	for {
		ctx := context.Background()
		stream, err := c.Changes(ctx, &pb.ChangesRequest{})
		code := status.Code(err)

		if code == codes.Unavailable || code == codes.Canceled || code == codes.Aborted {
			time.Sleep(1.0)
			grpcConn.Close()
			grpcConn, c, _ = client()
			continue
		}

		if err != nil {
			log.Fatalf("could not start: %v", err)
		}

		var streamErr error
		for {
			change, streamErr := stream.Recv()
			if streamErr == io.EOF {
				return
			}
			if streamErr != nil {
				break
			}
			log.Printf("%v %v", change.Client, change.Command)
		}

		streamCode := status.Code(streamErr)
		if streamErr == nil || streamCode == codes.Unavailable || streamCode == codes.Canceled || streamCode == codes.Aborted {
			continue
		} else {
			log.Fatalf("%v.Changes() = %v", c, streamErr)
		}
	}
}

func main() {
	flag.Usage = func() {
		cmds := []string{"start", "define", "changes", "machines", "instances"}
		inCmdArr := []string{}
		if len(os.Args) > 0 {
			inCmdArr = os.Args[1:len(os.Args)]
		}
		inCmd := strings.Join(inCmdArr, " ")
		validCmds := strings.Join(cmds, ", ")

		fmt.Fprintf(os.Stderr, "Specified command was: %v\nValid commands are: %v\n", inCmd, validCmds)
		flag.PrintDefaults()
	}

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatalf("must specify at least one parameters, specified %v", len(args))
	}

	grpcConn, c, err := client()
	if err != nil {
		log.Fatalf("%v", err)
	}

	switch args[0] {
	case "start":
		start(args, c)
	case "define":
		define(args, c)
	case "changes":
		changes(grpcConn, c)
	case "machines":
		getMachines(args, c)
	case "instances":
		getInstances(args, c)		
	default:
		flag.Usage()
	}
}
