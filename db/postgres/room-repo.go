package postgres

type PostgresRoomRepo struct {
	dbRepo *PostgresRepo
}

const (
	GetAllRoomsByHotelIdQuery = `
		SELECT * FROM rooms WHERE hotel_id = $1 
	`
	GetHotelByIdQuery = `
		SELECT * FROM hotels WHERE id = $1 
	`
	CreateHotelQuery = `
		INSERT INTO hotels 
		(hotel_name, location, stars)
		VALUES 
		($1, $2, $3)
		RETURNING id, hotel_name, location, stars, created_at, updated_at
	`
)
