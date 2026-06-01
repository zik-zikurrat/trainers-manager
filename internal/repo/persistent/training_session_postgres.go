package persistent

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (r *TrainingRepo) CreateTraining(ctx context.Context) (uuid.UUID, error) {
	var id uuid.UUID
	if err := r.Pool.QueryRow(ctx, insertTrainingQuery).Scan(&id); err != nil {
		return uuid.Nil, fmt.Errorf("create training: %w", err)
	}
	return id, nil
}

func (r *TrainingRepo) LinkExercises(ctx context.Context, trainingID uuid.UUID, exerciseIDs []uuid.UUID) error {
	if len(exerciseIDs) == 0 {
		return nil
	}
	if _, err := r.Pool.Exec(ctx, linkExercisesQuery, trainingID, exerciseIDs); err != nil {
		return fmt.Errorf("link exercises: %w", err)
	}
	return nil
}
