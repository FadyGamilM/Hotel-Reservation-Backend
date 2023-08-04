package api

import (
	"context"

	"github.com/FadyGamilM/hotelreservationapi/db"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	repo db.UserRepository
}

func NewUserHandler(r db.UserRepository) *UserHandler {
	return &UserHandler{
		repo: r,
	}
}

// handlers
func HandleReadiness(c *fiber.Ctx) error {
	return c.JSON("server is ready !")
}

func (uh *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON("all users")
}

func (uh *UserHandler) HandleGetUserByID(c *fiber.Ctx) error {
	var (
		user_id = c.Params("id")
		ctx     = context.Background()
	)

	user, err := uh.repo.GetUserById(ctx, user_id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}
