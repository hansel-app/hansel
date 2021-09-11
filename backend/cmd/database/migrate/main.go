package main

import (
	"github.com/hansel-app/hansel/internal/adapters/secondary/repositories/database"
	"github.com/hansel-app/hansel/internal/config"
	"log"
	"github.com/rubenv/sql-migrate"
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

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}
	steps, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Printf("Applied %d migrations!", steps)
}
