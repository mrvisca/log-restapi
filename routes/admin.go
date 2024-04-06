package routes

import (
	"log-restapi/config"
	"log-restapi/models"

	"github.com/gin-gonic/gin"
)

func Changeadmin(c *gin.Context) {
	id := c.Param("id")
	keypass := c.PostForm("password")

	var item models.User
	if config.DB.First(&item, "id = ?", id).RecordNotFound() {
		c.JSON(400, gin.H{
			"status":  "Elor",
			"message": "Elor, data pengguna tidak ditemukan",
		})
		c.Abort()
		return
	}

	if keypass != "backend ganteng" {
		c.JSON(400, gin.H{
			"status":  "Elor",
			"message": "Kode akses salah, pastikan anda memiliki akses tersebut",
		})
		c.Abort()
		return
	} else {
		config.DB.Model(&item).Where("id = ?", id).Update("is_mimin", true)

		c.JSON(201, gin.H{
			"status":  "Berhasil mengubah akses admin",
			"message": "Berhasil mengubah akses pengguna ke admin",
		})
	}
}

func UbahLimit(c *gin.Context) {
	id := c.Param("id")
	is_mimin := bool(c.MustGet("jwt_is_mimin").(bool))
	keypass := c.PostForm("password")

	var item models.User
	if config.DB.First(&item, "id = ?", id).RecordNotFound() {
		c.JSON(400, gin.H{
			"status":  "Elor",
			"message": "Elor, data pengguna tidak ditemukan",
		})
		c.Abort()
		return
	}

	if keypass != "backend ganteng" && !is_mimin {
		c.JSON(400, gin.H{
			"status":  "Elor",
			"message": "Kode akses salah, pastikan anda memiliki akses tersebut",
		})
		c.Abort()
		return
	} else {
		config.DB.Model(&item).Where("id = ?", id).Update("limit", c.PostForm("limit"))

		c.JSON(201, gin.H{
			"status":  "Berhasil mengubah limit pengguna",
			"message": "Berhasil mengubah data limit pengguna",
		})
	}
}
