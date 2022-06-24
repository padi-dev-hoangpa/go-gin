package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "example/hoangdz/gin/square/square"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedSquareRpcServer
}

func (s *server) FindSquareUnary(ctx context.Context, in *pb.Input) (*pb.Output, error) {
	log.Printf("Received: %v", in.GetNumber())
	result := in.GetNumber() * in.GetNumber()
	return &pb.Output{Number: in.GetNumber(), Result: result}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSquareRpcServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
