package share

import (
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
	"github.com/tienla0409/gofiber-boilerplate/config"
	db "github.com/tienla0409/gofiber-boilerplate/db/sqlc"
)

type ApiServer struct {
	App     *fiber.App
	Router  fiber.Router
	Env     *config.EnvConfig
	Logger  zerolog.Logger
	Queries *db.Queries
}
