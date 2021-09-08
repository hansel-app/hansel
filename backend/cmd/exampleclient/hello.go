package main

import (
	"context"
	"github.com/hansel-app/hansel/protobuf/gem"
	"google.golang.org/grpc"
	"log"
)

// TODO: Remove this proof of concept. To be executed when the server is running.
func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := gem.NewGemServiceClient(conn)

	response, err := c.HelloWorld(context.Background(), &gem.HelloWorldRequest{Name: "Uncle Soo"})
	if err != nil {
		log.Fatalf("Error when calling HelloWorld: %s", err)
	}

	log.Printf("Response from server: %s", response.Message)
}
