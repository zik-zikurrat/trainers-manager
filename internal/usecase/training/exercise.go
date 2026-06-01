package training

import (
	"context"

	"trainers-manager/internal/entity"

	"github.com/google/uuid"
)

func (us *UseCase) CreateExercise(ctx context.Context, e entity.Exercise) (uuid.UUID, error) {
	const op = "training.CreateExercise"
	id, err := us.repo.CreateExercise(ctx, e)
	if err != nil {
		us.log.Error("Failed to store exercise", err, op)
		return uuid.Nil, err
	}
	return id, nil
}

func (us *UseCase) ListExercises(ctx context.Context) ([]entity.Exercise, error) {
	const op = "training.ListExercises"
	exercises, err := us.repo.ListExercises(ctx)
	if err != nil {
		us.log.Error("Failed to list exercises", err, op)
		return nil, err
	}
	return exercises, nil
}

func (us *UseCase) UpdateExercise(ctx context.Context, e entity.Exercise, id uuid.UUID) error {
	const op = "training.UpdateExercise"
	if err := us.repo.UpdateExercise(ctx, e, id); err != nil {
		us.log.Error("Failed to update exercise", err, op)
		return err
	}
	return nil
}

func (us *UseCase) DeleteExercise(ctx context.Context, id uuid.UUID) error {
	const op = "training.DeleteExercise"
	if err := us.repo.DeleteExercise(ctx, id); err != nil {
		us.log.Error("Failed to delete exercise", err, op)
		return err
	}
	return nil
}
func (us *UseCase) LinkExercises(ctx context.Context, trainingID uuid.UUID, exerciseIDs []uuid.UUID) error {
	if err := us.repo.LinkExercises(ctx, trainingID, exerciseIDs); err != nil {
		us.log.Error("Failed to link exercises", err, "training.LinkExercises")
		return err
	}
	return nil
}
