package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebas1012/go-rest-api/config"
)

func Index(c *fiber.Ctx) error {
	p := config.Config("DB_PORT")

	return c.JSON(fiber.Map{"message": "Hi!", "env": p})
}
