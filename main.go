package main

import (
	"fmt"
	"time"

	"github.com/Balugrizzly/LogHubber/controllers"
	"github.com/Balugrizzly/LogHubber/database"
	"github.com/gofiber/fiber/v2"
)

func main() {

	fmt.Println(time.Now())

	database.Init()
	app := fiber.New()

	app.Get("/ping", controllers.Ping)

	app.Post("/log", controllers.Log)

	app.Listen(":3000")
}
