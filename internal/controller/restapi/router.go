package restapi

import (
	"trainers-manager/internal/config"
	"trainers-manager/internal/controller/restapi/middleware"
	v1 "trainers-manager/internal/controller/restapi/v1"
	"trainers-manager/internal/usecase"
	"trainers-manager/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// NewRouter -.
// Swagger spec:
// @title       Trainers manager
// @description Create training plan
// @version     1.0
// @host        localhost:3033
// @BasePath    /v1
func NewRouter(
	app *fiber.App,
	cfg *config.Config,
	training usecase.Training,
	exercise usecase.Exercise,
	structure usecase.Structure,
	plan usecase.Plan,
	planHistory usecase.PlanHistory,
	group usecase.Group,
	generator usecase.Generate,
	l logger.Interface,
) {
	// Options
	// app.Use(middleware.Logger(l))
	// app.Use(middleware.Recovery(l))

	// Prometheus metrics TODO
	// Swagger TODO
	// app.Get("/swagger/*", swagger.HandlerDefault)

	apiV1Group := app.Group("/v1")
	{
		apiV1Group.Use(middleware.TracingMiddleware())
		apiV1Group.Use(cors.New())
		v1.NewTrainingRoutes(
			apiV1Group,
			training,
			exercise,
			structure,
			plan,
			planHistory,
			group,
			generator,
			l,
		)
	}
}
