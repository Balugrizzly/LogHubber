package main

import (
	"github.com/Balugrizzly/LogHubber/controllers"
	"github.com/Balugrizzly/LogHubber/database"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Init()
	app := fiber.New()

	app.Get("/", controllers.Log)

	app.Listen(":3000")
}
