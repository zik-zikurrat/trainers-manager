package training

import (
	"context"
	"fmt"

	"trainers-manager/internal/controller/restapi/v1/request"
	"trainers-manager/internal/entity"
	"trainers-manager/internal/usecase"

	"github.com/google/uuid"
)

const _limitPlans = 4

func (us *UseCase) Generate(ctx context.Context, in request.Generate) (entity.TrainingPlan, error) {
	const op = "training.Generate"

	group, err := us.repo.GetGroupByName(ctx, in.TrainType)
	if err != nil {
		us.log.Error("group lookup failed", err, op)
		return entity.TrainingPlan{}, err
	}

	recent, err := us.repo.RecentPlans(ctx, group.ID, _limitPlans)
	if err != nil {
		us.log.Error("recent plans failed", err, op)
		return entity.TrainingPlan{}, err
	}

	var lastAccent, lastSkills string
	if len(recent) > 0 {
		lastAccent = recent[0].Accent
		lastSkills = recent[0].Skills
	}

	accent := nextInCycle(group.AccentCycle, lastAccent)
	skills := nextInCycle(group.SkillCycle, lastSkills)

	structure, err := us.repo.GetStructure(ctx, in.StructureID)
	if err != nil {
		us.log.Error("structure lookup failed", err, op)
		return entity.TrainingPlan{}, err
	}
	pool, err := us.repo.ListExercises(ctx)
	if err != nil {
		us.log.Error("exercises pool failed", err, op)
		return entity.TrainingPlan{}, err
	}

	gen, err := us.gen.Generate(ctx, usecase.GeneratePrompt{
		Structure: structure.Structure,
		Accent:    accent,
		Skills:    skills,
		Recent:    recent,
		Pool:      pool,
	})
	if err != nil {
		us.log.Error("llm generate failed", err, op)
		return entity.TrainingPlan{}, err
	}

	validIDs, err := validateExerciseIDs(gen.ExerciseIDs, pool)
	if err != nil {
		us.log.Error("llm returned invalid exercises", err, op)
		return entity.TrainingPlan{}, err
	}

	trainID, err := us.repo.CreateTraining(ctx)
	if err != nil {
		us.log.Error("create training failed", err, op)
		return entity.TrainingPlan{}, err
	}
	if err := us.linkExercises(ctx, trainID, validIDs); err != nil {
		us.log.Error("link exercises failed", err, op)
		return entity.TrainingPlan{}, err
	}

	plan := entity.TrainingPlan{
		Plan:                gen.PlanText,
		TrainID:             trainID,
		GroupID:             group.ID,
		Accent:              accent,
		Skills:              skills,
		TrainingStructureID: in.StructureID,
	}
	planID, err := us.repo.StoreTrainingPlan(ctx, plan)
	if err != nil {
		us.log.Error("store plan failed", err, op)
		return entity.TrainingPlan{}, err
	}
	plan.ID = planID

	return plan, nil
}

func validateExerciseIDs(got []uuid.UUID, pool []entity.Exercise) ([]uuid.UUID, error) {
	allowed := make(map[uuid.UUID]struct{}, len(pool))
	for _, e := range pool {
		allowed[e.ID] = struct{}{}
	}
	out := make([]uuid.UUID, 0, len(got))
	for _, id := range got {
		if _, ok := allowed[id]; !ok {
			return nil, fmt.Errorf("llm returned id not in pool: %s", id)
		}
		out = append(out, id)
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("llm returned no valid exercises")
	}
	return out, nil
}

func nextInCycle(cycle []string, last string) string {
	if len(cycle) == 0 {
		return ""
	}
	for i, v := range cycle {
		if v == last {
			return cycle[(i+1)%len(cycle)]
		}
	}
	return cycle[0]
}

func (us *UseCase) linkExercises(ctx context.Context, trainingID uuid.UUID, exerciseIDs []uuid.UUID) error {
	if err := us.repo.LinkExercises(ctx, trainingID, exerciseIDs); err != nil {
		us.log.Error("Failed to link exercises", err, "training.linkExercises")
		return err
	}
	return nil
}
