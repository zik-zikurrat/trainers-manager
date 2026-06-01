package entity

import (
	"github.com/google/uuid"
)

// TrainingGroup -.
type TrainingGroup struct {
	ID          uuid.UUID
	Name        string
	AccentCycle []string
	SkillCycle  []string
}
