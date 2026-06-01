package training

import (
	"context"
	"trainers-manager/internal/entity"

	"github.com/google/uuid"
)

// CreateTrainingGroup -.
func (us *UseCase) CreateTrainingGroup(ctx context.Context, trainingGroup entity.TrainingGroup) error {
	const op = "training.CreateTrainingGroup"

	if err := us.repo.CreateTrainingGroup(ctx, trainingGroup); err != nil {
		us.log.Error("Failed to store trainingGroup", err, op)
		return err
	}
	return nil
}

// UpdateTrainingGroup -.
func (us *UseCase) UpdateTrainingGroup(ctx context.Context, trainingGroup entity.TrainingGroup, id uuid.UUID) error {
	const op = "training.UpdateTrainingGroup"
	if err := us.repo.UpdateTrainingGroup(ctx, trainingGroup, id); err != nil {
		us.log.Error("Failed to update trainingGroup", err, op)
		return err
	}
	return nil
}

// DeleteTrainingGroup -.
func (us *UseCase) DeleteTrainingGroup(ctx context.Context, id uuid.UUID) error {
	const op = "training.UpdateTrainingGroup"
	if err := us.repo.DeleteTrainingGroup(ctx, id); err != nil {
		us.log.Error("Failed to delete trainingGroup", err, op)
		return err
	}
	return nil
}

// GetTrainingGroup -.
func (us *UseCase) GetTrainingGroup(ctx context.Context, id uuid.UUID) (entity.TrainingGroup, error) {
	const op = "training.TrainingGroup"
	TrainingGroup, err := us.repo.GetTrainingGroup(ctx, id)
	if err != nil {
		us.log.Error("Failed to get trainingGroup", err, op)
		return entity.TrainingGroup{}, err
	}
	return TrainingGroup, nil
}
