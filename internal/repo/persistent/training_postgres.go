package persistent

import (
	"context"
	"fmt"
	"trainers-manager/internal/entity"
	"trainers-manager/internal/repo"
	"trainers-manager/pkg/postgres"

	"github.com/google/uuid"
)

const _defaultEntityCap = 64

// TrainingRepo -.
type TrainingRepo struct {
	*postgres.Posgtres
}

// New -.
func New(pg *postgres.Posgtres) *TrainingRepo {
	return &TrainingRepo{pg}
}

func (r *TrainingRepo) CreateStructure(ctx context.Context, structure entity.TrainingStructure) error {
	_, err := r.Pool.Exec(ctx, insertTrainingStructureQuery, structure.Structure)
	if err != nil {
		return fmt.Errorf("insert training structure: %w", err)
	}
	return nil
}

func (r *TrainingRepo) UpdateStructure(ctx context.Context, structure entity.TrainingStructure, id uuid.UUID) error {
	ct, err := r.Pool.Exec(ctx, updateTrainingStructureQuery, structure.Structure, id)
	if err != nil {
		return fmt.Errorf("update training structure: %w", err)
	}
	if ct.RowsAffected() == 0 {
		return repo.ErrNotFound
	}
	return nil
}

func (r *TrainingRepo) DeleteStructure(ctx context.Context, id uuid.UUID) error {
	ct, err := r.Pool.Exec(ctx, deleteTrainingStructureQuery, id)
	if err != nil {
		return fmt.Errorf("delete training structure: %w", err)
	}
	if ct.RowsAffected() == 0 {
		return repo.ErrNotFound
	}
	return nil
}
