package training

import (
	"context"
	"trainers-manager/internal/repo"
	"trainers-manager/pkg/logger"

	"github.com/google/uuid"
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

func (us *UseCase) CreateTraining(ctx context.Context) (uuid.UUID, error) {
	id, err := us.repo.CreateTraining(ctx)
	if err != nil {
		us.log.Error("Failed to create training", err, "training.CreateTraining")
	}
	return id, err
}
