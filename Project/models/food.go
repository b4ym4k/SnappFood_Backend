package models

import "gorm.io/gorm"

type Food struct {
	ID           uint   `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Price        int    `json:"price,omitempty"`
	Ready        bool   `json:"isReady,omitempty"`
	Rate         int    `json:"rate,omitempty"`
	RestaurantID int    `json:"restaurantID,omitempty"`
	gorm.Model
}
