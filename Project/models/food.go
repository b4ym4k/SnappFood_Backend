package models

import "gorm.io/gorm"

type Food struct {
	ID           uint   `gorm:"primaryKey" json:"id,omitempty"`
	Name         string `gorm:"not null" json:"name,omitempty"`
	Price        string `gorm:"not null" json:"price,omitempty"`
	Ready        bool   `json:"isReady,omitempty"`
	Rate         int    `json:"rate,omitempty"`
	RestaurantID int    `json:"restaurantID,omitempty"`
	gorm.Model
}
