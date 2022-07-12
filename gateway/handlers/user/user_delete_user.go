package userHandlers

import (
	"github.com/gofiber/fiber/v2"
)

func DeleteUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"msg": "try delete user",
	})
}
