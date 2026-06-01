package restapi

import (
	"trainers-manager/internal/config"
	"trainers-manager/internal/controller/restapi/middleware"
	v1 "trainers-manager/internal/controller/restapi/v1"
	"trainers-manager/internal/usecase"
	"trainers-manager/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

// NewRouter -.
// Swagger spec:
// @title       Trainers manager
// @description Create training plan
// @version     1.0
// @host        localhost:3033
// @BasePath    /v1
func NewRouter(app *fiber.App, cfg *config.Config, t usecase.Training, l logger.Interface) {
	// Options
	// app.Use(middleware.Logger(l))
	// app.Use(middleware.Recovery(l))

	// Prometheus metrics TODO

	// Swagger TODO
	// app.Get("/swagger/*", swagger.HandlerDefault)

	apiV1Group := app.Group("/v1")
	{
		apiV1Group.Use(middleware.TracingMiddleware())
		v1.NewTrainingRoutes(apiV1Group, t, l)
	}
}
