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

func CreateSperpart(c *gin.Context) {
	var input models.Sperpat
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create sperpart"})
		return
	}
	c.JSON(http.StatusCreated, input)
}
