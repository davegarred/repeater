package main

import (
	"log"
	"net"

	pb "github.com/davegarred/repeater/grpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Pushfile(ctx context.Context, in *pb.Filecontent) (*pb.Filekey, error) {
	return &pb.Filekey{Key: "this is a key"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterFilemoverServer(s, &server{})
	//	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
