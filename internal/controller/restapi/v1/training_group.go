package v1

import (
	"errors"
	"net/http"
	"strings"

	"trainers-manager/internal/controller/restapi/v1/request"
	"trainers-manager/internal/entity"
	"trainers-manager/internal/repo"
	"trainers-manager/internal/usecase/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Summary     CreateGroup
// @Description Create group
// @ID          createGroup
// @Tags  	    createGroup
// @Accept      json
// @Produce     json
// @Success     201
// @Failure     500 {object} response.Error
// @Router      /training/group [post]
func (r *V1) CreateGroup(ctx *fiber.Ctx) error {
	var req request.TrainingGroup
	req.Name = strings.TrimSpace(req.Name)
	if err := ctx.BodyParser(&req); err != nil {
		r.l.Error(err, "restapi - v1 - group")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}
	if err := r.v.Struct(req); err != nil {
		r.l.Error(err, "restapi - v1 - group")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid request body")
	}

	id, err := r.group.CreateGroup(ctx.UserContext(), entity.TrainingGroup{
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

// @Summary     ListGroup
// @Description List group
// @ID          listGroup
// @Tags  	    lsitGroup
// @Accept      json
// @Produce     json
// @Success     200
// @Failure     500 {object} response.Error
// @Router      /training/group [get]
func (r *V1) ListGroups(ctx *fiber.Ctx) error {
	groups, err := r.group.ListGroups(ctx.UserContext())
	if err != nil {
		r.l.Error(err, "restapi - v1 - group")
		return errorResponse(ctx, http.StatusInternalServerError, "Group service error")
	}
	return ctx.Status(http.StatusOK).JSON(groups)
}

// @Summary     UpdateGroup
// @Description Update group
// @ID          UpdateGroup
// @Tags  	    updateGroup
// @Accept      json
// @Produce     json
// @Success     200
// @Failure     500 {object} response.Error
// @Router      /training/group/:id [patch]
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
	input := dto.UpdateGroupInput{
		ID:          uuidID,
		Name:        &req.Name,
		AccentCycle: &req.AccentCycle,
		SkillCycle:  &req.SkillCycle,
	}
	err = r.group.UpdateGroup(ctx.UserContext(), input)
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

// @Summary     DeleteGroup
// @Description Delete group
// @ID          DeleteGroup
// @Tags  	    deleteGroup
// @Accept      json
// @Produce     json
// @Success     204
// @Failure     500 {object} response.Error
// @Router      /training/group/:id [delete]
func (r *V1) DeleteGroup(ctx *fiber.Ctx) error {
	uuidID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		r.l.Error(err, "restapi - v1 - group")
		return errorResponse(ctx, http.StatusBadRequest, "Invalid id")
	}

	err = r.group.DeleteGroup(ctx.UserContext(), uuidID)
	switch {
	case errors.Is(err, repo.ErrNotFound):
		return errorResponse(ctx, http.StatusNotFound, "Group not found")
	case err != nil:
		r.l.Error(err, "restapi - v1 - group")
		return errorResponse(ctx, http.StatusInternalServerError, "Group service error")
	}

	return ctx.Status(http.StatusNoContent).Send(nil)
}
