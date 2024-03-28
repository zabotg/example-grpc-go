package client

import (
	"context"
	"log"
	"os"
	"time"

	pb "journey-grpc/hellopb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// StartClient initializes and starts the gRPC client.
func StartClient() {
	// Establish a connection to the server using insecure credentials.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// Determine the name to be used in the greeting, default is "World".
	name := "World"
	if len(os.Args) > 2 {
		name = os.Args[2]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Make the call to the service and print the response.
	response, err := client.SayHello(ctx, &pb.HelloRequest{Name: name, LastName: "BRASIL"})
	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}
	log.Printf("Greeting: %s", response.GetMessage())
}
