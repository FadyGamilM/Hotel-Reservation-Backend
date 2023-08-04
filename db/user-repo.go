package db

import (
	"context"

	"github.com/FadyGamilM/hotelreservationapi/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUsers(context.Context) ([]*types.User, error)
	CreateUser(context.Context, *types.User) (*types.User, error)
	GetUserById(context.Context, string) (*types.User, error)
}

type MongoUserRepo struct {
	db         *mongo.Client
	collection *mongo.Collection
}

const (
	DB_URI          = "mongodb://localhost:27017"
	USER_COLLECTION = "users"
)

func NewMongoUserRepo(client *mongo.Client) *MongoUserRepo {
	return &MongoUserRepo{
		db:         client,
		collection: client.Database(DB_NAME).Collection(USER_COLLECTION),
	}
}

func (m *MongoUserRepo) GetUsers(ctx context.Context) ([]*types.User, error) {
	users := []*types.User{}

	// define a cursor which is a pointer to the query result
	cur, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	// deserialize the result into our entity model and handle if there are any errors
	if err := cur.All(ctx, &users); err != nil {
		return []*types.User{}, nil
	}

	return users, nil
}

func (m *MongoUserRepo) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := m.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	// now set the id of the user document
	user.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return user, nil
}

func (m *MongoUserRepo) GetUserById(ctx context.Context, id string) (*types.User, error) {
	// define user entity to deserialize the data from mongodb to your entity model
	var user types.User
	// filter to get the user by its id
	// 1. first validate the id to be a right mongodb id
	obj_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = m.collection.FindOne(ctx, bson.M{"_id": obj_id}).Decode(&user)
	// handle errors
	if err != nil {
		return nil, err
	}
	return &user, nil
}
