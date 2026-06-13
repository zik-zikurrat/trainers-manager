package training

import (
	"testing"
	"trainers-manager/internal/entity"

	"github.com/google/uuid"
)

func TestNextInCycle(t *testing.T) {
	cycle := []string{"спина", "грудь", "плечи"}
	tests := []struct {
		name string
		last string
		want string
	}{
		{"первый план — пустой last", "", "спина"},
		{"следующий по циклу", "спина", "грудь"},
		{"середина цикла", "грудь", "плечи"},
		{"закольцовка с последнего", "плечи", "спина"},
		{"last не из цикла → начало", "ноги", "спина"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextInCycle(cycle, tt.last); got != tt.want {
				t.Errorf("nextInCycle(last=%q) = %q, want %q", tt.last, got, tt.want)
			}
		})
	}
}

func TestNextInCycle_EmptyCycle(t *testing.T) {
	if got := nextInCycle(nil, ""); got != "" {
		t.Errorf("nextInCycle(last=%q) = %q, want %q", "", got, "")
	}
}

func TestValidateExerciseIDs(t *testing.T) {
	response := make([]uuid.UUID, 0, 3)
	poolID := make([]uuid.UUID, 0, 3)

	for i := 0; i < 3; i++ {
		id := uuid.New()
		poolID = append(poolID, id)
		response = append(response, id)
	}

	pool := []entity.Exercise{
		{ID: poolID[0]},
		{ID: poolID[1]},
		{ID: poolID[2]},
	}

	tests := []struct {
		name     string
		response []uuid.UUID
		pool     []entity.Exercise
		wantErr  bool
	}{
		{
			name:     "success - all ids valid",
			response: poolID,
			pool:     pool,
			wantErr:  false,
		},
		{
			name:     "fail - one invalid id",
			response: []uuid.UUID{uuid.New()},
			pool:     pool,
			wantErr:  true,
		},
		{
			name:     "fail - empty response",
			response: nil,
			pool:     pool,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := validateExerciseIDs(tt.response, tt.pool)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if len(got) != len(tt.response) {
				t.Fatalf("expected %d ids, got %d", len(tt.response), len(got))
			}
		})
	}
}
