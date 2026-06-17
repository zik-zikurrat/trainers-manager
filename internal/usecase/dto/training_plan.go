package dto

import (
	"time"

	"github.com/google/uuid"
)

type UpdateTrainingPlan struct {
	ID                  uuid.UUID
	Plan                *string
	Accent              *string
	Skills              *string
	TrainingStructureID uuid.UUID
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
