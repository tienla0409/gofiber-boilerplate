package user

import (
	"github.com/gofiber/fiber/v3"
	"github.com/tienla0409/gofiber-boilerplate/api/v1/share"
	"github.com/tienla0409/gofiber-boilerplate/util"
)

type userHandler struct {
	*share.ApiServer
}

func newUserHandler(apiServer *share.ApiServer) *userHandler {
	return &userHandler{ApiServer: apiServer}
}

func (h *userHandler) getUsers(ctx fiber.Ctx) error {
	users, err := h.Queries.GetUsers(ctx.Context())
	if err != nil {
		return util.SendError(ctx, util.Response{
			Status:  fiber.StatusInternalServerError,
			Message: "Failed to get user",
			Data:    nil,
		}, err.Error())
	}

	return util.SendSuccess(ctx, util.Response{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    users,
	})
}
