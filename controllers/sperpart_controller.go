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

func UpdateSperpart(c *gin.Context) {
	id := c.Param("id")
	var sperpart models.Sperpat

	// cek dulu datanya ada atau tidak
	if err := config.DB.First(&sperpart, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sperpart not found"})
		return
	}

	if err := c.ShouldBindJSON(&sperpart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&sperpart).Updates(sperpart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update sperpart"})
		return
	}

	c.JSON(http.StatusOK, sperpart)
}

func DeleteSperpart(c *gin.Context) {
	id := c.Param("id")
	var sperpart models.Sperpat
	config.DB.Delete(&sperpart, "id = ?", id)
	
	c.JSON(http.StatusOK, gin.H{"message": "Sperpart deleted successfully"})
}
