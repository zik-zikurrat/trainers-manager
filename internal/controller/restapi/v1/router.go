package v1

import (
	"trainers-manager/internal/usecase"
	"trainers-manager/pkg/logger"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewTrainingRoutes(
	apiV1Group fiber.Router,
	training usecase.Training,
	exercise usecase.Exercise,
	structure usecase.Structure,
	plan usecase.Plan,
	planHistory usecase.PlanHistory,
	group usecase.Group,
	generator usecase.Generate,
	l logger.Interface,
) {
	r := &V1{
		training:    training,
		exercise:    exercise,
		structure:   structure,
		plan:        plan,
		planHistory: planHistory,
		group:       group,
		generator:   generator,
		l:           l,
		v:           validator.New(validator.WithRequiredStructEnabled()),
	}
	trainingGroup := apiV1Group.Group("/training")

	{
		// structure
		trainingGroup.Post("/structure", r.CreateStructure)
		trainingGroup.Get("/structure", r.ListStructure)
		trainingGroup.Get("/structure/:id", r.GetStructure)
		trainingGroup.Patch("/structure/:id", r.UpdateStructure)
		trainingGroup.Delete("/structure/:id", r.DeleteStructure)
		// exercise
		trainingGroup.Post("/exercise", r.CreateExercise)
		trainingGroup.Get("/exercise", r.ListExercises)
		trainingGroup.Patch("/exercise/:id", r.UpdateExercise)
		trainingGroup.Delete("/exercise/:id", r.DeleteExercise)
		// group
		trainingGroup.Post("/group", r.CreateGroup)
		trainingGroup.Get("/group", r.ListGroups)
		trainingGroup.Patch("/group/:id", r.UpdateGroup)
		trainingGroup.Delete("/group/:id", r.DeleteGroup)
		// plan
		trainingGroup.Get("/plan", r.ListPlan)
		// generate
		trainingGroup.Post("/generate", r.Generate)
	}
}
