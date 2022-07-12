package userHandlers

import (
	"github.com/gofiber/fiber/v2"
	// connectdb "github.com/penkong/data4life/pkg/connect_db"
)

func GetUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"msg": "try get user with id",
	})
}