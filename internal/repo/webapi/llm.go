package webapi

import (
	"trainers-manager/internal/config"
	"trainers-manager/pkg/postgres"
)

type Generator struct {
	cfg *config.Config
	*postgres.Posgtres
}

// func New(cfg *config.Config)

// func (g *Generator) Generate(ctx context.Context, prompt usecase.GeneratePrompt) {

// }
