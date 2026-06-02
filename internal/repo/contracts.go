package repo

import (
	"context"
	"errors"
	"trainers-manager/internal/entity"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("not found")
var ErrAlreadyExists = errors.New("already exists")

type TrainingRepo interface {

	// Structure
	CreateStructure(context.Context, entity.TrainingStructure) error
	UpdateStructure(context.Context, entity.TrainingStructure, uuid.UUID) error
	DeleteStructure(context.Context, uuid.UUID) error
	GetStructure(context.Context, uuid.UUID) (entity.TrainingStructure, error)
	ListStructure(context.Context) ([]entity.TrainingStructure, error)
	// Exercise
	CreateExercise(context.Context, entity.Exercise) (uuid.UUID, error)
	UpdateExercise(context.Context, entity.Exercise, uuid.UUID) error
	DeleteExercise(context.Context, uuid.UUID) error
	ListExercises(context.Context) ([]entity.Exercise, error)
	LinkExercises(context.Context, uuid.UUID, []uuid.UUID) error
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
	// Generation
	RecentPlans(context.Context, uuid.UUID, int) ([]entity.TrainingPlan, error)
	// Partitions
	EnsureHistoryPartitions(context.Context, int) error
}
