package main

import (
	"fmt"
	"log"
	"os"
	"flag"

	pb "github.com/davegarred/repeater/grpcfile"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	get := flag.String("get", "", "Retrieve a file using this key")
	fmt.Printf("getting: %v\n", *get)
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//	c := pb.NewGreeterClient(conn)
	c := pb.NewFilemoverClient(conn)

	if *get != "" {
		getFile(c, *get)
	} else {
		pushFile(c)
	}
}
func getFile(c pb.FilemoverClient, key string) {
	result, err := c.Getfile(context.Background(), &pb.Filekey{key})
	if err != nil {
		log.Fatalf("Could not retrieve file '%s', recieved %v\n", key, err)
	}
	log.Printf("Retrieved: %v\n", string(result.Content))
}

func pushFile(c pb.FilemoverClient) {
	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.Pushfile(context.Background(), &pb.Filecontent{Content: []byte(name)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Key: %s", r.Key)
}
