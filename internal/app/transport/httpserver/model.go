package httpserver

type ShowResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ShowListResponse struct {
	Response []ShowResponse `json:"response"`
}

type EventResponse struct {
	ID     int    `json:"id"`
	ShowID int    `json:"showId"`
	Date   string `json:"date"`
}

type EventListResponse struct {
	Response []EventResponse `json:"response"`
}

type PlaceResponse struct {
	ID          int     `json:"id"`
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
	Width       float64 `json:"width"`
	Height      float64 `json:"height"`
	IsAvailable bool    `json:"is_available"` // nolint: tagliatelle
}

type PlaceListResponse struct {
	Response []PlaceResponse `json:"response"`
}
