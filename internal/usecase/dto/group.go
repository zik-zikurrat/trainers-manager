package dto

import "github.com/google/uuid"

type UpdateGroupInput struct {
	ID          uuid.UUID
	Name        *string
	AccentCycle *[]string
	SkillCycle  *[]string
}
