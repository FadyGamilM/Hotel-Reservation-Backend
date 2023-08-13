package postgres

import "github.com/FadyGamilM/hotelreservationapi/types"

type HotelPostgresRepo struct {
	dbRepo *PostgresRepo
}

const (
	GetAllHotelsQuery = `SELECT * FROM hotels`
	GetHotelByIdQuery = `SELECT * FROM hotels WHERE id = $1`
	CreateHotelQuery  = `
		INSERT INTO hotels 
		(hotel_name, location, stars)
		VALUES 
		($1, $2, $3)
		RETURNING id, hotel_name, location, stars, created_at, updated_at
	`
)

func NewHotelPostgresRepo(pr *PostgresRepo) *HotelPostgresRepo {
	return &HotelPostgresRepo{
		dbRepo: pr,
	}
}

func (hr *HotelPostgresRepo) GetHotels() ([]*types.Hotel, error) {
	ctx, cancel := CreateContext()
	defer cancel()

	rows, err := hr.dbRepo.db.QueryContext(ctx, GetAllHotelsQuery)
	if err != nil {
		return nil, err
	}

	dbEntities := []types.PostgresHotel{}
	for rows.Next() {
		dbEntity := types.PostgresHotel{}
		err = rows.Scan(
			&dbEntity.ID,
			&dbEntity.HotelName,
			&dbEntity.Location,
			&dbEntity.Stars,
			&dbEntity.CreatedAt,
			&dbEntity.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		dbEntities = append(dbEntities, dbEntity)
	}

	domainEntities := []*types.Hotel{}
	for _, dbHotel := range dbEntities {
		domainEntities = append(domainEntities, &types.Hotel{
			ID:        dbHotel.ID,
			HotelName: dbHotel.HotelName,
			Location:  dbHotel.Location,
			Stars:     dbHotel.Stars,
		})
	}
	return domainEntities, nil
}

func (hr *HotelPostgresRepo) CreateHotel(hotels types.Hotel) (*types.Hotel, error) {
	ctx, cancel := CreateContext()
	defer cancel()

	postgresHotelEntity := new(types.PostgresHotel)
	err := hr.dbRepo.db.QueryRowContext(ctx, CreateHotelQuery, hotels.HotelName, hotels.Location, hotels.Stars).Scan(&postgresHotelEntity)
	if err != nil {
		return nil, err
	}

	domainHotelEntity := new(types.Hotel)
	domainHotelEntity.ID = postgresHotelEntity.ID
	domainHotelEntity.HotelName = postgresHotelEntity.HotelName
	domainHotelEntity.Location = postgresHotelEntity.Location
	domainHotelEntity.Stars = postgresHotelEntity.Stars
	return domainHotelEntity, nil
}
