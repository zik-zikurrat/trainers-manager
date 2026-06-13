package dto

import "github.com/google/uuid"

type UpdateExerciseInput struct {
	ID          uuid.UUID
	Muscle      *string
	Position    *string
	Description *string
}
