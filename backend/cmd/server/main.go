package main

import (
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
	err = s.Serve(l)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
