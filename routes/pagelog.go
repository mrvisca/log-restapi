package routes

import (
	"log-restapi/config"
	"log-restapi/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ResLog(logact models.LogAct) models.Dlogact {
	return models.Dlogact{
		ID:         logact.ID,
		OfficeName: logact.Office.Name,
		Username:   logact.Username,
		Endpoint:   logact.Endpoint,
		Halaman:    logact.Halaman,
		Aksi:       logact.Aksi,
		Keterangan: logact.Keterangan,
		Tipe:       logact.Tipe,
		IsLogin:    logact.IsLogin,
		Latitude:   logact.Latitude,
		Longtitude: logact.Longtitude,
		Alamat:     logact.Alamat,
		Mac:        logact.Mac,
		DeviceName: logact.DeviceName,
		DeviceType: logact.DeviceType,
		Version:    logact.Version,
	}
}

func Getlog(c *gin.Context) {
	items := []models.LogAct{}

	userid := uint(c.MustGet("jwt_user_id").(float64))
	id := c.Param("id")

	valoff, _ := strconv.Atoi(id)
	if valoff == 0 {
		config.DB.Where("user_id = ?", userid).Preload("Office", func(db *gorm.DB) *gorm.DB {
			return db.Order("name")
		}).Find(&items)

		list := []models.Dlogact{}
		for _, item := range items {
			list = append(list, ResLog(item))
		}

		c.JSON(200, gin.H{
			"status":  "Berhasil akses data log",
			"message": "Berhasil akses data log usaha keseluruhan",
			"data":    list,
		})
	} else {
		config.DB.Where("user_id = ?", userid).Preload("Office", "id = ?", valoff, func(db *gorm.DB) *gorm.DB {
			return db.Order("name")
		}).Find(&items)

		list := []models.Dlogact{}
		for _, item := range items {
			list = append(list, ResLog(item))
		}

		c.JSON(200, gin.H{
			"status":  "Berhasil akses data log",
			"message": "Berhasil akses data log usaha keseluruhan",
			"data":    list,
		})
	}
}

func PostLog(c *gin.Context) {
	userid := uint(c.MustGet("jwt_user_id").(float64))
	id := c.PostForm("officeid")
	valoff, _ := strconv.Atoi(id)

	if valoff == 0 || id == "" {
		c.JSON(400, gin.H{
			"status":  "Elor Post",
			"message": "Pastikan body request officeid tidak boleh kosong / 0",
		})
		c.Abort()
		return
	}

	is_login := c.PostForm("is_login")
	LoginIs, _ := strconv.ParseBool(is_login)

	item := models.LogAct{
		UserId:     userid,
		OfficeId:   uint(valoff),
		Username:   c.PostForm("username"),
		Endpoint:   c.PostForm("endpoint"),
		Halaman:    c.PostForm("halaman"),
		Aksi:       c.PostForm("aksi"),
		Keterangan: c.PostForm("keterangan"),
		Tipe:       c.PostForm("tipe"),
		IsLogin:    LoginIs,
		Latitude:   c.PostForm("latitude"),
		Longtitude: c.PostForm("longtitude"),
		Alamat:     c.PostForm("alamat"),
		Mac:        c.PostForm("mac"),
		DeviceName: c.PostForm("device_name"),
		DeviceType: c.PostForm("device_type"),
		Version:    c.PostForm("version"),
	}

	config.DB.Create(&item)
	response := ResLog(item)

	c.JSON(201, gin.H{
		"status":  "Berhasil post data",
		"message": "Berhasil post data log aktivitas kantor",
		"data":    response,
	})
}
