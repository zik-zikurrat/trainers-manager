package usecase

import (
	"context"
	"trainers-manager/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_training_test.go -package=usecase_test

type (
	// Training -.
	Training interface {
		// Generate(context.Context, entity.Training) (entity.TrainingPlan, error)
		// History(context.Context) (entity.TrainingPlanHistory, error)
		StoreStructure(context.Context, entity.TrainingStructure) error
	}
)
