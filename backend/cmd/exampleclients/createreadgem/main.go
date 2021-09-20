package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/hansel-app/hansel/protobuf/gemsapi"
)

// TODO: Remove this proof of concept. To be executed when the server is running.
//       To run this, create a database with the same name as DB_NAME in your .env.development.local file.
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

	dropResponse, err := c.Drop(context.Background(), &gemsapi.DropRequest{GemMessage: &gemsapi.GemMessage{
		Message:    "hello!",
		Latitude:   1.0,
		Longitude:  1.0,
		CreatorId:  1,
		ReceiverId: 1,
		Color:      gemsapi.GemColor_BLUE,
	}})
	if err != nil {
		log.Fatalf("error when adding gem: %s", err)
	}

	log.Printf("Response from server: %v", dropResponse)
}
