package api

import (
	"github.com/FadyGamilM/hotelreservationapi/db"
	"github.com/FadyGamilM/hotelreservationapi/types"
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
func (uh *UserHandler) HandleGetUsers(c *fiber.Ctx) error {

	users, err := uh.repo.GetUsers(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(users)
}

func (uh *UserHandler) HandleGetUserByID(c *fiber.Ctx) error {
	var (
		user_id = c.Params("id")
	)

	user, err := uh.repo.GetUserById(c.Context(), user_id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

/*
@ responsibilites:

	➜ parse request body data into requestDto type
	➜ convert requestDto type into domainEntity type to handle core logic
	➜ convert domainEntity type into databaseDto type
	➜ call repo method and pass the databaseDto type

@ Returns:

	➜ returns json_response / or error if there is one
*/
func (uh *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	// 1. Get data from request body and parse it to the request dto
	var reqDto types.CreateUserRequest
	err := c.BodyParser(&reqDto)
	if err != nil {
		return err
	}

	// i should have a conversion between the requestDto => domain model entity
	// then the domain model entity will validate the data in the middle and return any errors if there are
	// then another conversion between the domain entity => db entity and pass it here ( i will do this in hex arch version of this project )
	userEntity, err := types.NewUserEntity(reqDto)
	if err != nil {
		return err
	}
	createdUser, err := uh.repo.CreateUser(c.Context(), userEntity)
	if err != nil {
		return err
	}

	// Now convert the returned data into the response dto form
	userResponseDto := types.CreateUserResponse{
		ID:        createdUser.ID,
		FirstName: createdUser.FirstName,
		LastName:  createdUser.LastName,
		Email:     createdUser.Email,
	}

	return c.JSON(userResponseDto)
}
