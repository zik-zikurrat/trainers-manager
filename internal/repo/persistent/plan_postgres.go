package persistent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"trainers-manager/internal/entity"
	"trainers-manager/internal/repo"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *TrainingRepo) StoreTrainingPlan(ctx context.Context, p entity.TrainingPlan) (uuid.UUID, error) {
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("store plan begin: %w", err)
	}
	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx, insertTrainingPlanQuery,
		p.Plan, p.Status, p.TrainID, p.GroupID, p.Accent, p.Skills, p.TrainingStructureID,
	).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return uuid.Nil, fmt.Errorf("store plan insert: %w", err)
	}

	snapshot, err := json.Marshal(p)
	if err != nil {
		return uuid.Nil, fmt.Errorf("store plan marshal: %w", err)
	}
	if _, err := tx.Exec(ctx, insertPlanHistoryQuery, p.ID, entity.HistoryActionCreate, snapshot); err != nil {
		return uuid.Nil, fmt.Errorf("store plan history: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return uuid.Nil, fmt.Errorf("store plan commit: %w", err)
	}
	return p.ID, nil
}

func (r *TrainingRepo) UpdateTrainingPlan(ctx context.Context, p entity.TrainingPlan, id uuid.UUID) error {
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("update plan begin: %w", err)
	}
	defer tx.Rollback(ctx)

	p.ID = id
	err = tx.QueryRow(ctx, updateTrainingPlanQuery, p.Plan, p.Accent, p.Skills, id).
		Scan(&p.TrainID, &p.TrainingStructureID, &p.CreatedAt, &p.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return repo.ErrNotFound
	}
	if err != nil {
		return fmt.Errorf("update plan: %w", err)
	}

	snapshot, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("update plan marshal: %w", err)
	}
	if _, err := tx.Exec(ctx, insertPlanHistoryQuery, id, entity.HistoryActionUpdate, snapshot); err != nil {
		return fmt.Errorf("update plan history: %w", err)
	}

	return tx.Commit(ctx)
}

func (r *TrainingRepo) GetTrainingPlan(ctx context.Context, id uuid.UUID) (entity.TrainingPlan, error) {
	var p entity.TrainingPlan
	err := r.Pool.QueryRow(ctx, getTrainingPlanQuery, id).Scan(
		&p.ID, &p.Plan, &p.TrainID, &p.Accent, &p.Skills, &p.TrainingStructureID, &p.CreatedAt, &p.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return entity.TrainingPlan{}, repo.ErrNotFound
	}
	if err != nil {
		return entity.TrainingPlan{}, fmt.Errorf("get plan: %w", err)
	}
	return p, nil
}

func (r *TrainingRepo) GetPlanHistory(ctx context.Context, planID uuid.UUID) ([]entity.TrainingPlanHistory, error) {
	rows, err := r.Pool.Query(ctx, getPlanHistoryQuery, planID)
	if err != nil {
		return nil, fmt.Errorf("get plan history: %w", err)
	}
	defer rows.Close()

	out := make([]entity.TrainingPlanHistory, 0, _defaultEntityCap)
	for rows.Next() {
		var h entity.TrainingPlanHistory
		if err := rows.Scan(&h.ID, &h.PlanID, &h.Action, &h.Snapshot, &h.CreatedAt); err != nil {
			return nil, fmt.Errorf("get plan history scan: %w", err)
		}
		out = append(out, h)
	}
	return out, rows.Err()
}

func (r *TrainingRepo) EnsureHistoryPartitions(ctx context.Context, ahead int) error {
	if _, err := r.Pool.Exec(ctx, "SELECT ensure_history_partitions($1)", ahead); err != nil {
		return fmt.Errorf("ensure history partitions: %w", err)
	}
	return nil
}
