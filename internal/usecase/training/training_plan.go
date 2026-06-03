package training

import (
	"context"

	"trainers-manager/internal/entity"

	"github.com/google/uuid"
)

func (us *UseCase) StoreTrainingPlan(ctx context.Context, p entity.TrainingPlan) (uuid.UUID, error) {
	id, err := us.repo.StoreTrainingPlan(ctx, p)
	if err != nil {
		us.log.Error("Failed to store plan", err, "training.StoreTrainingPlan")
	}
	return id, err
}

func (us *UseCase) UpdateTrainingPlan(ctx context.Context, p entity.TrainingPlan, id uuid.UUID) error {
	if err := us.repo.UpdateTrainingPlan(ctx, p, id); err != nil {
		us.log.Error("Failed to update plan", err, "training.UpdateTrainingPlan")
		return err
	}
	return nil
}

func (us *UseCase) GetTrainingPlan(ctx context.Context, id uuid.UUID) (entity.TrainingPlan, error) {
	return us.repo.GetTrainingPlan(ctx, id)
}

func (us *UseCase) ListTrainingPlan(ctx context.Context) ([]entity.TrainingPlan, error) {
	return us.repo.ListTrainingPlan(ctx)
}
