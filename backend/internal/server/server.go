package server

import (
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"

	"github.com/hansel-app/hansel/internal/adapters/primary/handlers"
)

func New(db *sqlx.DB) *grpc.Server {
	s := grpc.NewServer()
	handlers.RegisterServices(s, db)
	return s
}
