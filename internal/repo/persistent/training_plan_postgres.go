package persistent

import (
	"context"
	"fmt"
	"trainers-manager/internal/entity"
)

func (r *TrainingRepo) ListTrainingPlan(ctx context.Context) ([]entity.TrainingPlan, error) {
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
