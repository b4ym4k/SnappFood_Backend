package models

import "gorm.io/gorm"

type Restaurant struct {
	ID             uint    `gorm:"primaryKey" json:"id,omitempty"`
	RestaurantName string  `gorm:"not null" json:"restaurantName,omitempty"`
	Region         string  `gorm:"not null" json:"region,omitempty"` //region: enum
	Address        string  `json:"address,omitempty"`
	Balance        float64 `json:"balance,omitempty"`
	ManagerID      int     `gorm:"primaryKey; not null" json:"managerID,omitempty"`
	//Food           []Food  `gorm:"foreignKey:RestaurantID;references:ID" json:"food,omitempty"`
	//Food []Food `gorm:"foreignKey:RestaurantID; references:ID" json:"food,omitempty"`
	Food []Food `gorm:"foreignKey:ID; references:ID" json:"food,omitempty"`
	gorm.Model
}
