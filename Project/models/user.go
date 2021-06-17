package models

type Manager struct {
	ID       uint   `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `gorm:"unique" json:"email,omitempty"`
	Region   uint   `json:"region,omitempty"` //region: enum
	Address  string `json:"address,omitempty"`
	Password []byte `json:"-"`
	//6- service regions: array of enum
	//7- working hours: array
	//8- delivery time
	//9- delivery cost
	//10- food list: array
	//11- order list
	//12- comment list
}

type User struct {
	ID          uint   `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	PhoneNumber string `gorm:"unique" json:"email,omitempty"`
	Region      uint   `json:"region,omitempty"` //region: enum
	Address     string `json:"address,omitempty"`
	Credit      uint   `json:"credit,omitempty"`
	Password    []byte `json:"-"`

	//8- favourite food ids : array
	//9- order history : array
}
