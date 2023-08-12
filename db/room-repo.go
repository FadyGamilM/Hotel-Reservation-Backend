package db

import "github.com/FadyGamilM/hotelreservationapi/types"

type RoomRepo interface {
	CreateRoom(types.Room) (*types.Room, error)
	GetRoomsByHotelID(int64) ([]*types.Room, error)
	DeleteRoomByHotelID(int64) error
}
