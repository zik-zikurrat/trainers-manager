package webapi

import (
	"context"
	"errors"
	"net/http"
	"time"

	"trainers-manager/internal/config"
	"trainers-manager/internal/usecase"
)

type Generator struct {
	cfg    *config.Config
	client *http.Client
}

func New(cfg *config.Config) *Generator {
	return &Generator{
		cfg:    cfg,
		client: &http.Client{Timeout: 60 * time.Second},
	}
}

// Generate — сигнатура ТОЧНО как в usecase.PlanGenerator,
// поэтому *Generator автоматически реализует интерфейс.
func (g *Generator) Generate(ctx context.Context, prompt usecase.GeneratePrompt) (usecase.GeneratedPlan, error) {
	// TODO: построить промпт из prompt, POST на api, распарсить JSON
	return usecase.GeneratedPlan{}, errors.New("llm generator not implemented")
}
