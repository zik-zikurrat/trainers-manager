package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"trainers-manager/internal/config"
	"trainers-manager/internal/controller/restapi"
	"trainers-manager/internal/repo/persistent"
	"trainers-manager/internal/repo/webapi"
	"trainers-manager/internal/usecase/training"
	"trainers-manager/pkg/httpserver"
	"trainers-manager/pkg/logger"
	"trainers-manager/pkg/postgres"
)

func Run(cfg *config.Config) error {
	Migrate(cfg)
	l := logger.New(cfg.Logging.Level)

	pg, err := postgres.New(cfg, l)
	if err != nil {
		l.Fatal("app - Run - postgres.New: %w", err)
		return err
	}
	defer pg.Close()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	trainingRepo := persistent.New(pg)
	trainingUseCase := training.New(persistent.New(pg), webapi.Generator{}, l)
	StartPartitionMaintainer(ctx, trainingRepo, l)

	httpserver := httpserver.New(ctx, l, cfg)
	restapi.NewRouter(httpserver.App, cfg, trainingUseCase, l)
	httpserver.Start()

	select {
	case <-ctx.Done():
		l.Info("app - Run - shutdown signal received")
	case err := <-httpserver.Notify():
		l.Error(fmt.Errorf("app - Run - httpserver.Notify: %w", err))
	}

	return httpserver.Shutdown()
}
