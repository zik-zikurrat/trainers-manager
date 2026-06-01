package response

import (
	"time"
	"trainers-manager/internal/entity"
)

type TrainingStructure struct {
	Structure string    `json:"structure"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToTrainingStructure(structure entity.TrainingStructure) TrainingStructure {
	return TrainingStructure{
		Structure: structure.Structure,
		CreatedAt: structure.CreatedAt,
		UpdatedAt: structure.UpdatedAt,
	}
}
