package main

import (
	"context"
	"fmt"
	"log"
	"net"

	hello "github.com/foresightyj/hello-grpc-go/HelloWorld"
	"google.golang.org/grpc"
)

const (
	port = ":8088"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	hello.UnimplementedHelloServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &hello.HelloReply{Message: "Hello from golang: " + in.GetName()}, nil
}

var _ hello.HelloServiceServer = server{}

func main() {
	fmt.Println("grpc server at:", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hello.RegisterHelloServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
