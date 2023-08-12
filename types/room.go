package types

import "time"

type Room struct {
	ID         int64
	RoomNumber string
	RoomTypeID int64
	HotelID    int64
}

type GetRoomResponse struct {
	ID         int64  `json:"id"`
	RoomNumber string `json:"room_number"`
	RoomTypeID int64  `json:"room_type_id"`
	HotelID    int64  `json:"hotel_id"`
}

type CreateRoomRequest struct {
	RoomNumber string `json:"room_number"`
	RoomTypeID int64  `json:"room_type_id"`
	HotelID    int64  `json:"hotel_id"`
}

type CreateRoomResponse struct {
	GetRoomResponse
}

type PostgresRoom struct {
	ID         int64
	RoomNumber string
	RoomTypeID int64
	HotelID    int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
