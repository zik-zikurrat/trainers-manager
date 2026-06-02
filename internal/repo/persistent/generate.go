package persistent

import (
	"context"
	"fmt"
	"trainers-manager/internal/entity"

	"github.com/google/uuid"
)

func (r *TrainingRepo) CreateTraining(ctx context.Context) (uuid.UUID, error) {
	var id uuid.UUID
	if err := r.Pool.QueryRow(ctx, insertTrainingQuery).Scan(&id); err != nil {
		return uuid.Nil, fmt.Errorf("create training: %w", err)
	}
	return id, nil
}

func (r *TrainingRepo) RecentPlans(ctx context.Context, groupID uuid.UUID, limit int) ([]entity.TrainingPlan, error) {
	rows, err := r.Pool.Query(ctx, getRecentPlans, groupID, limit)
	if err != nil {
		return nil, fmt.Errorf("recent plans: %v", err)
	}
	defer rows.Close()
	out := make([]entity.TrainingPlan, 0, limit)
	for rows.Next() {
		var t entity.TrainingPlan
		if err := rows.Scan(&t.Accent, &t.Skills, &t.CreatedAt); err != nil {
			return nil, fmt.Errorf("recent plans scan: %w", err)
		}
		out = append(out, t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("recent plans rows: %w", err)
	}
	return out, nil
}
