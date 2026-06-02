package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"trainers-manager/internal/controller/restapi/v1/request"

	"github.com/gofiber/fiber/v2"
)

// @Summary     GenerateTrainingPlan
// @Description Generating training plan
// @ID          generateTrainingPlan
// @Tags  	    generateTrainingPlan
// @Accept      json
// @Produce     json
// @Success     202
// @Failure     500 {object} response.Error
// @Router      /training/generate [post]
func (r *V1) Generate(c *fiber.Ctx) error {
	// Validating
	var req request.Generate
	if err := c.BodyParser(&req); err != nil {
		r.l.Error(err, "restapi - v1 - generate")
		return errorResponse(c, http.StatusBadRequest, "Invalid request body")
	}
	if err := r.v.Struct(req); err != nil {
		r.l.Error(err, "restapi - v1 - generate")
		return errorResponse(c, http.StatusBadRequest, "Invalid request body")
	}

	// CreatePrompt
	prompt, err := r.g.CreatePrompt(c.UserContext(), req.TrainType, req.StructureID)
	if err != nil {
		r.l.Error(err, "restapi - v1 - generate")
		return errorResponse(c, http.StatusInternalServerError, "Invalid request body")
	}

	// Generate in goroutine
	// TODO: implement this with KAFKA
	go func() {
		defer func() {
			if rec := recover(); rec != nil {
				r.l.Error(fmt.Errorf("panic in generate: %v", rec), "background")
			}
		}()
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()

		if _, err := r.g.Generate(ctx, prompt); err != nil {
			r.l.Error(err, "background generate frailed")
		}
	}()

	return c.Status(http.StatusAccepted).JSON(map[string]string{"status": "generating"})
}
