package v1

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
	"trainers-manager/internal/controller/restapi/v1/request"
	"trainers-manager/internal/repo"
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
			errMsg := err.Error()
			r.genCh <- workers.GenEvent{
				TaskID: task_id,
				Status: "ERROR",
				Error:  &errMsg,
			}
		}
	}()

	return c.Status(http.StatusAccepted).JSON(map[string]string{"task_id": fmt.Sprintf("%v", task_id)})
}

func (r *V1) GetGenerationTask(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, "Invalid task id")
	}
	task, err := r.generationTask.GetGenerationTask(c.UserContext(), id)
	switch {
	case errors.Is(err, repo.ErrNotFound):
		return errorResponse(c, http.StatusNotFound, "Task not found")
	case err != nil:
		r.l.Error("GetGenerationTask: %v", err)
		return errorResponse(c, http.StatusInternalServerError, "Task service error")
	}
	return c.Status(http.StatusOK).JSON(map[string]string{"status": task.Status})
}
