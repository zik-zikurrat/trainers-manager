package training

import (
	"context"
	"fmt"
	"trainers-manager/internal/entity"
	"trainers-manager/internal/repo"
	"trainers-manager/internal/usecase"
	"trainers-manager/pkg/logger"

	"github.com/google/uuid"
)

const _limitPlans = 4

type GenerateUseCase struct {
	l   *logger.Logger
	r   repo.GenerationRepo
	gen usecase.PlanGenerator
}

func NewGenerateUseCase(l *logger.Logger, r repo.GenerationRepo, gen usecase.PlanGenerator) *GenerateUseCase {
	return &GenerateUseCase{l: l, r: r, gen: gen}
}

func (us *GenerateUseCase) Generate(ctx context.Context, trainType string, structureID uuid.UUID) (entity.TrainingPlan, error) {
	const op = "training.Generate"

	prompt, group, err := us.buildPrompt(ctx, trainType, structureID)
	if err != nil {
		return entity.TrainingPlan{}, err
	}

	gen, err := us.gen.Generate(ctx, prompt)
	if err != nil {
		us.l.Error("llm generate failed %v (op=%v)", err, op)
		return entity.TrainingPlan{}, err
	}

	validIDs, err := validateExerciseIDs(gen.ExerciseIDs, prompt.Pool)
	if err != nil {
		us.l.Error("llm returned invalid exercises %v (op=%v)", err, op)
		return entity.TrainingPlan{}, err
	}

	trainID, err := us.r.CreateTraining(ctx)
	if err != nil {
		us.l.Error("create training failed %v (op=%v)", err, op)
		return entity.TrainingPlan{}, err
	}
	if err := us.r.LinkExercises(ctx, trainID, validIDs); err != nil {
		return entity.TrainingPlan{}, err
	}

	plan := entity.TrainingPlan{
		Plan:                gen.PlanText,
		TrainID:             trainID,
		Status:              "ACTIVE",
		GroupID:             group.ID,
		Accent:              prompt.Accent,
		Skills:              prompt.Skills,
		TrainingStructureID: structureID,
	}
	planID, err := us.r.StoreTrainingPlan(ctx, plan)
	if err != nil {
		us.l.Error("store plan failed %v (op=%v)", err, op)
		return entity.TrainingPlan{}, err
	}
	plan.ID = planID
	return plan, nil
}

func (us *GenerateUseCase) buildPrompt(ctx context.Context, trainType string, structureID uuid.UUID) (usecase.GeneratePrompt, entity.TrainingGroup, error) {
	const op = "training.buildPrompt"

	group, err := us.r.GetGroupByName(ctx, trainType)
	if err != nil {
		us.l.Error("group lookup failed %v (op=%v)", err, op)
		return usecase.GeneratePrompt{}, entity.TrainingGroup{}, err
	}

	recent, err := us.r.RecentPlans(ctx, group.ID, _limitPlans)
	if err != nil {
		us.l.Error("recent plans failed %v (op=%v)", err, op)
		return usecase.GeneratePrompt{}, entity.TrainingGroup{}, err
	}

	var lastAccent, lastSkills string
	if len(recent) > 0 {
		lastAccent = recent[0].Accent
		lastSkills = recent[0].Skills
	}
	accent := nextInCycle(group.AccentCycle, lastAccent)
	skills := nextInCycle(group.SkillCycle, lastSkills)

	structure, err := us.r.GetStructure(ctx, structureID)
	if err != nil {
		us.l.Error("structure lookup failed %v (op=%v)", err, op)
		return usecase.GeneratePrompt{}, entity.TrainingGroup{}, err
	}

	pool, err := us.r.ListExercises(ctx)
	if err != nil {
		us.l.Error("exercises pool failed %v (op=%v)", err, op)
		return usecase.GeneratePrompt{}, entity.TrainingGroup{}, err
	}

	return usecase.GeneratePrompt{
		Structure: structure.Structure,
		Accent:    accent,
		Skills:    skills,
		Recent:    recent,
		Pool:      pool,
	}, group, nil
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
