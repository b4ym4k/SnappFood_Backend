package models

import "gorm.io/gorm"

type Food struct {
	ID           uint   `gorm:"primaryKey" json:"id,omitempty"`
	FoodName     string `gorm:"not null" json:"foodName,omitempty"`
	Price        string `gorm:"not null" json:"price,omitempty"`
	Ready        bool   `json:"isReady,omitempty"`
	Rate         int    `json:"rate,omitempty"`
	RestaurantID int    `gorm:"primaryKey" json:"restaurantID,omitempty"`
	gorm.Model
}
