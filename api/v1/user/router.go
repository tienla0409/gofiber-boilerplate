package user

import (
	"github.com/tienla0409/gofiber-boilerplate/api/v1/share"
)

type userRouter struct {
	*share.ApiServer
}

func NewUserRouter(apiServer *share.ApiServer) *userRouter {
	return &userRouter{ApiServer: apiServer}
}

func (r *userRouter) RegisterRoutes() {
	handler := newUserHandler(r.ApiServer)
	router := r.Router.Group("/user")

	router.Get("/", handler.getUsers)
}
