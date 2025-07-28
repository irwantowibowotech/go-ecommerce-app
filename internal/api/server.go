package api

import (
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/api/rest/handlers"
	"go-ecommerce-app/internal/domain"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection server error : %v", err)
	}

	log.Println("database connected")

	// run migrations
	db.AutoMigrate(&domain.User{})

	app.Get("/health", HealtCheck)

	restHandler := &rest.RestHandler{
		App: app,
		DB:  db,
	}

	setupRoutes(restHandler)

	app.Listen(config.ServerPort)
}

func HealtCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "I am breathng!",
	})
}

func setupRoutes(rh *rest.RestHandler) {
	// user handler
	handlers.SetupUserRoutes(rh)
	// transaction handler
	// catalog handler
}
