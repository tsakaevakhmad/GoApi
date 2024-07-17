package domain

type Album struct {
	Id     int64   `json:"id,omitempty"`
	Price  float64 `json:"price"`
	Name   string  `json:"name"`
	Artist string  `json:"artist"`
}
