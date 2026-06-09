package training

import (
	"context"

	"trainers-manager/internal/entity"
	"trainers-manager/pkg/logger"

	"trainers-manager/internal/repo"

	"github.com/google/uuid"
)

type PlanUseCase struct {
	l *logger.Logger
	r repo.TrainingPlanRepo
}

func NewPlanUseCase(r repo.TrainingPlanRepo, l *logger.Logger) *PlanUseCase {
	return &PlanUseCase{
		l: l,
		r: r,
	}
}

func (us *PlanUseCase) StoreTrainingPlan(ctx context.Context, p entity.TrainingPlan) (uuid.UUID, error) {
	const op = "training.StoreTrainingPlan"
	id, err := us.r.StoreTrainingPlan(ctx, p)
	if err != nil {
		us.l.Error("Failed to store plan", err, op)
	}
	return id, err
}

func (us *PlanUseCase) UpdateTrainingPlan(ctx context.Context, p entity.TrainingPlan, id uuid.UUID) error {
	const op = "training.UpdateTrainingPlan"
	if err := us.r.UpdateTrainingPlan(ctx, p, id); err != nil {
		us.l.Error("Failed to update plan", err, op)
		return err
	}
	return nil
}

func (us *PlanUseCase) GetTrainingPlan(ctx context.Context, id uuid.UUID) (entity.TrainingPlan, error) {
	const op = "training.GetTrainingPlan"
	plan, err := us.r.GetTrainingPlan(ctx, id)
	if err != nil {
		us.l.Error("Failed to get plan", err, op)
	}
	return plan, nil
}

func (us *PlanUseCase) ListTrainingPlan(ctx context.Context) ([]entity.TrainingPlan, error) {
	const op = "training.ListTrainingPlan"
	plans, err := us.r.ListTrainingPlan(ctx)
	if err != nil {
		us.l.Error("Failed to list plans", err, op)
	}
	return plans, nil
}
