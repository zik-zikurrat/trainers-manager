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

// StructureRepo -.
type StructureRepo struct {
	*postgres.Posgtres
}

// New -.
func NewStructureRepo(pg *postgres.Posgtres) *StructureRepo {
	return &StructureRepo{pg}
}

func (r *StructureRepo) CreateStructure(ctx context.Context, structure entity.TrainingStructure) error {
	_, err := r.Pool.Exec(ctx, insertTrainingStructureQuery, structure.Structure)
	if err != nil {
		return fmt.Errorf("insert training structure: %w", err)
	}
	return nil
}

func (r *StructureRepo) UpdateStructure(ctx context.Context, structure entity.TrainingStructure, structureID uuid.UUID) error {
	ct, err := r.Pool.Exec(ctx, updateTrainingStructureQuery, structure.Structure, structureID)
	if err != nil {
		return fmt.Errorf("update training structure: %w", err)
	}
	if ct.RowsAffected() == 0 {
		return repo.ErrNotFound
	}
	return nil
}

func (r *StructureRepo) DeleteStructure(ctx context.Context, structureID uuid.UUID) error {
	ct, err := r.Pool.Exec(ctx, deleteTrainingStructureQuery, structureID)
	if err != nil {
		return fmt.Errorf("delete training structure: %w", err)
	}
	if ct.RowsAffected() == 0 {
		return repo.ErrNotFound
	}
	return nil
}

func (r *StructureRepo) GetStructure(ctx context.Context, structureID uuid.UUID) (entity.TrainingStructure, error) {
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

func (r *StructureRepo) ListStructure(ctx context.Context) ([]entity.TrainingStructure, error) {
	rows, err := r.Pool.Query(ctx, listTrainingStructure)
	if err != nil {
		return nil, fmt.Errorf("list structures: %w", err)
	}
	defer rows.Close()

	out := make([]entity.TrainingStructure, 0, _defaultEntityCap)
	for rows.Next() {
		var e entity.TrainingStructure
		if err := rows.Scan(&e.ID, &e.Structure, &e.CreatedAt, &e.UpdatedAt); err != nil {
			return nil, fmt.Errorf("list structures scan: %w", err)
		}
		out = append(out, e)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("list structures rows: %w", err)
	}
	return out, nil
}
