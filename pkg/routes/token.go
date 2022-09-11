package routes

import (
	"github.com/dogab/notes-api/app/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupTokenRoutes(app *fiber.App) {
	token := app.Group("/token")
	// Get a Token
	token.Post("/", handlers.GetToken)
	// Refresh a Token
	token.Get("/refresh", handlers.RefreshToken)
}
