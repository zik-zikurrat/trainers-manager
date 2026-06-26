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
	"trainers-manager/pkg/ssoclient"
	"trainers-manager/pkg/workers"

	ssov1 "buf.build/gen/go/zik-zikurrat-sso/sso/protocolbuffers/go/sso/v1"
	"connectrpc.com/connect"
	"github.com/gofiber/fiber/v2/log"
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

	ssoClient := ssoclient.New(cfg.SSO.URL)

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
		generationTaskRepo,
		genEventCh,
	)
	httpserver.Start()

	// REGISTER ENDPOINTS IN SSO
	var endpoints []*ssov1.Endpoint
	for _, r := range httpserver.App.GetRoutes() {
		if !isCorrectRoute(r.Method, r.Path) {
			continue
		}
		endpoints = append(endpoints, &ssov1.Endpoint{
			Method: r.Method,
			Path:   r.Path,
		})
	}

	resp, err := ssoClient.RegisterService(ctx, connect.NewRequest(&ssov1.RegisterServiceRequest{
		Name:      "trainers-manager",
		Metadata:  map[string]string{"env": cfg.Env},
		Endpoints: endpoints,
	}))
	if err != nil {
		log.Error("failed to register in SSO", err)
	}
	log.Info("registered in SSO", "service_id", resp.Msg.GetMsg())

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

func isCorrectRoute(method string, path string) bool {
	if path == "/" {
		return false
	}
	badMethods := []string{"HEAD", "CONNECT", "OPTIONS", "TRACE"}
	for _, val := range badMethods {
		if method == val {
			return false
		}
	}
	return true
}
