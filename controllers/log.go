package controllers

import (
	"github.com/Balugrizzly/LogHubber/database"
	"github.com/Balugrizzly/LogHubber/models"
	"github.com/gofiber/fiber/v2"
)

func Log(c *fiber.Ctx) error {

	l := new(models.Log)

	if err := c.BodyParser(l); err != nil {
		return err
	}

	isValid, err := l.IsValid()
	if !isValid || err != nil {
		c.SendStatus(fiber.StatusBadRequest)
	}

	db := database.GetDB()
	result := db.Create(l)

	if result.Error != nil {
		panic(result.Error)
	}

	return c.SendStatus(fiber.StatusOK)
}
