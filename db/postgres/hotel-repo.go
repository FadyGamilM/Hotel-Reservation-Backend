package postgres

import "github.com/FadyGamilM/hotelreservationapi/types"

type HotelPostgresRepo struct {
	dbRepo *PostgresRepo
}

func NewHotelPostgresRepo(pr *PostgresRepo) *HotelPostgresRepo {
	return &HotelPostgresRepo{
		dbRepo: pr,
	}
}

func (hr *HotelPostgresRepo) GetHotels() ([]*types.Hotel, error) {
	ctx, cancel := CreateContext()
	defer cancel()
}
