package main

import (
	"fmt"
	"log/slog"
	"os"
	"sso/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	log := setupLogger(cfg.Env)

	log.Info("Starting application",
		slog.String("env", cfg.Env),
		slog.String("storage_path", cfg.StoragePath),
		slog.String("token_ttl", cfg.TokenTTL.String()),
		slog.Int("port", cfg.GRPC.Port),
		slog.String("timeout", cfg.GRPC.Timeout.String()))

	log.Debug("debug message")

	log.Error("error message")

	log.Warn("warning message")

	// ToDo: инициализировать объект конфигурации

	// ToDo: инициализировать логгер

	// ToDo: запустить приложение

	// ToDo: запустить сервер
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "dev":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "prod":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)

	}
	return log
}
