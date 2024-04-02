package config

import (
	"log-restapi/models"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("mysql", "root:@(localhost)/loggin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Koneksi ke database gagal!")
	}

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Office{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	DB.AutoMigrate(&models.LogAct{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").AddForeignKey("office_id", "offices(id)", "CASCADE", "CASCADE")

	DB.Model(&models.User{}).Related(&models.Office{}).Related(&models.LogAct{})
}
