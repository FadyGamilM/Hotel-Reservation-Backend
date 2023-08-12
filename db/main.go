package db

type Store struct {
	User  UserRepo
	Hotel HotelRepo
	Room  RoomRepo
}
