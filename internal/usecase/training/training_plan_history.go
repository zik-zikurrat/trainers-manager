package training

import (
	"context"
	"trainers-manager/internal/entity"
	"trainers-manager/internal/repo"
	"trainers-manager/pkg/logger"

	"github.com/google/uuid"
)

type PlanHistoryUseCase struct {
	l *logger.Logger
	r repo.PlanHistoryRepo
}

func NewPlanHistoryUseCase(l *logger.Logger, r repo.PlanHistoryRepo) *PlanHistoryUseCase {
	return &PlanHistoryUseCase{
		l: l,
		r: r,
	}
}

func (us *PlanHistoryUseCase) GetPlanHistory(ctx context.Context, planID uuid.UUID) ([]entity.TrainingPlanHistory, error) {
	const op = "training.GetPlanHistory"
	planHistory, err := us.r.GetPlanHistory(ctx, planID)
	if err != nil {
		us.l.Error("Failed to get plan history", err, op)
		return nil, err
	}
	return planHistory, nil
}
