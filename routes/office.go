package routes

import (
	"log-restapi/config"
	"log-restapi/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetDofice(office models.Office) models.Doffice {
	return models.Doffice{
		ID:     office.ID,
		Name:   office.Name,
		Email:  office.Email,
		Alamat: office.Alamat,
		Telpon: office.Telpon,
		Join:   office.Join,
	}
}

func GetOffice(c *gin.Context) {
	items := []models.Office{}

	userid := uint(c.MustGet("jwt_user_id").(float64))
	config.DB.Where("user_id = ?", userid).Find(&items)

	// Proses filtering data tanpa relasi
	list := []models.Doffice{}
	for _, item := range items {
		list = append(list, GetDofice(item))
	}

	c.JSON(200, gin.H{
		"status": "Berhasil akses list office",
		"data":   list,
	})
}

func DetailOffice(c *gin.Context) {
	id := c.Param("id")

	var item models.Office

	if config.DB.First(&item, "id = ?", id).RecordNotFound() {
		c.JSON(404, gin.H{
			"status":  "Elor",
			"message": "Elor, Data kantor tidak ditemukan",
		})
		c.Abort()
		return
	}

	res_data := GetDofice(item)

	c.JSON(200, gin.H{
		"status": "Berhasil akses data profil kantor",
		"data":   res_data,
	})
}

func PostOffice(c *gin.Context) {
	var olditem models.Office
	name := c.PostForm("name")
	userid := uint(c.MustGet("jwt_user_id").(float64))
	limit := uint(c.MustGet("jwt_limit").(float64))
	if !config.DB.First(&olditem, "name = ?", name).RecordNotFound() {
		c.JSON(400, gin.H{
			"status":  "Elor",
			"message": "Elor, Post data gagal, pastikan nama kantor tidak boleh sama dalam 1 user",
		})
	} else {
		// Check limit office pada 1 user
		var count int64
		config.DB.Model(&models.User{}).Where("user_id = ?", userid).Count(&count)

		if limit > uint(count) {
			c.JSON(400, gin.H{
				"status":  "Elor",
				"message": "Limit telah habis, tidak dapat membuat data kantor baru",
			})
			c.Abort()
			return
		}

		// Format tanggal saat ini
		tanggal := time.Now()

		item := models.Office{
			UserId: userid,
			Name:   c.PostForm("name"),
			Email:  c.PostForm("email"),
			Alamat: c.PostForm("alamat"),
			Telpon: c.PostForm("telpon"),
			Join:   tanggal.Format("2006-01-02"),
		}

		config.DB.Create(&item)
		response := GetDofice(item)

		c.JSON(201, gin.H{
			"status":  "Berhasil",
			"message": "Post data kantor berhasil!",
			"data":    response,
		})
	}
}

func UpdateOffice(c *gin.Context) {
	id := c.Param("id")

	var item models.Office

	if config.DB.First(&item, "id = ?", id).RecordNotFound() {
		c.JSON(404, gin.H{
			"status":  "Elor",
			"message": "Elor, data kantor tidak ditemukan",
		})
		c.Abort()
		return
	}

	// Filter update sesuai dengan akses user_id

	// Jika data ditemukan maka akan dilakukan update data office
	config.DB.Model(&item).Where("id = ?", id).Updates(models.Office{
		Name:   c.PostForm("name"),
		Email:  c.PostForm("email"),
		Alamat: c.PostForm("alamat"),
		Telpon: c.PostForm("telpon"),
	})

	response := GetDofice(item)

	c.JSON(201, gin.H{
		"status": "Update data kator berhasil",
		"data":   response,
	})
}

func DeleteOffice(c *gin.Context) {
	id := c.Param("id")
	var office models.Office

	config.DB.Where("id = ?", id).Delete(&office)

	c.JSON(201, gin.H{
		"status":  "Berhasil hapus data!",
		"message": "Berhasil menghapus data kantor!",
	})
}
