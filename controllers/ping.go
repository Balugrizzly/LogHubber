package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Ping(c *fiber.Ctx) error {

	payload := struct {
		Ping string `json:"Ping" params:"Ping"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	log.Printf("Body Payload is: %v", payload.Ping)

	return c.JSON(fiber.Map{
		"Ping": "Pong",
	})
}
