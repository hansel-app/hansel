package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/hansel-app/hansel/internal/adapters/secondary/repositories/database"
	"github.com/hansel-app/hansel/internal/config"
	"github.com/hansel-app/hansel/internal/constants"
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

	gretelAvatarPath := filepath.Join("assets", "gretel-avatar.png")
	imageBytes, ioErr := ioutil.ReadFile(filepath.Clean(gretelAvatarPath))
	if ioErr != nil {
		log.Fatalf("failed to read gretel avatar file: %v", ioErr)
	}

	var qb = goqu.Dialect("postgres")
	sql, args, _ := qb.Update(goqu.T("users")).Where(goqu.C("id").Eq(constants.GretelUserId)).Set(goqu.Record{
		"avatar": imageBytes,
	}).Prepared(true).ToSQL()

	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatalf("failed to prepare statement: %w", err)
	}
	_, err = stmt.Exec(args...)

	if err != nil {
		log.Fatalf("failed to seed database with Gretel's avatar: %v", err)
	}

	log.Printf("Seeded database!")
}
