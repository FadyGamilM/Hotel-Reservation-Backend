package db

import "github.com/FadyGamilM/hotelreservationapi/types"

type HotelRepo interface {
	GetHotels() ([]*types.Hotel, error)
	CreateHotel(types.Hotel) (*types.Hotel, error)
	GetHotelByID(int64) (*types.Hotel, error)
	GetHotelRoomsByHotelID(int64) ([]*types.Room, error)
}
