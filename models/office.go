package models

import "github.com/jinzhu/gorm"

type Office struct {
	gorm.Model
	LogActs []LogAct
	User    User `gorm:"foreignKey:UserId"`
	UserId  uint
	Name    string
	Email   string
	Alamat  string
	Telpon  string
	Join    string
}
