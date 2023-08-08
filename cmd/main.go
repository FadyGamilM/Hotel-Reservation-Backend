package main

import (
	"flag"
	"log"

	"github.com/FadyGamilM/hotelreservationapi/api"
	"github.com/FadyGamilM/hotelreservationapi/db/postgres"
	"github.com/gofiber/fiber/v2"
)

const db_uri = "mongodb://localhost:27017"
const db_name = "hotel-reservation"
const user_collection = "users"

func main() {
	// define fiber app instance with the custom error handler config
	app := fiber.New(api.Config)

	// setup the port to be received from terminal while running or give it a default value
	listenAddr := flag.String("listen_port", ":5000", "port for server to listen to")
	flag.Parse()

	// initilize postgres user store impl
	developmentDSN := "host=localhost port=1111 user=postgres password=postgres dbname=hrdb sslmode=disable timezone=UTC connect_timeout=5"
	userPostgresRepo, err := postgres.NewPostgresRepo(developmentDSN)
	if err != nil {
		log.Printf("error while creating a new postgres repository instance")
	}

	// initialize user handler
	userHandler := api.NewUserHandler(userPostgresRepo)

	// group and versioning
	users_router_v1 := app.Group("/api/v1/users")
	users_router_v1.Get("/", userHandler.HandleGetUsers)
	users_router_v1.Get("/:id", userHandler.HandleGetUserByID)
	users_router_v1.Post("/", userHandler.HandleCreateUser)
	users_router_v1.Delete("/:id", userHandler.HandleDeleteUser)
	users_router_v1.Put("/:id", userHandler.HandleUpdateUser)

	// listen to the port and start the server
	app.Listen(*listenAddr)
}
