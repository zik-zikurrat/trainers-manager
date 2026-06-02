package persistent

import (
	"context"
	"errors"
	"fmt"
	"trainers-manager/internal/entity"
	"trainers-manager/internal/repo"
	"trainers-manager/pkg/postgres"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

func (r *TrainingRepo) UpdateStructure(ctx context.Context, structure entity.TrainingStructure, structureID uuid.UUID) error {
	ct, err := r.Pool.Exec(ctx, updateTrainingStructureQuery, structure.Structure, structureID)
	if err != nil {
		return fmt.Errorf("update training structure: %w", err)
	}
	if ct.RowsAffected() == 0 {
		return repo.ErrNotFound
	}
	return nil
}

func (r *TrainingRepo) DeleteStructure(ctx context.Context, structureID uuid.UUID) error {
	ct, err := r.Pool.Exec(ctx, deleteTrainingStructureQuery, structureID)
	if err != nil {
		return fmt.Errorf("delete training structure: %w", err)
	}
	if ct.RowsAffected() == 0 {
		return repo.ErrNotFound
	}
	return nil
}

func (r *TrainingRepo) GetStructure(ctx context.Context, structureID uuid.UUID) (entity.TrainingStructure, error) {
	var s entity.TrainingStructure
	err := r.Pool.QueryRow(ctx, getTrainingStructure, structureID).Scan(&s.ID, &s.Structure, &s.CreatedAt, &s.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return entity.TrainingStructure{}, repo.ErrNotFound
	}
	if err != nil {
		return entity.TrainingStructure{}, fmt.Errorf("get training structure: %w", err)
	}
	return s, nil
}
