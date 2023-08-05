package api

import (
	"errors"
	"fmt"

	"github.com/FadyGamilM/hotelreservationapi/db"
	"github.com/FadyGamilM/hotelreservationapi/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
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

	responseDtos := []types.CreateUserResponse{}

	for _, domainUser := range users {
		responseDtos = append(responseDtos, types.CreateUserResponse{
			ID:        domainUser.ID,
			FirstName: domainUser.FirstName,
			LastName:  domainUser.LastName,
			Email:     domainUser.Email,
		})
	}

	return c.JSON(responseDtos)
}

func (uh *UserHandler) HandleGetUserByID(c *fiber.Ctx) error {
	var (
		user_id = c.Params("id")
	)

	user, err := uh.repo.GetUserById(c.Context(), user_id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("not found")
		}
		return err
	}

	// convert to responseDto
	userResponseDto := types.CreateUserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
	return c.JSON(userResponseDto)
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

/*
@ Responsibilites:

	➜ Get the user id from the request params
	➜ Call the repo delete  method to handle the request using the db

@ Returns:

	➜ Empty response or error if there is any
*/
func (uh *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	err := uh.repo.DeleteUserById(c.Context(), userID)
	if err != nil {
		return errors.New("internal server error while deleting user from database")
	}

	return nil
}

/*
@ Responsibilites:

	➜ Get the user id from the request params
	➜ Call the repo delete  method to handle the request using the db

@ Returns:

	➜ Empty response or error if there is any
*/
func (uh *UserHandler) HandleUpdateUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	updateRequestDto := new(types.UpdateUserRequest)
	err := c.BodyParser(&updateRequestDto)
	if err != nil {
		fmt.Println("ok")

		return err
	}
	fmt.Println(updateRequestDto.FirstName)

	// // check for allowed fields to be updated from the provided data to check if at least one of them is there
	// fieldExists, ok := reflect.TypeOf(updateRequestDto).FieldByName("FirstName")
	// if !ok {
	// 	fieldExists, ok = reflect
	// }

	err = uh.repo.UpdateUserById(c.Context(), userId, *updateRequestDto)
	if err != nil {
		return err
	}

	return c.JSON("updated")
}
