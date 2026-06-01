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

func (r *V1) CreateGroup(ctx *fiber.Ctx) error {
	var req request.TrainingGroup
	if err := ctx.BodyParser(&req); err != nil {
		r.l.Error(err, "restapi - v1 - group")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}
	if err := r.v.Struct(req); err != nil {
		r.l.Error(err, "restapi - v1 - group")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}

	id, err := r.t.CreateGroup(ctx.UserContext(), entity.TrainingGroup{
		Name:        req.Name,
		AccentCycle: req.AccentCycle,
		SkillCycle:  req.SkillCycle,
	})
	switch {
	case errors.Is(err, repo.ErrAlreadyExists):
		return errorResponse(ctx, http.StatusConflict, "Group with this name already exists")
	case err != nil:
		r.l.Error(err, "restapi - v1 - group")
		return errorResponse(ctx, http.StatusInternalServerError, "Group service error")
	}

	return ctx.Status(http.StatusCreated).JSON(map[string]string{"id": id.String()})
}

func (r *V1) ListGroups(ctx *fiber.Ctx) error {
	groups, err := r.t.ListGroups(ctx.UserContext())
	if err != nil {
		r.l.Error(err, "restapi - v1 - group")
		return errorResponse(ctx, http.StatusInternalServerError, "Group service error")
	}
	return ctx.Status(http.StatusOK).JSON(groups)
}

func (r *V1) UpdateGroup(ctx *fiber.Ctx) error {
	uuidID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		r.l.Error(err, "restapi - v1 - group")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid id")
	}

	var req request.TrainingGroup
	if err := ctx.BodyParser(&req); err != nil {
		r.l.Error(err, "restapi - v1 - group")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}
	if err := r.v.Struct(req); err != nil {
		r.l.Error(err, "restapi - v1 - group")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}

	err = r.t.UpdateGroup(ctx.UserContext(), entity.TrainingGroup{
		Name:        req.Name,
		AccentCycle: req.AccentCycle,
		SkillCycle:  req.SkillCycle,
	}, uuidID)
	switch {
	case errors.Is(err, repo.ErrNotFound):
		return errorResponse(ctx, http.StatusNotFound, "Group not found")
	case errors.Is(err, repo.ErrAlreadyExists):
		return errorResponse(ctx, http.StatusConflict, "Group with this name already exists")
	case err != nil:
		r.l.Error(err, "restapi - v1 - group")
		return errorResponse(ctx, http.StatusInternalServerError, "Group service error")
	}

	return ctx.Status(http.StatusOK).JSON(map[string]string{"msg": "UPDATED"})
}

func (r *V1) DeleteGroup(ctx *fiber.Ctx) error {
	uuidID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		r.l.Error(err, "restapi - v1 - group")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid id")
	}

	err = r.t.DeleteGroup(ctx.UserContext(), uuidID)
	switch {
	case errors.Is(err, repo.ErrNotFound):
		return errorResponse(ctx, http.StatusNotFound, "Group not found")
	case err != nil:
		r.l.Error(err, "restapi - v1 - group")
		return errorResponse(ctx, http.StatusInternalServerError, "Group service error")
	}

	return ctx.Status(http.StatusNoContent).Send(nil)
}
