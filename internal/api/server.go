package api

import (
	"go-ecommerce-app/config"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	app.Get("/health", HealtCheck)

	app.Listen(config.ServerPort)
}

func HealtCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "I am breathng!",
	})
}
