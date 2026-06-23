package restapi

import (
	"trainers-manager/internal/config"
	"trainers-manager/internal/controller/restapi/middleware"
	v1 "trainers-manager/internal/controller/restapi/v1"
	"trainers-manager/internal/repo"
	"trainers-manager/internal/usecase"
	"trainers-manager/pkg/logger"
	"trainers-manager/pkg/workers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jackc/pgx/v5/pgxpool"
)

// NewRouter -.
// Swagger spec:
// @title       Trainers manager
// @description Create training plan
// @version     1.0
// @host        localhost:9045
// @BasePath    /v1
func NewRouter(
	app *fiber.App,
	cfg *config.Config,
	pool *pgxpool.Pool,
	l logger.Interface,
	training usecase.Training,
	exercise usecase.Exercise,
	structure usecase.Structure,
	plan usecase.Plan,
	planHistory usecase.PlanHistory,
	group usecase.Group,
	generator usecase.Generate,
	generationTask repo.GenerationTaskRepo,
	genCh chan workers.GenEvent,
) {

	app.Use(recover.New())
	// Prometheus metrics TODO
	// Swagger TODO
	// app.Get("/swagger/*", swagger.HandlerDefault)

	apiV1Group := app.Group("/v1")
	apiV1Group.Get("/health", func(c *fiber.Ctx) error {
		if err := pool.Ping(c.Context()); err != nil {
			return c.SendStatus(503)
		}
		return c.SendStatus(200)
	})
	{
		// Tracing
		apiV1Group.Use(middleware.TracingMiddleware())
		// Options
		apiV1Group.Use(middleware.LoggerMiddleware(l))
		// Cors
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
			generationTask,
			l,
			genCh,
		)
	}
}
