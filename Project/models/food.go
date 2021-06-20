package models

type Food struct {
	ID         uint       `json:"id,omitempty"`
	Name       string     `json:"name,omitempty"`
	Price      int        `json:"price,omitempty"`
	Ready      bool       `json:"isReady,omitempty"`
	Restaurant Restaurant `json:"RestaurantID,omitempty"`
	Rate       int        `json:"rate,omitempty"`
}
