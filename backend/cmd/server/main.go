package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/hansel-app/hansel/internal/adapters/primary/handlers"
	"github.com/hansel-app/hansel/internal/adapters/secondary/repositories/database"
	"github.com/hansel-app/hansel/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	addr := fmt.Sprintf("localhost:%d", cfg.ServerPort)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	handlers.RegisterServices(s, db)
	log.Printf("Starting server on port %d...", cfg.ServerPort)
	err = s.Serve(l)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
