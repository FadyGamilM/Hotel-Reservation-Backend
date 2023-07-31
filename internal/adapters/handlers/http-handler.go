package handlers

import "github.com/gofiber/fiber/v2"

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
