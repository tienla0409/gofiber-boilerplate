package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tienla0409/gofiber-boilerplate/api/v1/share"
	"github.com/tienla0409/gofiber-boilerplate/api/v1/user"
	"github.com/tienla0409/gofiber-boilerplate/config"
	db "github.com/tienla0409/gofiber-boilerplate/db/sqlc"
	"github.com/tienla0409/gofiber-boilerplate/middleware"
	"github.com/tienla0409/gofiber-boilerplate/util"
	"log"
	"os"
)

type apiServerInternal struct {
	*share.ApiServer
	cleanupFunc func()
}

func NewApiServer(env *config.EnvConfig, poolConn *pgxpool.Pool) *apiServerInternal {
	instance := &apiServerInternal{
		ApiServer: &share.ApiServer{
			Env:     env,
			Queries: db.New(poolConn),
		},
	}

	instance.setupFiberServer()
	fileLogger := instance.setupLogger()

	instance.cleanupFunc = func() {
		if fileLogger != nil {
			err := fileLogger.Close()
			if err != nil {
				log.Printf("Failed to close file Logger: %v", err)
			}
		}
	}

	return instance
}

func (s *apiServerInternal) setupLogger() *os.File {
	fileLogger, err := config.InitLoggerConfig(s.Env.LoggerPath, "api")
	if err != nil {
		log.Panicf("Failed to create Logger: %v", err)
		return nil
	}

	return fileLogger
}

func (s *apiServerInternal) setupFiberServer() {
	s.App = fiber.New(fiber.Config{
		StructValidator: config.NewValidator(),
	})
	s.Router = s.App.Group("/api/v1")

	s.App.Use(middleware.NewLoggerMiddleware)

	user.NewUserRouter(s.ApiServer).RegisterRoutes()

	s.Router.Get(healthcheck.DefaultLivenessEndpoint, healthcheck.NewHealthChecker(healthcheck.Config{
		Probe: func(ctx fiber.Ctx) bool {
			return true
		},
	}))

	s.App.All("*", func(c fiber.Ctx) error {
		return util.SendError(c, util.Response{
			Status:  fiber.StatusNotFound,
			Message: "API endpoint not found.",
			Data:    nil,
		}, fmt.Sprintf("API endpoint not found: %s", c.Path()))
	})
}

func (s *apiServerInternal) Start() {
	defer s.cleanupFunc()

	log.Fatal(s.App.Listen(fmt.Sprintf(":%s", s.Env.Port)))
}
