package training

import (
	"context"

	"trainers-manager/internal/entity"
	"trainers-manager/internal/repo"
	"trainers-manager/pkg/logger"

	"github.com/google/uuid"
)

type ExerciseUseCase struct {
	l *logger.Logger
	r repo.ExerciseRepo
}

func NewExerciseUseCase(l *logger.Logger, r repo.ExerciseRepo) *ExerciseUseCase {
	return &ExerciseUseCase{
		l: l,
		r: r,
	}

}

func (us *ExerciseUseCase) CreateExercise(ctx context.Context, e entity.Exercise) (uuid.UUID, error) {
	const op = "training.CreateExercise"
	id, err := us.r.CreateExercise(ctx, e)
	if err != nil {
		us.l.Error("Failed to store exercise", err, op)
		return uuid.Nil, err
	}
	return id, nil
}

func (us *ExerciseUseCase) ListExercises(ctx context.Context) ([]entity.Exercise, error) {
	const op = "training.ListExercises"
	exercises, err := us.r.ListExercises(ctx)
	if err != nil {
		us.l.Error("Failed to list exercises", err, op)
		return nil, err
	}
	return exercises, nil
}

func (us *ExerciseUseCase) UpdateExercise(ctx context.Context, e entity.Exercise, id uuid.UUID) error {
	const op = "training.UpdateExercise"
	if err := us.r.UpdateExercise(ctx, e, id); err != nil {
		us.l.Error("Failed to update exercise", err, op)
		return err
	}
	return nil
}

func (us *ExerciseUseCase) DeleteExercise(ctx context.Context, id uuid.UUID) error {
	const op = "training.DeleteExercise"
	if err := us.r.DeleteExercise(ctx, id); err != nil {
		us.l.Error("Failed to delete exercise", err, op)
		return err
	}
	return nil
}
func (us *ExerciseUseCase) LinkExercises(ctx context.Context, trainingID uuid.UUID, exerciseIDs []uuid.UUID) error {
	const op = "training.LinkExercises"
	if err := us.r.LinkExercises(ctx, trainingID, exerciseIDs); err != nil {
		us.l.Error("Failed to link exercises", err, op)
		return err
	}
	return nil
}
