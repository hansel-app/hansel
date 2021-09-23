package main

import (
	"log"

	"github.com/hansel-app/hansel/internal/adapters/secondary/repositories/database"
	"github.com/hansel-app/hansel/internal/config"
	"github.com/omeid/pgerror"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	dbName := cfg.DBName
	cfg.DBName = ""
	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	_, err = db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		if err := pgerror.DuplicateDatabase(err); err != nil {
			log.Printf("Database '%s' already exists", dbName)
			return
		}
		log.Fatalf("failed to create database: %v", err)
	}

	log.Printf("Successfully created database '%s'", dbName)
}
