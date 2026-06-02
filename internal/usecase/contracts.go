package usecase

import (
	"context"
	"trainers-manager/internal/entity"

	"github.com/google/uuid"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_training_test.go -package=usecase_test

type (
	// Training -.
	Training interface {
		// Structure
		CreateStructure(context.Context, entity.TrainingStructure) error
		UpdateStructure(context.Context, entity.TrainingStructure, uuid.UUID) error
		DeleteStructure(context.Context, uuid.UUID) error
		GetStructure(context.Context, uuid.UUID) (entity.TrainingStructure, error)
		// Exercises
		CreateExercise(context.Context, entity.Exercise) (uuid.UUID, error)
		UpdateExercise(context.Context, entity.Exercise, uuid.UUID) error
		DeleteExercise(context.Context, uuid.UUID) error
		ListExercises(context.Context) ([]entity.Exercise, error)
		// Training / Plan / History
		CreateTraining(context.Context) (uuid.UUID, error)
		StoreTrainingPlan(context.Context, entity.TrainingPlan) (uuid.UUID, error)
		UpdateTrainingPlan(context.Context, entity.TrainingPlan, uuid.UUID) error
		GetTrainingPlan(context.Context, uuid.UUID) (entity.TrainingPlan, error)
		GetPlanHistory(context.Context, uuid.UUID) ([]entity.TrainingPlanHistory, error)
		// Groups
		CreateGroup(context.Context, entity.TrainingGroup) (uuid.UUID, error)
		ListGroups(context.Context) ([]entity.TrainingGroup, error)
		UpdateGroup(context.Context, entity.TrainingGroup, uuid.UUID) error
		DeleteGroup(context.Context, uuid.UUID) error
		GetGroupByName(context.Context, string) (entity.TrainingGroup, error)
		// Generate
		Generate(ctx context.Context, trainType string, structureID uuid.UUID) (entity.TrainingPlan, error)
	}
)
type PlanGenerator interface {
	Generate(ctx context.Context, in GeneratePrompt) (GeneratedPlan, error)
}

type GeneratePrompt struct {
	Structure string
	Accent    string
	Skills    string
	Recent    []entity.TrainingPlan
	Pool      []entity.Exercise
}

// GeneratedPlan — что вернула LLM (после парсинга JSON).
type GeneratedPlan struct {
	ExerciseIDs []uuid.UUID
	PlanText    string
}
