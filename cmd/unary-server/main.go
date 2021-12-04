package main

import (
	"context"
	"log"
	"net"

	pb "github.com/grpc-otlp/protos"
	"google.golang.org/grpc"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello.world"}, nil
}

func main() {
	// A grpc server
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	println("Hello, World!")
}
