package main

import (
	"context"
	"github.com/caarlos0/env/v11"
	"github.com/crazybolillo/hook360/repo"
	"github.com/crazybolillo/hook360/web"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"net/http"
	"net/url"
	"os"
)

type serveCfg struct {
	DatabaseURL string `env:"DATABASE_URL,required"`
	Token       string `env:"TOKEN,required"`
	Port        string `env:"PORT" envDefault:"8080"`
}

func serve(ctx context.Context) int {
	var config serveCfg
	err := env.Parse(&config)
	if err != nil {
		slog.Error("Failed to read settings from environment variables", "reason", err)
		return 1
	}

	u, err := url.Parse(config.DatabaseURL)
	if err != nil {
		slog.Error("Invalid database URL", "reason", err)
		return 1
	}

	slog.Info(
		"Connecting to database",
		slog.String("host", u.Host),
		slog.String("port", u.Port()),
		slog.String("user", u.User.Username()),
		slog.String("database", u.Path[1:]),
	)

	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		slog.Error("Failed to establish database connection", "reason", err)
		return 1
	}

	eventRepo := repo.NewEvent(pool)
	handler := web.NewHandler(eventRepo)

	http.HandleFunc("POST /"+config.Token, handler.Handle)

	address := ":" + config.Port
	slog.Info("Starting server", "address", address)

	err = http.ListenAndServe(address, nil)
	if err != nil {
		slog.Error("Failed to start server", "reason", err)
		return 1
	}

	return 0
}
