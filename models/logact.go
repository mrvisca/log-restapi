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

type Dlogact struct {
	ID         uint   `json:"id"`
	OfficeName string `json:"office_name"`
	Username   string `json:"username"`
	Endpoint   string `json:"endpoint"`
	Halaman    string `json:"halaman"`
	Aksi       string `json:"aksi"`
	Keterangan string `json:"keterangan"`
	Tipe       string `json:"tipe"`
	IsLogin    bool   `json:"is_login"`
	Latitude   string `json:"latitude"`
	Longtitude string `json:"longtitude"`
	Alamat     string `json:"alamat"`
	Mac        string `json:"mac"`
	DeviceName string `json:"device_name"`
	DeviceType string `json:"device_type"`
	Version    string `json:"version"`
}
