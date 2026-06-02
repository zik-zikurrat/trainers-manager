package entity

import (
	"time"

	"github.com/google/uuid"
)

type TrainingPlan struct {
	ID                  uuid.UUID
	Plan                string
	Status              string
	TrainID             uuid.UUID
	GroupID             uuid.UUID
	Accent              string
	Skills              string
	TrainingStructureID uuid.UUID
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
