package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Offices  []Office
	LogActs  []LogAct
	Username string
	Fullname string
	Email    string `gorm:"unique;not null"`
	SocialId string
	Provider string
	Avatar   string
	Role     bool `gorm:"default:0"`
	Phone    string
	Limit    uint
	IsMimin  bool `gorm:"default:0"`
}
