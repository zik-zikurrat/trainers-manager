package entity

import (
	"time"

	"github.com/google/uuid"
)

type Training struct {
	ID        uuid.UUID
	Exercises []Exercise
	CreatedAt time.Time
	UpdatedAt time.Time
}
