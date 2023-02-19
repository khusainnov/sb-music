package main

import (
	"context"

	"github.com/caarlos0/env/v6"
	"github.com/khusainnov/sb-music/internal/app"
	"github.com/khusainnov/sb-music/internal/config"
	"go.uber.org/zap"
)

func main() {
	log, _ := zap.NewProduction()

	cfg := &config.Config{
		Logger: log,
	}

	if err := env.Parse(cfg); err != nil {
		cfg.Logger.Fatal("error due parse envs", zap.Error(err))
	}

	if err := app.Run(context.Background(), cfg); err != nil {
		cfg.Logger.Fatal("", zap.Error(err))
	}
}
