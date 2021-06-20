package models

type Restaurant struct {
	ID          uint    `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Manager     Manager `gorm:"unique" json:"email,omitempty"`
	Region      string  `json:"region,omitempty"` //region: enum
	Address     string  `json:"address,omitempty"`
	Balance     float64 `json:"balance,omitempty"`
	Food        []Food  `json:"food,omitempty"`
	LastUpdated string  `json:"last_updated"`
}
