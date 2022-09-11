package router

import (
	"github.com/dogab/notes-api/pkg/middleware"
	"github.com/dogab/notes-api/pkg/routes"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Use(logger.New())
	app.Use(cors.New())
	// app.Use(csrf.New())
	routes.SetupSwaggerRoutes(app)
	routes.SetupTokenRoutes(app)

	// Group api calls with param "/api"
	api := app.Group("/api")
	api.Use(middleware.JWTProtected())

	routes.SetupNoteRoutes(api)
	routes.SetupUserRoutes(api)

}
