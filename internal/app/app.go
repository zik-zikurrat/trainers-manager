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
	"trainers-manager/pkg/workers"
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

	genEventCh := make(chan workers.GenEvent, 100)

	// Repository
	exerciseRepo := persistent.NewExerciseRepo(pg)
	trainingRepo := persistent.NewTrainingRepo(pg)
	structureRepo := persistent.NewStructureRepo(pg)
	groupRepo := persistent.NewGroupRepo(pg)
	partitionsRepo := persistent.NewPartitionRepo(pg)
	planHistoryRepo := persistent.NewPlanHistoryRepo(pg)
	planRepo := persistent.NewPlanRepo(pg)
	generationTaskRepo := persistent.NewGenerationTaskRepo(pg)
	generationRepo := persistent.NewGeneratorRepo(pg)
	gen := webapi.New(cfg, genEventCh)
	// UseCase
	exerciseUseCase := training.NewExerciseUseCase(l, exerciseRepo)
	structureUseCase := training.NewStructureUseCase(l, structureRepo)
	planUseCase := training.NewPlanUseCase(l, planRepo)
	trainingUseCase := training.NewTrainingUseCase(l, trainingRepo)
	planHistoryUseCase := training.NewPlanHistoryUseCase(l, planHistoryRepo)
	generateUseCase := training.NewGenerateUseCase(l, generationRepo, gen)
	groupUseCase := training.NewGroupUseCase(l, groupRepo)

	// Generation Worker
	genWorker := workers.NewGenWorker(l, genEventCh, generationTaskRepo)
	go genWorker.Run(ctx)

	// Partition
	StartPartitionMaintainer(ctx, partitionsRepo, l)

	httpserver := httpserver.New(ctx, l, cfg)
	restapi.NewRouter(
		httpserver.App,
		cfg,
		pg.Pool,
		l,
		trainingUseCase,
		exerciseUseCase,
		structureUseCase,
		planUseCase,
		planHistoryUseCase,
		groupUseCase,
		generateUseCase,
		genEventCh,
	)
	httpserver.Start()

	select {
	case <-ctx.Done():
		close(genEventCh)
		l.Info("app - Run - shutdown signal received")
	case err := <-httpserver.Notify():
		close(genEventCh)
		l.Error(fmt.Errorf("app - Run - httpserver.Notify: %w", err))
	}

	return httpserver.Shutdown()
}
