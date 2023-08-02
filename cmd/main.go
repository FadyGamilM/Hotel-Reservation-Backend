package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/FadyGamilM/hotelreservationapi/api"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbUri = "mongodb://localhost:27017"

func main() {
	// define fiber app instance
	app := fiber.New()

	// connect to db and get client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(client)

	// setup the port to be received from terminal while running or give it a default value
	listenAddr := flag.String("listen_port", ":5000", "port for server to listen to")
	flag.Parse()

	// test the server
	app.Get("/readiness", api.HandleReadiness)

	// group and versioning
	users_router_v1 := app.Group("/api/v1/users")
	users_router_v1.Get("/", api.HandleGetUsers)
	users_router_v1.Get("/:id", api.HandleGetUserByID)

	// listen to the port and start the server
	app.Listen(*listenAddr)
}
