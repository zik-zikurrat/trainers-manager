package persistent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"trainers-manager/internal/entity"
	"trainers-manager/internal/repo"
	"trainers-manager/internal/usecase/dto"
	"trainers-manager/pkg/postgres"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// PlanRepo -.
type PlanRepo struct {
	*postgres.Posgtres
}

// New -.
func NewPlanRepo(pg *postgres.Posgtres) *PlanRepo {
	return &PlanRepo{pg}
}
func (r *PlanRepo) ListTrainingPlan(ctx context.Context) ([]entity.TrainingPlan, error) {
	rows, err := r.Pool.Query(ctx, listTrainingPlanQuery)
	if err != nil {
		return nil, fmt.Errorf("list training plan: %w", err)
	}
	defer rows.Close()

	out := make([]entity.TrainingPlan, 0, _defaultEntityCap)
	for rows.Next() {
		var e entity.TrainingPlan
		if err := rows.Scan(&e.ID, &e.Plan, &e.Status, &e.Accent, &e.Skills, &e.CreatedAt, &e.UpdatedAt); err != nil {
			return nil, fmt.Errorf("list training plans scan: %w", err)
		}
		out = append(out, e)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("list training plans rows: %w", err)
	}
	return out, nil
}
func (r *PlanRepo) CreateTrainingPlan(ctx context.Context, p entity.TrainingPlan) (uuid.UUID, error) {
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("create plan begin: %w", err)
	}
	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx, insertTrainingPlanQuery,
		p.Plan, p.Status, p.TrainID, p.GroupID, p.Accent, p.Skills, p.TrainingStructureID,
	).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return uuid.Nil, fmt.Errorf("create plan insert: %w", err)
	}

	snapshot, err := json.Marshal(p)
	if err != nil {
		return uuid.Nil, fmt.Errorf("create plan marshal: %w", err)
	}
	if _, err := tx.Exec(ctx, insertPlanHistoryQuery, p.ID, entity.HistoryActionCreate, snapshot); err != nil {
		return uuid.Nil, fmt.Errorf("create plan history: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return uuid.Nil, fmt.Errorf("create plan commit: %w", err)
	}
	return p.ID, nil
}

func (r *PlanRepo) UpdateTrainingPlan(ctx context.Context, p dto.UpdateTrainingPlan) error {
	tx, err := r.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("update plan begin: %w", err)
	}
	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx, updateTrainingPlanQuery, p.Plan, p.Accent, p.Skills, p.ID).
		Scan(&p.ID, &p.TrainingStructureID, &p.CreatedAt, &p.UpdatedAt)
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
	if _, err := tx.Exec(ctx, insertPlanHistoryQuery, p.ID, entity.HistoryActionUpdate, snapshot); err != nil {
		return fmt.Errorf("update plan history: %w", err)
	}

	return tx.Commit(ctx)
}

func (r *PlanRepo) GetTrainingPlan(ctx context.Context, id uuid.UUID) (entity.TrainingPlan, error) {
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
