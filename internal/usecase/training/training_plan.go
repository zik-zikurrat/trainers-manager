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


func (us *UseCase) Generate(ctx context.Context, in GenerateInput) (entity.TrainingPlan, error) {
	// 1. структура + недавние планы
	recent, _ := us.repo.RecentPlans(ctx, in.StructureID, 4)
	// 2. КОД решает акцент/навыки (детерминированно)
	accent, skills := nextRotation(recent)
	// 3. пул под акцент
	pool, _ := us.repo.ExercisesByMuscleGroup(ctx, musclesFor(accent))
	// 4. промпт -> LLM
	raw, _ := us.llm.Complete(ctx, systemPrompt, buildUserPrompt(accent, skills, recent, pool))
	// 5. парсинг + валидация id по pool
	exIDs, planText := parseAndValidate(raw, pool)
	// 6. персист (готовая транзакция): CreateTraining -> LinkExercises -> StoreTrainingPlan
	...
}

// Generate — ТВОЁ ядро. Вся обвязка ниже уже готова, тебе остаётся логика:
//  1. достать структуру (GetStructure по structureID)
//  2. us.repo.ListExercises(ctx) — пул упражнений
//  3. us.repo.GetPlanHistory — что было на прошлых неделях, чтобы крутить акцент/навыки
//  4. собрать accent + skills + отрендерить текст plan
//  5. CreateTraining -> LinkExercises -> StoreTrainingPlan (история запишется сама в tx)
