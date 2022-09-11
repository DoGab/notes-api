package routes

import (
	"github.com/dogab/notes-api/app/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")
	// Create a Note
	note.Post("/", handlers.CreateNotes)
	// Read all Notes
	note.Get("/", handlers.GetNotes)
	// // Read one Note
	note.Get("/:noteId", handlers.GetNote)
	// // Update one Note
	note.Put("/:noteId", handlers.UpdateNote)
	// // Delete one Note
	note.Delete("/:noteId", handlers.DeleteNote)
}
