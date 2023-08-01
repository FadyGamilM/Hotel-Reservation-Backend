package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserDto struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	FirstName string             `bson:"first_name" json:"first_name"`
	LastName  string             `bson:"last_name" json:"last_name"`
}
