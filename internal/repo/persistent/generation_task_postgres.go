package persistent

import (
	"context"
	"errors"
	"fmt"
	"trainers-manager/internal/repo"
	repoWorker "trainers-manager/internal/repo/workers"
	"trainers-manager/pkg/postgres"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type GenerationTaskRepo struct {
	*postgres.Posgtres
}

// New -.
func NewGenerationTaskRepo(pg *postgres.Posgtres) *GenerationTaskRepo {
	return &GenerationTaskRepo{pg}
}

func (r *GenerationTaskRepo) CreateGenerationTask(ctx context.Context, t repoWorker.GenerationTask) error {
	err := r.Pool.QueryRow(ctx, insertGenerationTask, t.ID, t.Status, t.Error).Scan(&t.ID)
	if err != nil {
		return fmt.Errorf("insert generation task: %w", err)
	}
	return nil
}

func (r *GenerationTaskRepo) UpdateGenerationTask(ctx context.Context, t repoWorker.UpdateGenerationTask) error {
	ct, err := r.Pool.Exec(ctx, updateGenerationTask, t.Status, t.Error, t.ID)
	if err != nil {
		return fmt.Errorf("update generation task: %w", err)
	}
	if ct.RowsAffected() == 0 {
		return repo.ErrNotFound
	}
	return nil
}

func (r *GenerationTaskRepo) GetGenerationTask(ctx context.Context, id uuid.UUID) (repoWorker.GenerationTask, error) {
	var t repoWorker.GenerationTask
	err := r.Pool.QueryRow(ctx, getGenerationTask, id).
		Scan(&t.ID, &t.Status, &t.Error, &t.CreatedAt, &t.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return repoWorker.GenerationTask{}, repo.ErrNotFound
	}
	if err != nil {
		return repoWorker.GenerationTask{}, fmt.Errorf("get generation task: %w", err)
	}
	return t, nil
}
