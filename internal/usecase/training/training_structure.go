package training

import (
	"context"
	"trainers-manager/internal/entity"

	"github.com/google/uuid"
)

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

// GetStructure -.
func (us *UseCase) GetStructure(ctx context.Context, id uuid.UUID) (entity.TrainingStructure, error) {
	const op = "training.TrainingStructure"
	structure, err := us.repo.GetStructure(ctx, id)
	if err != nil {
		us.log.Error("Failed to get structure", err, op)
		return entity.TrainingStructure{}, err
	}
	return structure, nil
}

// ListStructure -.
func (us *UseCase) ListStructure(ctx context.Context) ([]entity.TrainingStructure, error) {
	const op = "training.TrainingStructure"
	structures, err := us.repo.ListStructure(ctx)
	if err != nil {
		us.log.Error("Failed to get structure", err, op)
		return nil, err
	}
	return structures, nil
}
