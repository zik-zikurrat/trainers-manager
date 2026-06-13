package training

import (
	"context"
	"testing"

	"trainers-manager/internal/entity"
	repomocks "trainers-manager/internal/repo/mocks"
	"trainers-manager/internal/usecase"
	ucmocks "trainers-manager/internal/usecase/mocks"
	"trainers-manager/pkg/logger"

	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
)

func testLogger() *logger.Logger {
	return logger.New("error")
}

func TestGenerate_FirstPlanUsesFirstAccent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	exID := uuid.New()
	groupID := uuid.New()
	structID := uuid.New()

	genRepo := repomocks.NewMockGenerationRepo(ctrl)
	planGen := ucmocks.NewMockPlanGenerator(ctrl)

	genRepo.EXPECT().GetGroupByName(gomock.Any(), "upper body").
		Return(entity.TrainingGroup{
			ID:          groupID,
			AccentCycle: []string{"спина, бицепс", "грудь, трицепс"},
			SkillCycle:  []string{"баланс", "выносливость"},
		}, nil)
	genRepo.EXPECT().RecentPlans(gomock.Any(), groupID, gomock.Any()).Return(nil, nil)
	genRepo.EXPECT().GetStructure(gomock.Any(), structID).
		Return(entity.TrainingStructure{Structure: "кардио, сила"}, nil)
	genRepo.EXPECT().ListExercises(gomock.Any()).
		Return([]entity.Exercise{{ID: exID, Muscle: "спина"}}, nil)

	planGen.EXPECT().Generate(gomock.Any(), gomock.Any()).
		Return(usecase.GeneratedPlan{ExerciseIDs: []uuid.UUID{exID}, PlanText: "план"}, nil)

	genRepo.EXPECT().CreateTraining(gomock.Any()).Return(uuid.New(), nil)
	genRepo.EXPECT().LinkExercises(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	genRepo.EXPECT().StoreTrainingPlan(gomock.Any(), gomock.Any()).Return(uuid.New(), nil)

	uc := NewGenerateUseCase(testLogger(), genRepo, planGen)

	plan, err := uc.Generate(context.Background(), "upper body", structID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if plan.Accent != "спина, бицепс" {
		t.Errorf("первый план = cycle[0], got %q", plan.Accent)
	}
}
