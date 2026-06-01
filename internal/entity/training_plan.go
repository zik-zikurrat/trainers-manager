package entity

import (
	"time"

	"github.com/google/uuid"
)

type TrainingPlan struct {
	ID                  uuid.UUID
	Plan                string
	TrainID             uuid.UUID
	Accent              string
	Skills              []string
	TrainingStructureID uuid.UUID
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
