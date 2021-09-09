package main

import (
	"context"
	"github.com/hansel-app/hansel/protobuf/gemsapi"
	"google.golang.org/grpc"
	"log"
)

// TODO: Remove this proof of concept. To be executed when the server is running.
//       To run this, create a database with the same name as DB_NAME in your .env file.
//       Then, add the 'gems' table using:
//       CREATE TABLE gems (id BIGSERIAL PRIMARY KEY, message VARCHAR);
func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := gemsapi.NewGemServiceClient(conn)

	dropResponse, err := c.Drop(context.Background(), &gemsapi.DropRequest{Message: "Hello World!"})
	if err != nil {
		log.Fatalf("error when adding gem: %s", err)
	}

	getResponse, err := c.Get(context.Background(), &gemsapi.GetRequest{Id: dropResponse.Id})
	if err != nil {
		log.Fatalf("error when getting gem: %s", err)
	}

	log.Printf("Response from server: %v", getResponse)
}
