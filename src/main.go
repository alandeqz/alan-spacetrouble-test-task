package main

import (
	"context"
	"log/slog"

	"alan-tabeo-test-task/src/config"
	"alan-tabeo-test-task/src/drivers"
)

func main() {
	slog.Info("starting the application")

	defer func() {
		slog.Info("exiting the application")
	}()

	cfg, err := config.NewConfiguration()
	if err != nil {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	postgreSQL, err := drivers.InitPostgreSQL(ctx, cfg.DSN)
	if err != nil {
		return
	}

	defer func() {
		_ = postgreSQL.Connection.Close()
	}()
}
