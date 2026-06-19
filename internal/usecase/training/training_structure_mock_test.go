package training

import (
	"context"
	"errors"
	"testing"
	"trainers-manager/internal/entity"
	"trainers-manager/internal/repo"
	repomocks "trainers-manager/internal/repo/mocks"

	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
)

func Test_CreateStructure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	structureRepo := repomocks.NewMockTrainingStructureRepo(ctrl)

	input := entity.TrainingStructure{Structure: "some strucuture"}

	structureRepo.EXPECT().CreateStructure(gomock.Any(), input).Return(nil)

	uc := NewStructureUseCase(testLogger(), structureRepo)

	if err := uc.CreateStructure(context.Background(), input); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func Test_UpdateStructure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	structureRepo := repomocks.NewMockTrainingStructureRepo(ctrl)
	structureID := uuid.New()

	input := entity.TrainingStructure{Structure: "new structure"}

	structureRepo.EXPECT().UpdateStructure(gomock.Any(), input, structureID).Return(nil)

	uc := NewStructureUseCase(testLogger(), structureRepo)

	if err := uc.UpdateStructure(context.Background(), input, structureID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func Test_UpdateStructure_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	structureRepo := repomocks.NewMockTrainingStructureRepo(ctrl)
	structureID := uuid.New()

	input := entity.TrainingStructure{Structure: "new structure"}

	structureRepo.EXPECT().UpdateStructure(gomock.Any(), input, structureID).Return(repo.ErrNotFound)

	uc := NewStructureUseCase(testLogger(), structureRepo)

	err := uc.UpdateStructure(context.Background(), input, structureID)

	if !errors.Is(err, repo.ErrNotFound) {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}

func Test_DeleteStructure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	structureRepo := repomocks.NewMockTrainingStructureRepo(ctrl)
	structureID := uuid.New()

	structureRepo.EXPECT().DeleteStructure(gomock.Any(), structureID).Return(nil)

	uc := NewStructureUseCase(testLogger(), structureRepo)

	if err := uc.DeleteStructure(context.Background(), structureID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

}

func Test_DeleteStructure_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	structureRepo := repomocks.NewMockTrainingStructureRepo(ctrl)
	structureID := uuid.New()

	structureRepo.EXPECT().DeleteStructure(gomock.Any(), structureID).Return(repo.ErrNotFound)

	uc := NewStructureUseCase(testLogger(), structureRepo)

	err := uc.DeleteStructure(context.Background(), structureID)

	if !errors.Is(err, repo.ErrNotFound) {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}

func Test_GetStructure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	structureRepo := repomocks.NewMockTrainingStructureRepo(ctrl)
	structureID := uuid.New()

	structureRepo.EXPECT().GetStructure(gomock.Any(), structureID).Return(entity.TrainingStructure{}, nil)

	uc := NewStructureUseCase(testLogger(), structureRepo)

	if _, err := uc.GetStructure(context.Background(), structureID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

}
func Test_GetStructure_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	structureRepo := repomocks.NewMockTrainingStructureRepo(ctrl)
	structureID := uuid.New()

	structureRepo.EXPECT().GetStructure(gomock.Any(), structureID).Return(repo.ErrNotFound)

	uc := NewStructureUseCase(testLogger(), structureRepo)

	_, err := uc.GetStructure(context.Background(), structureID)

	if !errors.Is(err, repo.ErrNotFound) {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}
