package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/hansel-app/hansel/internal/adapters/secondary/repositories/database"
	"github.com/hansel-app/hansel/internal/auth"
	"github.com/hansel-app/hansel/internal/config"
	"github.com/hansel-app/hansel/internal/server"
)

var jwtDuration = 7 * 24 * time.Hour

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	addr := fmt.Sprintf(":%d", cfg.ServerPort)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	jwtManager, err := auth.NewJWTManager(cfg.SecretKey, &jwtDuration)
	if err != nil {
		log.Fatalf("failed to create JWT manager: %v", err)
	}

	s := server.New(db, jwtManager)
	log.Printf("Starting server on port %d...", cfg.ServerPort)
	err = s.Serve(l)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
