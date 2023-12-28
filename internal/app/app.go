package app

import (
	"log/slog"

	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       string
}

// New creates a new App
func New(log *slog.Logger, port string) *App {
	return &App{
		log:        log,
		port:       port,
		gRPCServer: grpc.NewServer(),
	}
}
