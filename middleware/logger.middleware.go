package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/google/uuid"
	"github.com/tienla0409/gofiber-boilerplate/constant"
	"time"
)

func NewLoggerMiddleware(ctx fiber.Ctx) error {
	requestId := uuid.NewString()
	ctx.Locals(constant.RequestID, requestId)
	ctx.Locals(constant.Method, ctx.Method())
	ctx.Locals(constant.Path, ctx.Path())
	ctx.Locals(constant.SourceIP, ctx.IP())

	f := logger.New(logger.Config{
		Format:     "${time} ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: time.DateTime,
	})

	return f(ctx)

}
