package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/FadyGamilM/hotelreservationapi/types"

	"github.com/FadyGamilM/hotelreservationapi/api"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const db_uri = "mongodb://localhost:27017"
const db_name = "hotel-reservation"
const user_collection = "users"

func main() {
	// define fiber app instance
	app := fiber.New()

	// connect to db and get client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db_uri))
	if err != nil {
		log.Fatal(err)
	}

	// => Insert user
	// define a context
	ctx := context.Background()
	// get the collection
	user_coll := client.Database(db_name).Collection(user_collection)
	// create user entity to be persisted
	user := types.User{
		FirstName: "fady",
		LastName:  "gamil",
	}
	// persist data
	res, err := user_coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	fmt.Println(res.InsertedID)

	// => Fetch First user matches the empty filter (will be the 1st record in db)
	var foundUser types.User
	err = user_coll.FindOne(ctx, bson.M{}).Decode(&foundUser)
	if err != nil {
		log.Fatal(err)
	}
	// print the decoded user
	fmt.Println(foundUser)

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
