package models

import "gorm.io/gorm"

type Restaurant struct {
	ID             uint    `json:"id,omitempty"`
	RestaurantName string  `json:"restaurantName,omitempty"`
	Manager        Manager `gorm:"foreignKey:Email"`
	Region         string  `json:"region,omitempty"` //region: enum
	Address        string  `json:"address,omitempty"`
	Balance        float64 `json:"balance,omitempty"`
	//Food           []Food  `json:"food,omitempty"`
	gorm.Model
}
