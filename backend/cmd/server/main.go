package main

import (
	"github.com/hansel-app/hansel/internal/handlers"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8000")
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
