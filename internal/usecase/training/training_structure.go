package training

import (
	"context"
	"trainers-manager/internal/entity"
	"trainers-manager/internal/repo"
	"trainers-manager/pkg/logger"

	"github.com/google/uuid"
)

type StructureUseCase struct {
	l *logger.Logger
	r repo.TrainingStructureRepo
}

func NewStructureUseCase(l *logger.Logger, r repo.TrainingStructureRepo) *StructureUseCase {
	return &StructureUseCase{
		l: l,
		r: r,
	}

}

// CreateStructure -.
func (us *StructureUseCase) CreateStructure(ctx context.Context, structure entity.TrainingStructure) error {
	const op = "training.CreateStructure"
	if err := us.r.CreateStructure(ctx, structure); err != nil {
		us.l.Error("Failed to store structure", err, op)
		return err
	}
	return nil
}

// UpdateStructure -.
func (us *StructureUseCase) UpdateStructure(ctx context.Context, structure entity.TrainingStructure, id uuid.UUID) error {
	const op = "training.UpdateStructure"
	if err := us.r.UpdateStructure(ctx, structure, id); err != nil {
		us.l.Error("Failed to update structure", err, op)
		return err
	}
	return nil
}

// DeleteStructure -.
func (us *StructureUseCase) DeleteStructure(ctx context.Context, id uuid.UUID) error {
	const op = "training.DeleteStructure"
	if err := us.r.DeleteStructure(ctx, id); err != nil {
		us.l.Error("Failed to delete structure", err, op)
		return err
	}
	return nil
}

// GetStructure -.
func (us *StructureUseCase) GetStructure(ctx context.Context, id uuid.UUID) (entity.TrainingStructure, error) {
	const op = "training.GetStructure"
	structure, err := us.r.GetStructure(ctx, id)
	if err != nil {
		us.l.Error("Failed to get structure", err, op)
		return entity.TrainingStructure{}, err
	}
	return structure, nil
}

// ListStructure -.
func (us *StructureUseCase) ListStructure(ctx context.Context) ([]entity.TrainingStructure, error) {
	const op = "training.ListStructure"
	structures, err := us.r.ListStructure(ctx)
	if err != nil {
		us.l.Error("Failed to list structure", err, op)
		return nil, err
	}
	return structures, nil
}
