package types

type Hotel struct {
	ID        int64
	HotelName string
	Location  string
	Stars     int8
}

type GetHotelResponse struct {
	ID        int64  `json:"id"`
	HotelName string `json:"hotel_name"`
	Location  string `json:"location"`
	Stars     int8   `json:"stars"`
}

type CreateHotelRequest struct {
	HotelName string `json:"hotel_name"`
	Location  string `json:"location"`
	Stars     int8   `json:"stars"`
}

type CreateHotelResponse struct {
	GetHotelResponse
}
