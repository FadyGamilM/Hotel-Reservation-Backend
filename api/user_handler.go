package api

import "github.com/gofiber/fiber/v2"

// handlers
func HandleReadiness(c *fiber.Ctx) error {
	return c.JSON("server is ready !")
}

func HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON("all users")
}

func HandleGetUserByID(c *fiber.Ctx) error {
	return c.JSON("fady")
}
