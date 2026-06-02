package webapi

import (
	"context"
	"errors"

	"trainers-manager/internal/usecase"
)

type StubGenerator struct{}

func (StubGenerator) Generate(ctx context.Context, in usecase.GeneratePrompt) (usecase.GeneratedPlan, error) {
	return usecase.GeneratedPlan{}, errors.New("llm generator not implemented")
}
