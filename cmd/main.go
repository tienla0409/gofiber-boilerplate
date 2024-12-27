package main

import (
	v1 "github.com/tienla0409/gofiber-boilerplate/api/v1"
	"github.com/tienla0409/gofiber-boilerplate/config"
	"github.com/tienla0409/gofiber-boilerplate/db/sqlc"
	"log"
	"log/slog"
)

func main() {
	slog.Info("Loading env config...")
	envConfig, err := config.LoadEnvConfig(".")
	if err != nil {
		log.Panicf("Failed to load env config: %v", err)
	}
	slog.Info("Env config loaded")

	slog.Info("Connecting to database...")
	postgresConnector := db.NewPostgresConnector(envConfig.DbSource)
	poolConn, err := postgresConnector.Connect()
	if err != nil {
		log.Panicf("Failed to connect to database: %v", err)
	}
	defer poolConn.Close()
	slog.Info("Database connected")

	api := v1.NewApiServer(envConfig, poolConn)
	api.Start()

}
