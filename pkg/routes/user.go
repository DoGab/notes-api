package routes

import (
	"github.com/dogab/notes-api/app/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")

	// Create a User
	user.Post("/", handlers.CreateUser)
	// Read all Users
	user.Get("/", handlers.GetUsers)
	// Read one User
	user.Get("/:userId", handlers.GetUser)
	// Update one User
	user.Put("/:userId", handlers.UpdateUser)
	// // Delete one User
	// user.Delete("/:userId", userHandler.DeleteUser)
}
