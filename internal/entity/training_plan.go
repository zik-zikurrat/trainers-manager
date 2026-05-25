package entity

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

// TrainingPlan
type TrainingPlan struct {
	ID        uuid.UUID
	Plan      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
