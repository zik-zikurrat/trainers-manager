package v1

import (
	"errors"
	"net/http"

	"trainers-manager/internal/controller/restapi/v1/request"
	"trainers-manager/internal/entity"
	"trainers-manager/internal/repo"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (r *V1) CreateExercise(ctx *fiber.Ctx) error {
	var req request.Exercise
	if err := ctx.BodyParser(&req); err != nil {
		r.l.Error(err, "restapi - v1 - exercise")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}
	if err := r.v.Struct(req); err != nil {
		r.l.Error(err, "restapi - v1 - exercise")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}

	id, err := r.t.CreateExercise(ctx.UserContext(), entity.Exercise{
		MuscleGroup: req.MuscleGroup,
		Description: req.Description,
	})
	if err != nil {
		r.l.Error(err, "restapi - v1 - exercise")
		return errorResponse(ctx, http.StatusInternalServerError, "Exercise service error")
	}

	return ctx.Status(http.StatusCreated).JSON(map[string]string{"id": id.String()})
}

func (r *V1) ListExercises(ctx *fiber.Ctx) error {
	exercises, err := r.t.ListExercises(ctx.UserContext())
	if err != nil {
		r.l.Error(err, "restapi - v1 - exercise")
		return errorResponse(ctx, http.StatusInternalServerError, "Exercise service error")
	}
	return ctx.Status(http.StatusOK).JSON(exercises)
}

func (r *V1) UpdateExercise(ctx *fiber.Ctx) error {
	uuidID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		r.l.Error(err, "restapi - v1 - exercise")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid id")
	}

	var req request.Exercise
	if err := ctx.BodyParser(&req); err != nil {
		r.l.Error(err, "restapi - v1 - exercise")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}
	if err := r.v.Struct(req); err != nil {
		r.l.Error(err, "restapi - v1 - exercise")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}

	err = r.t.UpdateExercise(ctx.UserContext(), entity.Exercise{
		MuscleGroup: req.MuscleGroup,
		Description: req.Description,
	}, uuidID)
	switch {
	case errors.Is(err, repo.ErrNotFound):
		return errorResponse(ctx, http.StatusNotFound, "Exercise not found")
	case err != nil:
		r.l.Error(err, "restapi - v1 - exercise")
		return errorResponse(ctx, http.StatusInternalServerError, "Exercise service error")
	}

	return ctx.Status(http.StatusOK).JSON(map[string]string{"msg": "UPDATED"})
}

func (r *V1) DeleteExercise(ctx *fiber.Ctx) error {
	uuidID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		r.l.Error(err, "restapi - v1 - exercise")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid id")
	}

	err = r.t.DeleteExercise(ctx.UserContext(), uuidID)
	switch {
	case errors.Is(err, repo.ErrNotFound):
		return errorResponse(ctx, http.StatusNotFound, "Exercise not found")
	case err != nil:
		r.l.Error(err, "restapi - v1 - exercise")
		return errorResponse(ctx, http.StatusInternalServerError, "Exercise service error")
	}

	return ctx.Status(http.StatusNoContent).Send(nil)
}
