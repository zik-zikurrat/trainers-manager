package repo

import (
	"context"
	"trainers-manager/internal/entity"
)

type TrainingRepo interface {
	// StoreTraining(context.Context, entity.Training) error
	// StoreTrainingPlan(context.Context, []entity.TrainingPlan) error
	StoreStructure(context.Context, entity.TrainingStructure) error
	// GetHistory(context.Context) ([]entity.TrainingPlanHistory, error)
	// GetTrainingPlan(context.Context, time.Time) (entity.TrainingPlan, error)
}
