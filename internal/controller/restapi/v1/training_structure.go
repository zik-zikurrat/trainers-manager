package v1

import (
	"net/http"
	"trainers-manager/internal/controller/restapi/middleware"
	"trainers-manager/internal/controller/restapi/v1/request"
	"trainers-manager/internal/controller/restapi/v1/response"
	"trainers-manager/internal/entity"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Summary     Create structure
// @Description Create training structure
// @ID          trainingStructure
// @Tags  	    trainingStructure
// @Accept      json
// @Produce     json
// @Success     201
// @Failure     500 {object} response.Error
// @Router      /training/structure [post]
func (r *V1) CreateStructure(ctx *fiber.Ctx) error {
	var req request.TrainingStructure

	if err := ctx.BodyParser(&req); err != nil {
		r.l.Error(err, "restapi - v1 - structure")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}

	if err := r.v.Struct(req); err != nil {
		r.l.Error(err, "restapi - v1 - structure")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}

	if err := r.structure.CreateStructure(
		ctx.UserContext(),
		entity.TrainingStructure{
			Structure: req.Structure,
		},
	); err != nil {
		r.l.Error(err, "restapi - v1 - structure")

		return errorResponse(ctx, http.StatusInternalServerError, "Structure service error")
	}
	r.l.Info(
		"structure successfully created trace_id: %s",
		middleware.GetTraceID(ctx),
	)

	return ctx.Status(http.StatusCreated).JSON(map[string]string{
		"msg": "CREATED",
	})
}

// @Summary     Update structure
// @Description Update training structure
// @ID          trainingStructure
// @Tags  	    trainingStructure
// @Accept      json
// @Produce     json
// @Success     200
// @Failure     500 {object} response.Error
// @Router      /training/structure/:id [patch]
func (r *V1) UpdateStructure(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		r.l.Error(err, "restapi - v1 - structure")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid query params")
	}

	var req request.TrainingStructure

	if err := ctx.BodyParser(&req); err != nil {
		r.l.Error(err, "restapi - v1 - structure")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}

	if err := r.v.Struct(req); err != nil {
		r.l.Error(err, "restapi - v1 - structure")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}

	if err := r.structure.UpdateStructure(
		ctx.UserContext(),
		entity.TrainingStructure{
			Structure: req.Structure,
		},
		uuidID,
	); err != nil {
		r.l.Error(err, "restapi - v1 - structure")

		return errorResponse(ctx, http.StatusInternalServerError, "Structure service error")
	}

	return ctx.Status(http.StatusOK).JSON(map[string]string{
		"msg": "UPDATED",
	})
}

// @Summary     Delete structure
// @Description Delete training structure
// @ID          trainingStructure
// @Tags  	    trainingStructure
// @Accept      json
// @Produce     json
// @Success     204
// @Failure     500 {object} response.Error
// @Router      /training/structure/:id [patch]
func (r *V1) DeleteStructure(ctx *fiber.Ctx) error {
	uuidID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		r.l.Error(err, "restapi - v1 - structure")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid id")
	}
	if err := r.structure.DeleteStructure(ctx.UserContext(), uuidID); err != nil {
		r.l.Error(err, "restapi - v1 - structure")
		return errorResponse(ctx, http.StatusInternalServerError, "Structure service error")
	}
	return ctx.Status(http.StatusNoContent).Send(nil)
}

// @Summary     Get structure
// @Description Get training structure
// @ID          trainingStructure
// @Tags  	    trainingStructure
// @Accept      json
// @Produce     json
// @Success     200
// @Failure     500 {object} response.Error
// @Router      /training/structure/:id [get]
func (r *V1) GetStructure(ctx *fiber.Ctx) error {
	uuidID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		r.l.Error(err, "restapi - v1 - structure")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid id")
	}
	structure, err := r.structure.GetStructure(ctx.UserContext(), uuidID)
	if err != nil {
		r.l.Error(err, "restapi - v1 - structure")
		return errorResponse(ctx, http.StatusInternalServerError, "Error while getting structure")
	}
	resp := response.ToTrainingStructure(structure)
	return ctx.Status(http.StatusOK).JSON(resp)
}

// @Summary     List structure
// @Description List training structure
// @ID          trainingStructure
// @Tags  	    trainingStructure
// @Accept      json
// @Produce     json
// @Success     200
// @Failure     500 {object} response.Error
// @Router      /training/structure [get]
func (r *V1) ListStructure(ctx *fiber.Ctx) error {
	structures, err := r.structure.ListStructure(ctx.UserContext())
	if err != nil {
		r.l.Error(err, "restapi - v1 - structure")
		return errorResponse(ctx, http.StatusInternalServerError, "Error while getting structure")
	}
	resp := make(map[uuid.UUID]response.TrainingStructure, len(structures))
	for _, structure := range structures {
		resp[structure.ID] = response.ToTrainingStructure(structure)
	}
	return ctx.Status(http.StatusOK).JSON(resp)
}
