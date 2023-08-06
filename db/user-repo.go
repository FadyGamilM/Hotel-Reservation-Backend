package db

import "github.com/FadyGamilM/hotelreservationapi/types"

type UserRepo interface {
	GetUsers() ([]*types.User, error)
	CreateUser(*types.User) (*types.User, error)
	GetUserById(int64) (*types.User, error)
	DeleteUserById(int64) error
	UpdateUserById(int64, *types.UpdateUserRequest) (*types.User, error)
}
