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
	// 1. parse request body data into requestDto type
	var reqDto types.CreateUserRequest
	err := c.BodyParser(&reqDto)
	if err != nil {
		return err
	}

	// 2. convert requestDto type into domainEntity type to handle core logic
	userEntity, err := types.NewUserEntity(reqDto)
	if err != nil {
		return err
	}
	// core logic validation
	validationErr := userEntity.Validate()
	if validationErr != nil {
		return validationErr
	}

	// 3. convert domainEntity type into databaseDto type
	dbEntity, _ := types.NewMongoDbUserEntity(*userEntity)
	createdUser, err := uh.repo.CreateUser(c.Context(), (dbEntity))
	if err != nil {
		return err
	}

	// Now convert the returned domain entity into the response dto form
	userResponseDto := types.CreateUserResponse{
		ID:        createdUser.ID,
		FirstName: createdUser.FirstName,
		LastName:  createdUser.LastName,
		Email:     createdUser.Email,
	}
	// 4. returns json_response / or error if there is one
	return c.JSON(userResponseDto)
}
