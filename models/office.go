package models

import "github.com/jinzhu/gorm"

type Office struct {
	gorm.Model
	LogActs []LogAct
	User    User `gorm:"foreignkey:UserId"`
	UserId  uint
	Name    string
	Email   string `gorm:"unique;not null"`
	Alamat  string
	Telpon  string
	Join    string
}

type Doffice struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Alamat string `json:"alamat"`
	Telpon string `json:"Telpon"`
	Join   string `json:"join"`
}
