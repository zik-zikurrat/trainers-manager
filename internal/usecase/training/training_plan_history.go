package training

import (
	"context"
	"trainers-manager/internal/entity"

	"github.com/google/uuid"
)

func (us *UseCase) GetPlanHistory(ctx context.Context, planID uuid.UUID) ([]entity.TrainingPlanHistory, error) {
	return us.repo.GetPlanHistory(ctx, planID)
}
