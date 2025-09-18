package controllers

import (
	config "listsperpat/db"
	"listsperpat/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllSperpart(c *gin.Context) {
	var sperpart []models.Sperpat
	config.DB.Find(&sperpart)
	c.JSON(http.StatusOK, sperpart)
}

