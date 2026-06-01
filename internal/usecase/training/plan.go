package training

import (
	"context"
	"time"

	"trainers-manager/internal/entity"

	"github.com/google/uuid"
)

type GenerateInput struct {
	StructureID uuid.UUID
	TrainID     uuid.UUID
	Date        time.Time
}

func (us *UseCase) CreateTraining(ctx context.Context) (uuid.UUID, error) {
	id, err := us.repo.CreateTraining(ctx)
	if err != nil {
		us.log.Error("Failed to create training", err, "training.CreateTraining")
	}
	return id, err
}

func (us *UseCase) LinkExercises(ctx context.Context, trainingID uuid.UUID, exerciseIDs []uuid.UUID) error {
	if err := us.repo.LinkExercises(ctx, trainingID, exerciseIDs); err != nil {
		us.log.Error("Failed to link exercises", err, "training.LinkExercises")
		return err
	}
	return nil
}

func (us *UseCase) StoreTrainingPlan(ctx context.Context, p entity.TrainingPlan) (uuid.UUID, error) {
	id, err := us.repo.StoreTrainingPlan(ctx, p)
	if err != nil {
		us.log.Error("Failed to store plan", err, "training.StoreTrainingPlan")
	}
	return id, err
}

func (us *UseCase) UpdateTrainingPlan(ctx context.Context, p entity.TrainingPlan, id uuid.UUID) error {
	if err := us.repo.UpdateTrainingPlan(ctx, p, id); err != nil {
		us.log.Error("Failed to update plan", err, "training.UpdateTrainingPlan")
		return err
	}
	return nil
}

func (us *UseCase) GetTrainingPlan(ctx context.Context, id uuid.UUID) (entity.TrainingPlan, error) {
	return us.repo.GetTrainingPlan(ctx, id)
}

func (us *UseCase) GetPlanHistory(ctx context.Context, planID uuid.UUID) ([]entity.TrainingPlanHistory, error) {
	return us.repo.GetPlanHistory(ctx, planID)
}

// func (us *UseCase) Generate(ctx context.Context, in GenerateInput) (entity.TrainingPlan, error) {
// 	structure, err := us.repo.GetStructure(ctx, in.StructureID)
// 	if err != nil {
// 		us.log.Error("Failed to get training structure", err)
// 		return entity.TrainingPlan{}, err
// 	}
// 	exercises := us.repo.LinkExercises(ctx, in.TrainID)
// }

// Generate — ТВОЁ ядро. Вся обвязка ниже уже готова, тебе остаётся логика:
//  1. достать структуру (GetStructure по structureID)
//  2. us.repo.ListExercises(ctx) — пул упражнений
//  3. us.repo.GetPlanHistory — что было на прошлых неделях, чтобы крутить акцент/навыки
//  4. собрать accent + skills + отрендерить текст plan
//  5. CreateTraining -> LinkExercises -> StoreTrainingPlan (история запишется сама в tx)
//
// func (us *UseCase) Generate(ctx context.Context, in GenerateInput) (entity.TrainingPlan, error) { ... }
