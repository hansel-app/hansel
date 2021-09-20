package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/hansel-app/hansel/internal/adapters/secondary/repositories/database"
	"github.com/hansel-app/hansel/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	seedPath := filepath.Join("seeds", "seed.sql")
	c, ioErr := ioutil.ReadFile(filepath.Clean(seedPath))
	if ioErr != nil {
		log.Fatalf("failed to read seed file: %v", err)
	}

	sql := string(c)
	_, err = db.Exec(sql)
	if err != nil {
		log.Fatalf("failed to seed database: %v", err)
	}

	log.Printf("Seeded database!")
}
