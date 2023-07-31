package main

import (
	"github.com/FadyGamilM/hotelreservationapi/internal/adapters/handlers"
	"github.com/gofiber/fiber/v2"
)

func (app *AppConfig) Routes() {
	// We need to pass the service impl here but this will be handled later (even by passing it to the Routes method)
	handler, _ := handlers.NewHttpHandler()
	// define an instance of the server handler
	router := fiber.New()

	// versioning our api router
	router_v1 := router.Group("/api/v1")

	router_v1.Get("/ready", handler.TestReadinessHandler)
	router.Listen(app.ListenAddr)
}
