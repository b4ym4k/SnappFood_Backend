package models

type Restaurant struct {
	ID             uint    `json:"id,omitempty"`
	RestaurantName string  `json:"restaurantName,omitempty"`
	Manager        Manager `gorm:"unique" json:"manager,omitempty"`
	Region         string  `json:"region,omitempty"` //region: enum
	Address        string  `json:"address,omitempty"`
	Balance        float64 `json:"balance,omitempty"`
	Food           []Food  `json:"food,omitempty"`
	LastUpdated    string  `json:"last_updated"`
}
