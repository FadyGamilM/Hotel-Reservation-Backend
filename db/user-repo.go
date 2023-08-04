package db

import (
	"context"

	"github.com/FadyGamilM/hotelreservationapi/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Notice that the repos must receive and return a user domain entity not the mongo entity type because handlers are sending these types and handlers should know nothing about the impl details ..
type UserRepository interface {
	GetUsers(context.Context) ([]*types.User, error)
	CreateUser(context.Context, *types.UserMongoDb) (*types.User, error)
	GetUserById(context.Context, string) (*types.User, error)
	DeleteUserById(context.Context, string) error
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
	domainUsers := []*types.User{}
	dbUsers := []*types.UserMongoDb{}

	// define a cursor which is a pointer to the query result
	cur, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	// deserialize the result into our entity model and handle if there are any errors
	if err := cur.All(ctx, &dbUsers); err != nil {
		return []*types.User{}, nil
	}

	// convert the db-related-users-type to domain entity type
	for _, dbUser := range dbUsers {
		domainUsers = append(domainUsers, &types.User{
			ID:                dbUser.ID,
			FirstName:         dbUser.FirstName,
			LastName:          dbUser.LastName,
			Email:             dbUser.Email,
			EncryptedPassword: dbUser.EncryptedPassword,
		})
	}

	return domainUsers, nil
}

func (m *MongoUserRepo) CreateUser(ctx context.Context, user *types.UserMongoDb) (*types.User, error) {
	res, err := m.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	// now set the id of the user document
	user.ID = res.InsertedID.(primitive.ObjectID).Hex()

	// return a user domain entity so the handlers don't have to convert from mongodb entity to a domain entity to make handlers more clean
	return &types.User{
		ID:                user.ID,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		Email:             user.Email,
		EncryptedPassword: user.EncryptedPassword,
	}, nil
}

func (m *MongoUserRepo) GetUserById(ctx context.Context, id string) (*types.User, error) {
	// filter to get the user by its id
	// 1. first validate the id to be a right mongodb id
	obj_id, err := convertFromStringToObjectID(id)
	if err != nil {
		return nil, err
	}

	// define user entity to deserialize the data from mongodb to your entity model
	var dbEntity types.UserMongoDb

	err = m.collection.FindOne(ctx, bson.M{"_id": obj_id}).Decode(&dbEntity)
	// handle errors
	if err != nil {
		return nil, err
	}

	// convert the dbEntity to domain entity and return it to tha handler
	domainEntity := types.User{
		ID:                dbEntity.ID,
		FirstName:         dbEntity.FirstName,
		LastName:          dbEntity.LastName,
		Email:             dbEntity.Email,
		EncryptedPassword: dbEntity.EncryptedPassword,
	}
	return &domainEntity, nil
}

func (m *MongoUserRepo) DeleteUserById(ctx context.Context, id string) error {
	// convert the id to a primitive objectID to filter based on it using mongo store
	obj_id, err := convertFromStringToObjectID(id)
	if err != nil {
		return err
	}

	_, err = m.collection.DeleteOne(ctx, bson.M{"_id": obj_id})
	if err != nil {
		return err
	}

	return nil
}

func convertFromStringToObjectID(id string) (*primitive.ObjectID, error) {
	obj_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return &obj_id, nil
}
