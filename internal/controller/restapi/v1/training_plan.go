package v1

import (
	"net/http"
	"trainers-manager/internal/controller/restapi/v1/response"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Summary     List plan
// @Description List training plan
// @ID          trainingPlan
// @Tags  	    trainingPlan
// @Accept      json
// @Produce     json
// @Success     200
// @Failure     500 {object} response.Error
// @Router      /training/plan [get]
func (r *V1) ListPlan(ctx *fiber.Ctx) error {
	plans, err := r.plan.ListTrainingPlan(ctx.UserContext())
	if err != nil {
		r.l.Error(err, "restapi - v1 - plan")
		return errorResponse(ctx, http.StatusInternalServerError, "Error while getting plan")
	}
	resp := make(map[uuid.UUID]response.TrainingPlan, len(plans))
	for _, plan := range plans {
		resp[plan.ID] = response.ToTrainingPlan(plan)
	}
	return ctx.Status(http.StatusOK).JSON(resp)
}

// @Summary     Get plan
// @Description Getn training plan
// @ID          trainingPlan
// @Tags  	    trainingPlan
// @Accept      json
// @Produce     json
// @Success     200
// @Failure     500 {object} response.Error
// @Router      /training/plan/:id [get]
func (r *V1) GetPlan(ctx *fiber.Ctx) error {
	uuidID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		r.l.Error(err, "restapi - v1 - plan")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid id")
	}
	plan, err := r.plan.GetTrainingPlan(ctx.UserContext(), uuidID)
	if err != nil {
		r.l.Error(err, "restapi - v1 - plan")
		return errorResponse(ctx, http.StatusInternalServerError, "Error while getting plan")
	}
	resp := response.ToTrainingPlan(plan)
	return ctx.Status(http.StatusOK).JSON(resp)
}
