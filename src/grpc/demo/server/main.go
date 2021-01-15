package main

import (
	"context"
	"fmt"
	grpc "google.golang.org/grpc"
	pb "grpc-demo/proto"
	"net"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Name: "hello.world"}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", "127.0.0.1:"+"12111")

	fmt.Println("Start listening...")
	server.Serve(lis)
}
