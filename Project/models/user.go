package models

type Manager struct {
	ID       uint   `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `gorm:"unique" json:"email,omitempty"`
	Password []byte `json:"-"`
}

type User struct {
	ID          uint   `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	PhoneNumber string `gorm:"unique" json:"email,omitempty"`
	Region      uint   `json:"region,omitempty"`
	Address     string `json:"address,omitempty"`
	Credit      uint   `json:"credit,omitempty"`
	Password    []byte `json:"-"`
}
