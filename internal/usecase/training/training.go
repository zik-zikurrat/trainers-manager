package training

import (
	"context"

	"trainers-manager/internal/repo"
	"trainers-manager/pkg/logger"

	"github.com/google/uuid"
)

type TrainingUseCase struct {
	l *logger.Logger
	r repo.TrainingRepo
}

func NewTrainingUseCase(r repo.TrainingRepo, l *logger.Logger) *TrainingUseCase {
	return &TrainingUseCase{
		r: r,
		l: l,
	}
}

func (us *TrainingUseCase) CreateTraining(ctx context.Context) (uuid.UUID, error) {
	id, err := us.r.CreateTraining(ctx)
	if err != nil {
		us.l.Error("Failed to create training", err, "training.CreateTraining")
	}
	return id, err
}
