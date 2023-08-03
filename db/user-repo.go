package db

import (
	"github.com/FadyGamilM/hotelreservationapi/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUsers() ([]*types.User, error)
	CreateUser(*types.User) error
	GetUserById(string) (*types.User, error)
}

type MongoUserRepo struct {
	db *mongo.Client
}

func NewMongoUserRepo(client *mongo.Client) *MongoUserRepo {
	return &MongoUserRepo{
		db: client,
	}
}
