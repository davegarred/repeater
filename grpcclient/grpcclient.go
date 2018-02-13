package main

import (
	"log"
	"os"

	pb "github.com/davegarred/repeater/grpcfile"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//	c := pb.NewGreeterClient(conn)
	c := pb.NewFilemoverClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	// r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	r, err := c.Pushfile(context.Background(), &pb.Filecontent{Content: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Key: %s", r.Key)
}
