package main

import (
	"context"
	"log"
	"time"

	hello_world "github.com/cwyu57/grpc-go-hello-world/hello-world"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := hello_world.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &hello_world.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("cloud not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
