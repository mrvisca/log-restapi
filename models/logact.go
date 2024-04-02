package models

import "github.com/jinzhu/gorm"

type LogAct struct {
	gorm.Model
	User       User   `gorm:"foreignKey:UserId"`
	Office     Office `gorm:"foreignKey:OfficeId"`
	UserId     uint
	OfficeId   uint
	Username   string
	Endpoint   string
	Halaman    string
	Aksi       string
	Keterangan string
	Tipe       string
	IsLogin    bool
	Latitude   string
	Longtitude string
	Alamat     string
	Mac        string
	DeviceName string
	DeviceType string
	Version    string
}
