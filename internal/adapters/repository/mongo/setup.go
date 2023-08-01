package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Users_collection = "users-collection"

type MongoDB struct {
	Uri              string
	Client           *mongo.Client
	InfoLog          log.Logger
	ErrLog           log.Logger
	DbName           string
	CollectionsNames map[string]string
}

// uri := "mongodb://localhost:27017"

func NewMongoDB(dburi string) *MongoDB {
	m := &MongoDB{
		Uri:              dburi,
		DbName:           "hotel_reservation_db",
		CollectionsNames: make(map[string]string),
	}
	m.CollectionsNames[Users_collection] = "users"
	return m
}

func (m *MongoDB) ConnectAndGetClient() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Uri))
	if err != nil {
		m.ErrLog.Fatalf(" Error while trying to connect to mongodb instance : %v \n", err)
	}

	m.Client = client

}
