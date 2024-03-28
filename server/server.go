package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "journey-grpc/hellopb"
)

// Default port for the gRPC server.
const serverPort = ":50051"

// StartServer initializes and starts the gRPC server.
func StartServer() {
	// Listen on the TCP port.
	listener, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", serverPort, err)
	}

	// Create a new gRPC server instance.
	grpcServer := grpc.NewServer()

	// Register the Greeter service with the server.
	pb.RegisterGreeterServer(grpcServer, &greeterService{})

	log.Printf("gRPC server is listening at %v", listener.Addr())

	// Start serving gRPC requests.
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// greeterService implements the GreeterServer interface.
type greeterService struct {
	pb.UnimplementedGreeterServer
}

// SayHello generates response to a HelloRequest.
func (s *greeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	name := req.GetName()
	log.Printf("Received: %v", name)
	message := "Hello " + name + ". Welcome to the gRPC project."
	return &pb.HelloReply{Message: message}, nil
}
