package api

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/FadyGamilM/hotelreservationapi/db"
	"github.com/FadyGamilM/hotelreservationapi/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	repo db.UserRepo
}

func NewUserHandler(r db.UserRepo) *UserHandler {
	return &UserHandler{
		repo: r,
	}
}

// handlers
func (uh *UserHandler) HandleGetUsers(c *fiber.Ctx) error {

	users, err := uh.repo.GetUsers()
	if err != nil {
		return err
	}

	responseDtos := []types.GetUserResponse{}

	for _, domainUser := range users {
		responseDtos = append(responseDtos, types.GetUserResponse{
			ID:        strconv.Itoa(int(domainUser.ID)),
			FirstName: domainUser.FirstName,
			LastName:  domainUser.LastName,
			Email:     domainUser.Email,
		})
	}

	return c.JSON(responseDtos)
}

func (uh *UserHandler) HandleGetUserByID(c *fiber.Ctx) error {
	var (
		id = c.Params("id")
	)

	user_id, err := strconv.ParseInt(id, 10, 64)

	user, err := uh.repo.GetUserById(user_id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("not found")
		}
		return err
	}

	// convert to responseDto
	userResponseDto := types.GetUserResponse{
		ID:        strconv.Itoa(int(user.ID)),
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
	userEntity, err := types.NewUserEntityFromUserRequestDto(reqDto)
	if err != nil {
		return err
	}
	// core logic validation
	validationErr := userEntity.Validate()
	if validationErr != nil {
		return validationErr
	}

	// 3. convert domainEntity type into databaseDto type
	createdUser, err := uh.repo.CreateUser(userEntity)
	if err != nil {
		return err
	}

	// Now convert the returned domain entity into the response dto form
	userResponseDto := types.CreateUserResponse{
		// ID:        strconv.Itoa(int(createdUser.ID)),
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
	id := c.Params("id")
	user_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errors.New("internal server error while deleting user from database")
	}

	err = uh.repo.DeleteUserById(user_id)
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
	id := c.Params("id")

	user_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errors.New("internal server error while deleting user from database")
	}

	updateRequestDto := new(types.UpdateUserRequest)
	err = c.BodyParser(&updateRequestDto)
	if err != nil {
		fmt.Println(updateRequestDto.FirstName)
		fmt.Println(err.Error())
		return errors.New("internal server error while deleting user from database")

	}

	// // check for allowed fields to be updated from the provided data to check if at least one of them is there
	// _, firstNameFieldExists := reflect.TypeOf(updateRequestDto).FieldByName("FirstName")
	// _, lastNameFieldExists := reflect.TypeOf(updateRequestDto).FieldByName("LastName")
	// if !firstNameFieldExists && !lastNameFieldExists {
	// 	return types.InvalidUpdateParameterErr{Msg: types.InvalidUpdateParameterMsg}
	// }

	updatedUser, err := uh.repo.UpdateUserById(user_id, updateRequestDto)
	updatedUserDto := new(types.UpdateUserResponse)
	updatedUserDto.FirstName = updatedUser.FirstName
	updatedUserDto.LastName = updatedUser.LastName
	if err != nil {
		return err
	}

	return c.JSON(updatedUserDto)
}
