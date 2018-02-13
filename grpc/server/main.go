package server

import (
	"log"
	"net"
	"errors"

	pb "github.com/davegarred/repeater/grpc/proto"
	"github.com/davegarred/repeater/persist"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	defaultPort = ":50051"
)

var storer Storer

type Storer interface {
	Store(string, string, string) error
	//Retrieve(string) (*persist.StoredObject, error)
	//Delete(string) error
}

type server struct{}

func (s *server) Pushfile(ctx context.Context, in *pb.Filecontent) (*pb.Filekey, error) {
	key := "a key"
	if err := storer.Store("image/slayer", key, in.Content); err != nil {
		return nil,errors.New("some sort of problem, I'm a bit too drunk to figure out")
	}
	return &pb.Filekey{Key: key}, nil
}

func main() {
	store := persist.NewMemStore()
	StartGRPCServer(store, defaultPort)
}
func StartGRPCServer(s Storer, listener string) {
	storer = s
	lis, err := net.Listen("tcp", listener)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterFilemoverServer(grpcServer, &server{})
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
