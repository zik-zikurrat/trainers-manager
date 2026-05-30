package persistent

import (
	"context"
	"fmt"
	"trainers-manager/internal/entity"
	"trainers-manager/pkg/postgres"
)

const _defaultEntityCap = 64

// TrainingRepo -.
type TrainingRepo struct {
	*postgres.Posgtres
}

// New -.
func New(pg *postgres.Posgtres) *TrainingRepo {
	return &TrainingRepo{pg}
}

// StoreStructure -.
func (r *TrainingRepo) StoreStructure(ctx context.Context, structure entity.TrainingStructure) error {
	_, err := r.Pool.Query(ctx, insertTrainingStructureQuery,
		structure.Structure,
	)
	if err != nil {
		return fmt.Errorf("insert training structure: %v", err)
	}
	return nil
}

// StoreTraining -.
// func (r *TrainingRepo) StoreTraining(context.Context, entity.Training) error{
// 	err := r.Pool.QueryRow(ctx, )
// }

// // StoreTrainingPlan -.
// func (r *TrainingRepo) StoreTrainingPlan(context.Context, []entity.TrainingPlan) error

// // GetHistory -.
// func (r *TrainingRepo) GetHistory(context.Context) ([]entity.TrainingPlanHistory, error)

// // GetTrainingPlan -.
// func (r *TrainingRepo) GetTrainingPlan(context.Context, time.Time) (entity.TrainingPlan, error)
