package v1

import (
	"trainers-manager/internal/usecase"
	"trainers-manager/pkg/logger"
	"trainers-manager/pkg/workers"

	"github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
	training    usecase.Training
	exercise    usecase.Exercise
	structure   usecase.Structure
	plan        usecase.Plan
	planHistory usecase.PlanHistory
	group       usecase.Group
	generator   usecase.Generate

	l     logger.Interface
	v     *validator.Validate
	genCh chan workers.GenEvent
}
