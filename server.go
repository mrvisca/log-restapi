package main

import (
	"log-restapi/config"
	"log-restapi/middleware"
	"log-restapi/routes"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	config.InitDB()
	defer config.DB.Close()
	gotenv.Load()

	router := gin.Default()

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
		}
	}

	router.Run()
}
