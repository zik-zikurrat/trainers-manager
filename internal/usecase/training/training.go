package training

import (
	"context"
	"trainers-manager/internal/entity"
	"trainers-manager/internal/repo"
	"trainers-manager/pkg/logger"
)

type UseCase struct {
	log  *logger.Logger
	repo repo.TrainingRepo
}

func New(r repo.TrainingRepo, log *logger.Logger) *UseCase {
	return &UseCase{
		repo: r,
		log:  log,
	}
}

// func (us *UseCase) GetHistory(ctx context.Context) ([]entity.TrainingPlanHistory, error) {
// 	const op = "training.History"
// 	log := us.log.With().Str("op", op)
// 	trainingPlanHistory, err := us.repo.GetHistory(ctx)
// 	if err != nil {
// 		log.Err(err)
// 		return nil, err
// 	}
// 	return trainingPlanHistory, nil
// }

func (us *UseCase) StoreStructure(ctx context.Context, structure entity.TrainingStructure) error {
	const op = "training.Structure"
	if err := us.repo.StoreStructure(ctx, structure); err != nil {
		us.log.Error("Failed to store structure", err, op)
		return err
	}
	return nil
}
