package entity

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

// Training
type Training struct {
	ID        uuid.UUID
	Type      string
	Skills    []string
	Plan      TrainingPlan
	Structure TrainingStructure
	CreatedAt time.Time
	UpdatedAt time.Time
}
