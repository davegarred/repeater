package server

import (
	"log"
	"net"
	"context"

	pb "github.com/davegarred/repeater/grpcfile"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/reflection"
	"gopkg.in/h2non/filetype.v1"
	"github.com/google/uuid"
)

const (
	defaultPort = ":50051"
)

type Storer interface {
	Store(string, string, string) error
	//Retrieve(string) (*persist.StoredObject, error)
	//Delete(string) error
}

type GrpcServer struct{
	storer Storer
}

func NewServer(s Storer) *GrpcServer {
	return &GrpcServer{s}
}

func (s *GrpcServer) Pushfile(ctx context.Context, in *pb.Filecontent) (*pb.Filekey, error) {
	key := uuid.New().String()

	data := in.Content
	mimetype := "application/octet-stream"
	kind,unknown := filetype.Match([]byte(data))
	if unknown == nil && kind.MIME.Value != ""  {
		mimetype = kind.MIME.Value
	}

	if err := s.storer.Store(mimetype, key, in.Content); err != nil {
		return nil, err
	}
	return &pb.Filekey{Key: key}, nil
}

func (s *GrpcServer) Start(listener string) {
	lis, err := net.Listen("tcp", listener)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterFilemoverServer(grpcServer, s)
	grpcServer.Serve(lis)
}
