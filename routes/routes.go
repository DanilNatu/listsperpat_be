package routes

import (
	"listSparepart/controllers"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8005"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/sparepart", controllers.GetAllSparepart)
	r.POST("/sparepart", controllers.CreateSparepart)
	r.PUT("/sparepart/:id", controllers.UpdateSparepart)
	r.DELETE("/sparepart/:id", controllers.DeleteSparepart)

	return r
}
