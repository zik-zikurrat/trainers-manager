package entity

import (
	"time"

	"github.com/google/uuid"
)

// Exercise -.
type Exercise struct {
	ID          uuid.UUID
	Muscle      string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
