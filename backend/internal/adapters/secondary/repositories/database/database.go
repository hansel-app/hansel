package database

import (
	"fmt"

	"github.com/hansel-app/hansel/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New(cfg *config.Config) (*sqlx.DB, error) {
	dsn := buildDSN(cfg)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return db, nil
}

func buildDSN(cfg *config.Config) string {
	dsn := ""
	if cfg.DBName != "" {
		dsn += fmt.Sprintf("dbname=%v", cfg.DBName)
	}
	if cfg.DBHostname != "" {
		dsn += fmt.Sprintf(" host=%v", cfg.DBHostname)
	}
	if cfg.DBPort != 0 {
		dsn += fmt.Sprintf(" port=%v", cfg.DBPort)
	}
	if cfg.DBUsername != "" {
		dsn += fmt.Sprintf(" user=%v", cfg.DBUsername)
	}
	if cfg.DBPassword != "" {
		dsn += fmt.Sprintf(" password=%v", cfg.DBPassword)
	}
	if cfg.DBSSLMode != "" {
		dsn += fmt.Sprintf(" sslmode=%v", cfg.DBSSLMode)
	}
	return dsn
}
