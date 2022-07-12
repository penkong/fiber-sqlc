package userHandlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetCurrent(c *fiber.Ctx) error {
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"msg": "try get current user",
	})
}
