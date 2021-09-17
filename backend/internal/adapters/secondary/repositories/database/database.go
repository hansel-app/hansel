package database

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/hansel-app/hansel/internal/config"
)

func NewDatabase(cfg *config.Config) (*sqlx.DB, error) {
	dsn := buildDSN(cfg)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	db.MapperFunc(camelToSnake)

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

// From https://stackoverflow.com/a/56616250
func camelToSnake(s string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
