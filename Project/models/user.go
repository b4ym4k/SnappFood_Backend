package models

type Manager struct {
	ID    uint   `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `gorm:"unique" json:"email,omitempty"`
	//Region   string    `json:"region,omitempty"` //region: enum
	//Address  string    `json:"address,omitempty"`
	Password []byte `json:"-"`
}

type User struct {
	ID          uint   `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	PhoneNumber string `gorm:"unique" json:"phoneNumber,omitempty"`
	Region      string `json:"region,omitempty"` //region: enum
	Address     string `json:"address,omitempty"`
	Credit      uint   `json:"credit,omitempty"`
	Password    []byte `json:"-"`
}
