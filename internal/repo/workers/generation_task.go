package workers

import (
	"time"

	"github.com/google/uuid"
)

type GenerationTask struct {
	ID        uuid.UUID
	Status    string
	Error     *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateGenerationTask struct {
	ID     uuid.UUID
	Status *string
	Error  *string
}
