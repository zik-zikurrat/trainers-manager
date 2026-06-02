package webapi

import (
	"context"

	"trainers-manager/internal/config"
	"trainers-manager/internal/usecase"
	"trainers-manager/pkg/postgres"
)

type Generator struct {
	cfg *config.Config
	*postgres.Posgtres
}

func (g *Generator) Generate(ctx context.Context, prompt usecase.GeneratePrompt) {

}
