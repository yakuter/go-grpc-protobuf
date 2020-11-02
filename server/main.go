package main

import (
	"context"
	"log"
	"net"

	"protobuf/pb"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedPersonServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Add(ctx context.Context, in *pb.PersonRequest) (*pb.PersonResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.PersonResponse{Name: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPersonServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
