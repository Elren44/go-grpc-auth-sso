package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/Elren44/go-grpc-auth-sso/internal/app/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	grpcapp := grpcapp.New(log, grpcPort)

	app := &App{
		GRPCServer: grpcapp,
	}

	return app
}
