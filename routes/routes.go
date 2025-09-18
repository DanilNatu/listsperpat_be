package routes

import (
	"listSparepart/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()

	r.GET("/sparepart", controllers.GetAllSparepart)
	r.POST("/sparepart", controllers.CreateSparepart)
	r.PUT("/sparepart/:id", controllers.UpdateSparepart)
	r.DELETE("/sparepart/:id", controllers.DeleteSparepart)

	return r
}
