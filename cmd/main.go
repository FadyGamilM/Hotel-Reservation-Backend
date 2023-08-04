package main

import (
	"context"
	"flag"
	"log"

	"github.com/FadyGamilM/hotelreservationapi/db"

	"github.com/FadyGamilM/hotelreservationapi/api"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const db_uri = "mongodb://localhost:27017"
const db_name = "hotel-reservation"
const user_collection = "users"

func main() {
	// define fiber app instance with the custom error handler config
	app := fiber.New(api.Config)

	// connect to db and get client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db_uri))
	if err != nil {
		log.Fatal(err)
	}

	// setup the port to be received from terminal while running or give it a default value
	listenAddr := flag.String("listen_port", ":5000", "port for server to listen to")
	flag.Parse()

	// initilize mongodb user store impl
	userMongoRepo := db.NewMongoUserRepo(client)

	// initialize user handler
	userHandler := api.NewUserHandler(userMongoRepo)

	// group and versioning
	users_router_v1 := app.Group("/api/v1/users")
	users_router_v1.Get("/", userHandler.HandleGetUsers)
	users_router_v1.Get("/:id", userHandler.HandleGetUserByID)
	users_router_v1.Post("/", userHandler.HandleCreateUser)

	// listen to the port and start the server
	app.Listen(*listenAddr)
}
