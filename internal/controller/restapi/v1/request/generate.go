package request

import "github.com/google/uuid"

type Generate struct {
	TrainType   string    `json:"train_type"   validate:"required"`
	StructureID uuid.UUID `json:"structure_id" validate:"required"`
}
