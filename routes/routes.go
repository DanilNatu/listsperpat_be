package routes

import (
	"listsperpat/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()

	r.GET("/sperpart", controllers.GetAllSperpart)
	r.POST("/sperpart", controllers.CreateSperpart)
	r.PUT("/sperpart/:id", controllers.UpdateSperpart)
	r.DELETE("/sperpart/:id", controllers.DeleteSperpart)

	return r
}
