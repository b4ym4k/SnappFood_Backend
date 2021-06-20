package models

import "gorm.io/gorm"

type Restaurant struct {
	ID             uint    `gorm:"primaryKey" json:"id,omitempty"`
	RestaurantName string  `json:"restaurantName,omitempty"`
	Manager        Manager `gorm:"foreignKey:ID"`
	Region         string  `json:"region,omitempty"` //region: enum
	Address        string  `json:"address,omitempty"`
	Balance        float64 `json:"balance,omitempty"`
	//Food           []Food  `json:"food,omitempty"`
	gorm.Model
}
