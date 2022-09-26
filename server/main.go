package main

import (
	"context"
	"fmt"
	"log"
	"net"

	hello_world "github.com/cwyu57/grpc-go-hello-world/hello-world"
	"google.golang.org/grpc"
)

// server is used to implement hello_world.GreeterServer.
type server struct {
	hello_world.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *hello_world.HelloRequest) (*hello_world.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &hello_world.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	hello_world.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
