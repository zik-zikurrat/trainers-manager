package entity

import (
	"time"

	"github.com/google/uuid"
)

// TrainingGroup -.
type TrainingGroup struct {
	ID           uuid.UUID
	accent_cycle []string
	skill_cycle  []string
	created_at   time.Time
	updated_at   time.Time
}
