package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"trainers-manager/internal/config"
	"trainers-manager/internal/controller/restapi"
	"trainers-manager/internal/repo/persistent"
	"trainers-manager/internal/usecase/training"
	"trainers-manager/pkg/httpserver"
	"trainers-manager/pkg/logger"
	"trainers-manager/pkg/postgres"
)

func Run(cfg *config.Config) error {
	// Run migrations
	Migrate(cfg)

	// Create logger
	l := logger.New(cfg.Logging.Level)

	// Create postgres connection
	pg, err := postgres.New(cfg, l)
	if err != nil {
		l.Fatal("app - Run - postgres.New: %w", err)
		return err
	}
	defer pg.Close()

	// Create training usecase
	trainingUseCase := training.New(persistent.New(pg), l)

	// Create new http server
	httpserver := httpserver.New(l, cfg)
	restapi.NewRouter(httpserver.App, cfg, trainingUseCase, l)
	httpserver.Start()

	// Notify
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: %s", s.String())
	case err := <-httpserver.Notify():
		l.Error(fmt.Errorf("app - Run - httpserver.Notify: %w", err))
	}

	return nil
}
