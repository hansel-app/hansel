package main

import (
	"github.com/hansel-app/hansel/internal/adapters/primary/handlers"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	handlers.RegisterServices(s)
	err = s.Serve(l)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
