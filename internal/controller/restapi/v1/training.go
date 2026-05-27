package v1

import (
	"net/http"
	"trainers-manager/internal/controller/restapi/v1/request"
	"trainers-manager/internal/entity"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Show history
// @Description Show all training plan history
// @ID          history
// @Tags  	    training
// @Accept      json
// @Produce     json
// @Success     200 {object} entity.TrainingPlanHistory
// @Failure     500 {object} response.Error
// @Router      /training/history [get]
func (r *V1) History(ctx *fiber.Ctx) error {
	trainingHistory, err := r.t.History(ctx.UserContext())
	if err != nil {
		r.l.Error(err, "restapi - v1 - history")
		return errorResponse(ctx, http.StatusInternalServerError, "Failed to get training history")
	}
	return ctx.Status(http.StatusOK).JSON(trainingHistory)
}

// @Summary     Create structure
// @Description Create training structure
// @ID          trainingStructure
// @Tags  	    trainingStructure
// @Accept      json
// @Produce     json
// @Success     201
// @Failure     500 {object} response.Error
// @Router      /training/structure [post]
func (r *V1) Structure(ctx *fiber.Ctx) error {
	var req request.TrainingStructure

	if err := ctx.BodyParser(&req); err != nil {
		r.l.Error(err, "restapi - v1 - structure")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}

	if err := r.v.Struct(req); err != nil {
		r.l.Error(err, "restapi - v1 - structure")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}

	if err := r.t.Structure(
		ctx.UserContext(),
		entity.TrainingStructure{
			Structure: req.Structure,
		},
	); err != nil {
		r.l.Error(err, "restapi - v1 - structure")

		return errorResponse(ctx, http.StatusInternalServerError, "Structure service error")
	}

	return ctx.Status(http.StatusCreated).JSON(map[string]string{
		"msg": "CREATED",
	})

}
