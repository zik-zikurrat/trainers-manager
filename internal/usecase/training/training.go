package training

import (
	"context"
	"trainers-manager/internal/entity"
	"trainers-manager/internal/repo"
	"trainers-manager/pkg/logger"

	"github.com/google/uuid"
)

type UseCase struct {
	log  *logger.Logger
	repo repo.TrainingRepo
}

func New(r repo.TrainingRepo, log *logger.Logger) *UseCase {
	return &UseCase{
		repo: r,
		log:  log,
	}
}

// CreateStructure -.
func (us *UseCase) CreateStructure(ctx context.Context, structure entity.TrainingStructure) error {
	const op = "training.CreateStructure"
	if err := us.repo.CreateStructure(ctx, structure); err != nil {
		us.log.Error("Failed to store structure", err, op)
		return err
	}
	return nil
}

// UpdateStructure -.
func (us *UseCase) UpdateStructure(ctx context.Context, structure entity.TrainingStructure, id uuid.UUID) error {
	const op = "training.UpdateStructure"
	if err := us.repo.UpdateStructure(ctx, structure, id); err != nil {
		us.log.Error("Failed to update structure", err, op)
		return err
	}
	return nil
}

// DeleteStructure -.
func (us *UseCase) DeleteStructure(ctx context.Context, id uuid.UUID) error {
	const op = "training.UpdateStructure"
	if err := us.repo.DeleteStructure(ctx, id); err != nil {
		us.log.Error("Failed to delete structure", err, op)
		return err
	}
	return nil
}
