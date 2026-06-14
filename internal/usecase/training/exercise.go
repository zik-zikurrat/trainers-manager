package training

import (
	"context"

	"trainers-manager/internal/entity"
	"trainers-manager/internal/repo"
	"trainers-manager/internal/usecase/dto"
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
		us.l.Error("Failed to store exercise: %v (op=%v)", err, op)
		return uuid.Nil, err
	}
	return id, nil
}

func (us *ExerciseUseCase) ListExercises(ctx context.Context) ([]entity.Exercise, error) {
	const op = "training.ListExercises"
	exercises, err := us.r.ListExercises(ctx)
	if err != nil {
		us.l.Error("Failed to list exercises: %v (op=%v)", err, op)
		return nil, err
	}
	return exercises, nil
}

func (us *ExerciseUseCase) UpdateExercise(ctx context.Context, e dto.UpdateExerciseInput, id uuid.UUID) error {
	const op = "training.UpdateExercise"
	if err := us.r.UpdateExercise(ctx, e, id); err != nil {
		us.l.Error("Failed to update exercise: %v (op=%v)", err, op)
		return err
	}
	return nil
}

func (us *ExerciseUseCase) DeleteExercise(ctx context.Context, id uuid.UUID) error {
	const op = "training.DeleteExercise"
	if err := us.r.DeleteExercise(ctx, id); err != nil {
		us.l.Error("Failed to delete exercise: %v (op=%v)", err, op)
		return err
	}
	return nil
}
