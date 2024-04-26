package main

import (
	"log-restapi/config"
	_ "log-restapi/docs"
	"log-restapi/middleware"
	"log-restapi/routes"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Service API Documentation
// @version 1.0
// @description Dokumentasi Layanan API untuk monitor log aktivitas
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support (Mr.Visca)
// @contact.url    https://resume.mrvisca.tech
// @contact.email  bimaputra@mrvisca.tech

// @host	localhost:8080
// @basePath /api/v1/
func main() {
	config.InitDB()
	defer config.DB.Close()
	gotenv.Load()

	router := gin.Default()

	// Tambahkan route dokumentasi swagger
	router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1/")
	{
		v1.GET("/auth/:provider", routes.RedirectHandler)
		v1.GET("/auth/:provider/callback", routes.CallbackHandler)
		v1.PUT("/auth/admin-pass/:id", routes.Changeadmin)

		kantor := v1.Group("office/")
		{
			kantor.GET("/", middleware.IsAuth(), routes.GetOffice)
			kantor.GET("detail/:id", middleware.IsAuth(), routes.DetailOffice)
			kantor.POST("/post", middleware.IsAuth(), routes.PostOffice)
			kantor.PUT("update/:id", middleware.IsAuth(), routes.UpdateOffice)
			kantor.DELETE("hapus/:id", middleware.IsAuth(), routes.DeleteOffice)
		}

		halog := v1.Group("page-log/")
		{
			halog.GET("/:officeid", middleware.IsAuth(), routes.Getlog)
			halog.POST("/add", middleware.IsAuth(), routes.PostLog)
		}

		devs := v1.Group("devs-only/")
		{
			devs.PUT("update-limit/:id", middleware.IsDev(), routes.UbahLimit)
			devs.DELETE("hapus-akun/:id", middleware.IsDev(), routes.HapusAkun)
		}
	}

	router.Run()
}
