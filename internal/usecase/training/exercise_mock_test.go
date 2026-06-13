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

func Test_CreateExercise(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	exRepo := repomocks.NewMockExerciseRepo(ctrl)
	exID := uuid.New()

	input := entity.Exercise{
		Muscle:      "Спина",
		Position:    "Стоя",
		Description: "Подтягивания широким хватом, 4x8-10",
	}

	exRepo.EXPECT().CreateExercise(gomock.Any(), input).Return(exID, nil)

	uc := NewExerciseUseCase(testLogger(), exRepo)
	gotID, err := uc.CreateExercise(context.Background(), input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotID != exID {
		t.Errorf("got id %v, want %v", gotID, exID)
	}
}

func Test_UpdateExercise(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	exRepo := repomocks.NewMockExerciseRepo(ctrl)
	exID := uuid.New()

	newDescription := "Тяга штанги в наклоне, 4x10"
	updateInput := dto.UpdateExerciseInput{
		Description: &newDescription,
	}

	exRepo.EXPECT().UpdateExercise(gomock.Any(), updateInput, exID).Return(nil)

	uc := NewExerciseUseCase(testLogger(), exRepo)

	err := uc.UpdateExercise(context.Background(), updateInput, exID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
func Test_UpdateExercise_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	exRepo := repomocks.NewMockExerciseRepo(ctrl)
	exID := uuid.New()
	desc := "x"
	input := dto.UpdateExerciseInput{Description: &desc}

	exRepo.EXPECT().UpdateExercise(gomock.Any(), input, exID).Return(repo.ErrNotFound)

	uc := NewExerciseUseCase(testLogger(), exRepo)
	err := uc.UpdateExercise(context.Background(), input, exID)
	if !errors.Is(err, repo.ErrNotFound) {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}

func Test_DeleteExercise(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	exRepo := repomocks.NewMockExerciseRepo(ctrl)
	exID := uuid.New()

	exRepo.EXPECT().DeleteExercise(gomock.Any(), exID).Return(nil)

	uc := NewExerciseUseCase(testLogger(), exRepo)

	if err := uc.DeleteExercise(context.Background(), exID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func Test_DeleteExercise_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	exRepo := repomocks.NewMockExerciseRepo(ctrl)
	exID := uuid.New()

	exRepo.EXPECT().DeleteExercise(gomock.Any(), exID).Return(repo.ErrNotFound)

	uc := NewExerciseUseCase(testLogger(), exRepo)

	err := uc.DeleteExercise(context.Background(), exID)
	if !errors.Is(err, repo.ErrNotFound) {
		t.Errorf("expected ErrNotFaound got %v", err)
	}
}
