package v1

import (
	"trainers-manager/internal/usecase"
	"trainers-manager/pkg/logger"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewTrainingRoutes(apiV1Group fiber.Router, t usecase.Training, l logger.Interface) {
	r := &V1{t: t, l: l, v: validator.New(validator.WithRequiredStructEnabled())}

	trainingGroup := apiV1Group.Group("/training")

	{
		trainingGroup.Get("/plan/history", r.History)
		trainingGroup.Post("/structure", r.Structure)
	}
}
