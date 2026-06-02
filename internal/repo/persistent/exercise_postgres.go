package persistent

import (
	"context"
	"fmt"

	"trainers-manager/internal/entity"
	"trainers-manager/internal/repo"

	"github.com/google/uuid"
)

func (r *TrainingRepo) CreateExercise(ctx context.Context, e entity.Exercise) (uuid.UUID, error) {
	var id uuid.UUID
	err := r.Pool.QueryRow(ctx, insertExerciseQuery, e.Muscle, e.Description).Scan(&id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("insert exercise: %w", err)
	}
	return id, nil
}

func (r *TrainingRepo) ListExercises(ctx context.Context) ([]entity.Exercise, error) {
	rows, err := r.Pool.Query(ctx, listExercisesQuery)
	if err != nil {
		return nil, fmt.Errorf("list exercises: %w", err)
	}
	defer rows.Close()

	out := make([]entity.Exercise, 0, _defaultEntityCap)
	for rows.Next() {
		var e entity.Exercise
		if err := rows.Scan(&e.ID, &e.Muscle, &e.Description, &e.CreatedAt, &e.UpdatedAt); err != nil {
			return nil, fmt.Errorf("list exercises scan: %w", err)
		}
		out = append(out, e)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("list exercises rows: %w", err)
	}
	return out, nil
}

func (r *TrainingRepo) UpdateExercise(ctx context.Context, e entity.Exercise, id uuid.UUID) error {
	ct, err := r.Pool.Exec(ctx, updateExerciseQuery, e.Muscle, e.Description, id)
	if err != nil {
		return fmt.Errorf("update exercise: %w", err)
	}
	if ct.RowsAffected() == 0 {
		return repo.ErrNotFound
	}
	return nil
}

func (r *TrainingRepo) DeleteExercise(ctx context.Context, id uuid.UUID) error {
	ct, err := r.Pool.Exec(ctx, deleteExerciseQuery, id)
	if err != nil {
		return fmt.Errorf("delete exercise: %w", err)
	}
	if ct.RowsAffected() == 0 {
		return repo.ErrNotFound
	}
	return nil
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
