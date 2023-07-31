package main

import (
	"github.com/FadyGamilM/hotelreservationapi/internal/adapters/handlers"
	"github.com/gofiber/fiber/v2"
)

func Routes() {
	handler, _ := handlers.NewHttpHandler()
	router := fiber.New()
	router.Get("/ready", handler.TestReadinessHandler)
	router.Listen(":5000")
}
