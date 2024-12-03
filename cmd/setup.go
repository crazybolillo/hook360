package main

import (
	"context"
	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
	"github.com/caarlos0/env/v11"
	"github.com/crazybolillo/hook360"
	"log/slog"
	"net/url"
)

type setupCfg struct {
	DatabaseURL string `env:"DATABASE_URL,required"`
}

func setup(_ context.Context) int {
	var cfg setupCfg
	err := env.Parse(&cfg)
	if err != nil {
		slog.Error("Failed to read settings from environment variables", "reason", err)
		return 1
	}

	u, err := url.Parse(cfg.DatabaseURL)
	if err != nil {
		slog.Error("Failed to parse database URL", "reason", err)
		return 1
	}

	db := dbmate.New(u)
	db.FS = hook360.Migrations

	slog.Info("Running migrations...")
	err = db.Migrate()
	if err != nil {
		slog.Error("Failed to migrate database", "reason", err)
		return 1
	}

	return 0
}
