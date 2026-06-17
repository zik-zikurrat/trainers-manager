package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"trainers-manager/internal/controller/restapi/v1/request"
	"trainers-manager/pkg/workers"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	var req request.Generate
	if err := c.BodyParser(&req); err != nil {
		return errorResponse(c, http.StatusBadRequest, "Invalid request body")
	}
	if err := r.v.Struct(req); err != nil {
		return errorResponse(c, http.StatusBadRequest, "Invalid request body")
	}

	task_id := uuid.New()
	go func() {
		defer func() {
			if rec := recover(); rec != nil {
				r.l.Error("panic in generate: %v (op=%s)", rec, "background")
			}
		}()
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()
		r.genCh <- workers.GenEvent{
			TaskID: task_id,
			Status: "CREATED",
		}
		if _, err := r.generator.Generate(ctx, req.TrainType, req.StructureID, task_id); err != nil {
			r.l.Error("background generate failed: %v", err)
		}
	}()

	return c.Status(http.StatusAccepted).JSON(map[string]string{"task_id": fmt.Sprintf("%v", task_id)})
}
