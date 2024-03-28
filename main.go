package main

import (
	"journey-grpc/client"
	"journey-grpc/server"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("Usage: go run main.go [client|server]")
		return
	}

	mode := os.Args[1]
	switch mode {
	case "server":
		server.StartServer()
	case "client":
		client.StartClient()
	default:
		println("Unknown mode:", mode)
	}
}
