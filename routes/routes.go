package routes

import (
	"listsperpat/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()

	r.GET("/sperpart", controllers.GetAllSperpart)
	r.POST("/sperpart", controllers.CreateSperpart)

	return r
}
