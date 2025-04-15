package router

import (
	"github.com/Satr10/1cak-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Get("/", handlers.ApiIndex)

	api := app.Group("/api")
	api.Get("/", handlers.RandomMeme)
	api.Get("/random", handlers.RandomMeme)
	api.Get("/popular", handlers.PopularMeme)
}
