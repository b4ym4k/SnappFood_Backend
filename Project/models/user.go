package models

import "gorm.io/gorm"

type Manager struct {
	ID         uint         `gorm:"primaryKey; not null" json:"id,omitempty"`
	Name       string       `gorm:"not null" json:"name,omitempty"`
	Email      string       `gorm:"unique; not null" json:"email,omitempty"`
	Password   []byte       `gorm:"not null" json:"-"`
	Restaurant []Restaurant `gorm:"foreignKey:ManagerID;references:ID"`
	gorm.Model
}

type User struct {
	ID          uint   `gorm:"primaryKey; not null" json:"id,omitempty"`
	Name        string `gorm:"not null" json:"name,omitempty"`
	PhoneNumber string `gorm:"unique; not null" json:"phoneNumber,omitempty"`
	Region      string `json:"region,omitempty"` //region: enum
	Address     string `json:"address,omitempty"`
	Credit      uint   `json:"credit,omitempty"`
	Password    []byte `gorm:"not null" json:"-"`
	gorm.Model
}

//example:

//type User struct {
//	OrganizationID uint   `gorm:"primaryKey; not null"`
//	Name           string `gorm:"primaryKey; not null"`
//}
//
//type Note struct {
//	ID             uint   `gorm:"primaryKey; not null"`
//	OrganizationID uint   `gorm:"not null"`
//	UserName       string `gorm:"not null"`
//	User           User   `gorm:"ForeignKey:OrganizationID,UserName;References:OrganizationID,Name"`
//}
