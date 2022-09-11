package main

import (
	"github.com/dogab/notes-api/database"
	_ "github.com/dogab/notes-api/docs"
	"github.com/dogab/notes-api/pkg/config"
	"github.com/dogab/notes-api/pkg/router"
	"github.com/gofiber/fiber/v2"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	// Init envs
	config.InitEnvs()

	// Start a new fiber app
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	// setup the router
	router.SetupRoutes(app)

	// Listen on PORT 300
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
