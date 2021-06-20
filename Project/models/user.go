package models

import "gorm.io/gorm"

type Manager struct {
	ID           uint   `gorm:"primaryKey" json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Email        string `gorm:"unique" json:"email,omitempty"`
	Password     []byte `json:"-"`
	RestaurantID uint   `json:"restaurantID,omitempty"`
	gorm.Model
}

type User struct {
	ID          uint   `gorm:"primaryKey" json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	PhoneNumber string `gorm:"unique" json:"phoneNumber,omitempty"`
	Region      string `json:"region,omitempty"` //region: enum
	Address     string `json:"address,omitempty"`
	Credit      uint   `json:"credit,omitempty"`
	Password    []byte `json:"-"`
	gorm.Model
}
