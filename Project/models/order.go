package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Status     string
	OrderItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID  uint
	ItemID   uint
	Item     Item
	Quantity int
}

type Item struct {
	gorm.Model
	ItemName string
	Amount   float32
}
