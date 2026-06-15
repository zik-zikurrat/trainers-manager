package training

import (
	"context"

	"trainers-manager/internal/entity"
	"trainers-manager/internal/repo"
	"trainers-manager/pkg/logger"

	"github.com/google/uuid"
)

type TrainingGroupUseCase struct {
	l *logger.Logger
	r repo.TrainingGroupRepo
}

func NewGroupUseCase(l *logger.Logger, r repo.TrainingGroupRepo) *TrainingGroupUseCase {
	return &TrainingGroupUseCase{
		l: l,
		r: r,
	}
}

func (us *TrainingGroupUseCase) CreateGroup(ctx context.Context, g entity.TrainingGroup) (uuid.UUID, error) {
	const op = "training.CreateGroup"
	id, err := us.r.CreateGroup(ctx, g)
	if err != nil {
		us.l.Error("Failed to create group %v (op=%v)", err, op)
	}
	return id, err
}

func (us *TrainingGroupUseCase) ListGroups(ctx context.Context) ([]entity.TrainingGroup, error) {
	const op = "training.ListGroups"
	groups, err := us.r.ListGroups(ctx)
	if err != nil {
		us.l.Error("Failed to list groups %v (op=%v)", err, op)
		return nil, err
	}
	return groups, nil
}

func (us *TrainingGroupUseCase) UpdateGroup(ctx context.Context, g entity.TrainingGroup, id uuid.UUID) error {
	const op = "training.UpdateGroup"
	if err := us.r.UpdateGroup(ctx, g, id); err != nil {
		us.l.Error("Failed to update group %v (op=%v)", err, op)
		return err
	}
	return nil
}

func (us *TrainingGroupUseCase) DeleteGroup(ctx context.Context, id uuid.UUID) error {
	const op = "training.DeleteGroup"
	if err := us.r.DeleteGroup(ctx, id); err != nil {
		us.l.Error("Failed to delete group %v (op=%v)", err, op)
		return err
	}
	return nil
}

func (us *TrainingGroupUseCase) GetGroupByName(ctx context.Context, name string) (entity.TrainingGroup, error) {
	const op = "training.CreateGroup"
	group, err := us.r.GetGroupByName(ctx, name)
	if err != nil {
		us.l.Error("Failed to get group by name %v (op=%v)", err, op)
		return entity.TrainingGroup{}, err
	}
	return group, nil
}
