package training

import (
	"context"
	"errors"
	"testing"
	"trainers-manager/internal/entity"
	"trainers-manager/internal/repo"
	repomocks "trainers-manager/internal/repo/mocks"
	"trainers-manager/internal/usecase/dto"

	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
)

func Test_CreateTrainingGroup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	groupRepo := repomocks.NewMockTrainingGroupRepo(ctrl)
	groupID := uuid.New()

	input := entity.TrainingGroup{
		Name:        "upper body",
		AccentCycle: []string{"спина", "бицепс", "грудь", "трицепс", "плечи", "предплечья"},
		SkillCycle:  []string{"баланс, взрывная сила", "координация, выносливость"},
	}

	groupRepo.EXPECT().CreateGroup(gomock.Any(), input).Return(groupID, nil)

	uc := NewGroupUseCase(testLogger(), groupRepo)
	gotID, err := uc.CreateGroup(context.Background(), input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotID != groupID {
		t.Errorf("got id %v, want %v", gotID, groupID)
	}
}
func Test_CreateTrainingGroup_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	groupRepo := repomocks.NewMockTrainingGroupRepo(ctrl)
	input := entity.TrainingGroup{
		Name:        "upper body",
		AccentCycle: []string{"спина", "бицепс", "грудь", "трицепс", "плечи", "предплечья"},
		SkillCycle:  []string{"баланс, взрывная сила", "координация, выносливость"},
	}

	groupRepo.EXPECT().CreateGroup(gomock.Any(), input).Return(uuid.Nil, repo.ErrAlreadyExists)

	uc := NewGroupUseCase(testLogger(), groupRepo)
	_, err := uc.CreateGroup(context.Background(), input)
	if !errors.Is(err, repo.ErrAlreadyExists) {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}

func Test_UpdateTrainingGroup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	groupRepo := repomocks.NewMockTrainingGroupRepo(ctrl)
	groupID := uuid.New()
	name := "x"
	input := dto.UpdateGroupInput{ID: groupID, Name: &name}

	groupRepo.EXPECT().UpdateGroup(gomock.Any(), input).Return(nil)

	uc := NewGroupUseCase(testLogger(), groupRepo)
	err := uc.UpdateGroup(context.Background(), input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
func Test_UpdateTrainingGroup_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	groupRepo := repomocks.NewMockTrainingGroupRepo(ctrl)
	groupID := uuid.New()
	name := "x"
	input := dto.UpdateGroupInput{ID: groupID, Name: &name}

	groupRepo.EXPECT().UpdateGroup(gomock.Any(), input).Return(repo.ErrNotFound)

	uc := NewGroupUseCase(testLogger(), groupRepo)
	err := uc.UpdateGroup(context.Background(), input)
	if !errors.Is(err, repo.ErrNotFound) {
		t.Errorf("expected ErrNotFound, got %v", err)
	}

}
