package repo

import (
	"context"
	"errors"
	"trainers-manager/internal/entity"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("not found")

type TrainingRepo interface {

	// Structure
	CreateStructure(context.Context, entity.TrainingStructure) error
	UpdateStructure(context.Context, entity.TrainingStructure, uuid.UUID) error
	DeleteStructure(context.Context, uuid.UUID) error
	// Exercise
	CreateExercise(context.Context, entity.Exercise) (uuid.UUID, error)
	UpdateExercise(context.Context, entity.Exercise, uuid.UUID) error
	DeleteExercise(context.Context, uuid.UUID) error
	ListExercises(context.Context) ([]entity.Exercise, error)
	// Training / Plan / History
	CreateTraining(context.Context) (uuid.UUID, error)
	LinkExercises(context.Context, uuid.UUID, []uuid.UUID) error
	StoreTrainingPlan(context.Context, entity.TrainingPlan) (uuid.UUID, error)
	UpdateTrainingPlan(context.Context, entity.TrainingPlan, uuid.UUID) error
	GetTrainingPlan(context.Context, uuid.UUID) (entity.TrainingPlan, error)
	GetPlanHistory(context.Context, uuid.UUID) ([]entity.TrainingPlanHistory, error)
}
