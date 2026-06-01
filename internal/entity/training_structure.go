package entity

import (
	"time"

	"github.com/google/uuid"
)

// TrainingStructure
type TrainingStructure struct {
	ID        uuid.UUID
	Structure string
	CreatedAt time.Time
	UpdatedAt time.Time
}
