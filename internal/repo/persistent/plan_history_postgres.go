package persistent

import (
	"context"
	"fmt"

	"trainers-manager/internal/entity"
	"trainers-manager/pkg/postgres"

	"github.com/google/uuid"
)

// PlanHistoryRepo -.
type PlanHistoryRepo struct {
	*postgres.Posgtres
}

// New -.
func NewPlanHistoryRepo(pg *postgres.Posgtres) *PlanHistoryRepo {
	return &PlanHistoryRepo{pg}
}
func (r *PlanHistoryRepo) GetPlanHistory(ctx context.Context, planID uuid.UUID) ([]entity.TrainingPlanHistory, error) {
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
