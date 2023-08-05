package api

import "github.com/gofiber/fiber/v2"

type CustomApiError struct {
	error string
}

var Config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(CustomApiError{
			error: err.Error(),
		})
	},
}
