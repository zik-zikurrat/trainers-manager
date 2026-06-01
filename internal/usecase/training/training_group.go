package training

import (
	"context"

	"trainers-manager/internal/entity"

	"github.com/google/uuid"
)

func (us *UseCase) CreateGroup(ctx context.Context, g entity.TrainingGroup) (uuid.UUID, error) {
	id, err := us.repo.CreateGroup(ctx, g)
	if err != nil {
		us.log.Error("Failed to store group", err, "training.CreateGroup")
	}
	return id, err
}

func (us *UseCase) ListGroups(ctx context.Context) ([]entity.TrainingGroup, error) {
	groups, err := us.repo.ListGroups(ctx)
	if err != nil {
		us.log.Error("Failed to list groups", err, "training.ListGroups")
		return nil, err
	}
	return groups, nil
}

func (us *UseCase) UpdateGroup(ctx context.Context, g entity.TrainingGroup, id uuid.UUID) error {
	if err := us.repo.UpdateGroup(ctx, g, id); err != nil {
		us.log.Error("Failed to update group", err, "training.UpdateGroup")
		return err
	}
	return nil
}

func (us *UseCase) DeleteGroup(ctx context.Context, id uuid.UUID) error {
	if err := us.repo.DeleteGroup(ctx, id); err != nil {
		us.log.Error("Failed to delete group", err, "training.DeleteGroup")
		return err
	}
	return nil
}

func (us *UseCase) GetGroupByName(ctx context.Context, name string) (entity.TrainingGroup, error) {
	return us.repo.GetGroupByName(ctx, name)
}
