package repo

import (
	"context"
	"errors"
	"trainers-manager/internal/entity"
	repoWorker "trainers-manager/internal/repo/workers"
	"trainers-manager/internal/usecase/dto"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("not found")
var ErrAlreadyExists = errors.New("already exists")

//go:generate mockgen -source=contracts.go -destination=mocks/mocks.go -package=mocks
type (
	TrainingStructureRepo interface {
		CreateStructure(context.Context, entity.TrainingStructure) error
		UpdateStructure(context.Context, entity.TrainingStructure, uuid.UUID) error
		DeleteStructure(context.Context, uuid.UUID) error
		GetStructure(context.Context, uuid.UUID) (entity.TrainingStructure, error)
		ListStructure(context.Context) ([]entity.TrainingStructure, error)
	}
	ExerciseRepo interface {
		CreateExercise(context.Context, entity.Exercise) (uuid.UUID, error)
		UpdateExercise(context.Context, dto.UpdateExerciseInput, uuid.UUID) error
		DeleteExercise(context.Context, uuid.UUID) error
		ListExercises(context.Context) ([]entity.Exercise, error)
	}
	TrainingRepo interface {
		CreateTraining(context.Context) (uuid.UUID, error)
	}
	TrainingPlanRepo interface {
		StoreTrainingPlan(context.Context, entity.TrainingPlan) (uuid.UUID, error)
		UpdateTrainingPlan(context.Context, dto.UpdateTrainingPlan) error
		GetTrainingPlan(context.Context, uuid.UUID) (entity.TrainingPlan, error)
		ListTrainingPlan(context.Context) ([]entity.TrainingPlan, error)
	}
	PlanHistoryRepo interface {
		GetPlanHistory(context.Context, uuid.UUID) ([]entity.TrainingPlanHistory, error)
	}
	TrainingGroupRepo interface {
		CreateGroup(context.Context, entity.TrainingGroup) (uuid.UUID, error)
		ListGroups(context.Context) ([]entity.TrainingGroup, error)
		UpdateGroup(context.Context, dto.UpdateGroupInput) error
		DeleteGroup(context.Context, uuid.UUID) error
		GetGroupByName(context.Context, string) (entity.TrainingGroup, error)
	}

	GenerationRepo interface {
		GetGroupByName(context.Context, string) (entity.TrainingGroup, error)
		GetStructure(context.Context, uuid.UUID) (entity.TrainingStructure, error)
		ListExercises(context.Context) ([]entity.Exercise, error)
		RecentPlans(context.Context, uuid.UUID, int) ([]entity.TrainingPlan, error)
		CreateTraining(context.Context) (uuid.UUID, error)
		LinkExercises(context.Context, uuid.UUID, []uuid.UUID) error
		StoreTrainingPlan(context.Context, entity.TrainingPlan) (uuid.UUID, error)
	}
	GenerationTaskRepo interface {
		CreateGenerationTask(context.Context, repoWorker.GenerationTask) error
		UpdateGenerationTask(context.Context, repoWorker.UpdateGenerationTask) error
		GetGenerationTask(context.Context, uuid.UUID) (repoWorker.GenerationTask, error)
	}

	PartitionsRepo interface {
		EnsureHistoryPartitions(context.Context, int) error
	}
)
