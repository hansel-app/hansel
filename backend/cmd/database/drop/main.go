package main

import (
	"log"
	"os"

	"github.com/hansel-app/hansel/internal/adapters/secondary/repositories/database"
	"github.com/hansel-app/hansel/internal/config"
)

const overrideEnvVar = "OVERRIDE_DATABASE_DELETE_SAFEGUARD"

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	if cfg.Environment != "development" && cfg.Environment != "test" {
		if os.Getenv(overrideEnvVar) != "1" {
			log.Fatalf("Unable to drop database in a non-development and non-test environment\n"+
				"Re-run this command with the prefix '%s=1' to override this safeguard", overrideEnvVar)
			return
		}
	}

	dbName := cfg.DBName
	cfg.DBName = ""
	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db.MustExec("DROP DATABASE " + dbName)
	log.Printf("Successfully dropped database '%s'", dbName)
}
