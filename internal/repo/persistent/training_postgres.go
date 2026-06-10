package persistent

import (
	"context"
	"fmt"
	"trainers-manager/pkg/postgres"

	"github.com/google/uuid"
)

// TrainingRepo -.
type TrainingRepo struct {
	*postgres.Posgtres
}

// New -.
func NewTrainingRepo(pg *postgres.Posgtres) *TrainingRepo {
	return &TrainingRepo{pg}
}

func (r *TrainingRepo) CreateTraining(ctx context.Context) (uuid.UUID, error) {
	var id uuid.UUID
	if err := r.Pool.QueryRow(ctx, insertTrainingQuery).Scan(&id); err != nil {
		return uuid.Nil, fmt.Errorf("create training: %w", err)
	}
	return id, nil
}
