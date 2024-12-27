package util

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mdobak/go-xerrors"
	"log/slog"
)

type Response struct {
	Status      int
	LogicStatus int
	Message     string
	Data        any
}

func SendSuccess(ctx fiber.Ctx, data Response) error {
	dataMap := fiber.Map{
		"data":    data.Data,
		"message": data.Message,
	}

	if data.LogicStatus == 0 {
		dataMap["status"] = data.Status
	} else {
		dataMap["status"] = data.LogicStatus
	}

	return ctx.Status(data.Status).JSON(dataMap)
}

func SendError(ctx fiber.Ctx, data Response, errLog string) error {
	slog.ErrorContext(ctx.Context(), "", xerrors.New(errLog))

	dataMap := fiber.Map{
		"data":    data.Data,
		"message": data.Message,
	}

	if data.LogicStatus == 0 {
		dataMap["status"] = data.Status
	} else {
		dataMap["status"] = data.LogicStatus
	}

	return ctx.Status(data.Status).JSON(dataMap)
}
