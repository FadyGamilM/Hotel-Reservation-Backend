package api

import (
	"strconv"

	"github.com/FadyGamilM/hotelreservationapi/db"
	"github.com/FadyGamilM/hotelreservationapi/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type HotelHandler struct {
	repo *db.Store
}

func NewHotelHandler(r *db.Store) *HotelHandler {
	return &HotelHandler{
		repo: r,
	}
}

/*
@ Logic :
➜ Call Repository Layer to get all hotels (domain entity)
➜ map the domain entity data into response dto
➜ returns the response
*/
func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	hotels, err := h.repo.Hotel.GetHotels()
	if err != nil {
		log.Info(err)
		return err
	}

	responseDtos := []types.GetHotelResponse{}

	for _, hotel := range hotels {
		responseDtos = append(responseDtos, types.GetHotelResponse{
			ID:        hotel.ID,
			HotelName: hotel.HotelName,
			Location:  hotel.Location,
			Stars:     hotel.Stars,
		})
	}

	return c.JSON(responseDtos)
}

func (h *HotelHandler) HandleGetHotelByID(c *fiber.Ctx) error {
	id := c.Params("id")
	hotelID, err := strconv.ParseInt(id, 10, 64)
	log.Infof("hotel id : %v \n", hotelID)

	if err != nil {
		log.Info(err)
		return err
	}
	hotel, err := h.repo.Hotel.GetHotelByID(hotelID)
	if err != nil {
		log.Info(err)
		return err
	}
	hotelResponse := types.GetHotelResponse{
		ID:        hotel.ID,
		HotelName: hotel.HotelName,
		Location:  hotel.Location,
		Stars:     hotel.Stars,
	}
	return c.JSON(hotelResponse)
}

func (h *HotelHandler) HandleCreateHotel(c *fiber.Ctx) error {
	var createHotelReqDto *types.CreateHotelRequest
	err := c.BodyParser(&createHotelReqDto)
	if err != nil {
		return err
	}

	var hotel types.Hotel
	hotel.HotelName = createHotelReqDto.HotelName
	hotel.Location = createHotelReqDto.Location
	hotel.Stars = createHotelReqDto.Stars

	createdHotel, err := h.repo.Hotel.CreateHotel(hotel)
	if err != nil {
		return err
	}
	var createHotelResponse *types.CreateHotelResponse
	createHotelResponse.ID = createdHotel.ID
	createHotelResponse.HotelName = createdHotel.HotelName
	createHotelResponse.Location = createdHotel.Location
	createHotelResponse.Stars = createdHotel.Stars
	return c.JSON(createHotelResponse)
}
