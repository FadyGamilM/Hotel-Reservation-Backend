package handlers

import (
	"github.com/FadyGamilM/hotelreservationapi/internal/adapters/repository/mongo"
	"github.com/gofiber/fiber/v2"
)

type HttpHandler struct {
	// should depends on the service port which is the input ports which is defined in the core/ports/input folder
}

// should receive the service port as a param
func NewHttpHandler() (*HttpHandler, error) {
	return &HttpHandler{}, nil
}

func (h *HttpHandler) TestReadinessHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]string{"resp": "working just fine !"})
}

func (h *HttpHandler) CreateUser(ctx *fiber.Ctx) error {
	mongodb := mongo.NewMongoDB("mongodb://localhost:27017")
	return ctx.JSON("inserted")
}
