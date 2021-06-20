package models

type Food struct {
	ID           uint   `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Price        int    `json:"price,omitempty"`
	Ready        bool   `json:"isReady,omitempty"`
	RestaurantID int    `json:"RestaurantID,omitempty"`
	Rate         int    `json:"rate,omitempty"`
}
