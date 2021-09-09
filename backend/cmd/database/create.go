package main

import (
	"github.com/hansel-app/hansel/internal/adapters/secondary/repositories/database"
	"github.com/hansel-app/hansel/internal/config"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	dbName := cfg.DBName
	cfg.DBName = ""
	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db.MustExec("CREATE DATABASE " + dbName)
	log.Printf("Successfully created database '%s'", dbName)
}
