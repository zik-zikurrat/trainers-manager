package v1

import (
	"trainers-manager/internal/usecase"
	"trainers-manager/pkg/logger"

	"github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
	t usecase.Training
	l logger.Interface
	v *validator.Validate
}
