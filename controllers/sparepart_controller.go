package controllers

import (
	config "listSparepart/db"
	"listSparepart/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllSparepart(c *gin.Context) {
	var Sparepart []models.Sparepart
	config.DB.Find(&Sparepart)
	c.JSON(http.StatusOK, Sparepart)
}

func CreateSparepart(c *gin.Context) {
	var input models.Sparepart
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Sparepart"})
		return
	}
	c.JSON(http.StatusCreated, input)
}

func UpdateSparepart(c *gin.Context) {
	id := c.Param("id")
	var Sparepart models.Sparepart

	// cek dulu datanya ada atau tidak
	if err := config.DB.First(&Sparepart, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sparepart not found"})
		return
	}

	if err := c.ShouldBindJSON(&Sparepart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&Sparepart).Updates(Sparepart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Sparepart"})
		return
	}

	c.JSON(http.StatusOK, Sparepart)
}

func DeleteSparepart(c *gin.Context) {
	id := c.Param("id")
	var Sparepart models.Sparepart
	config.DB.Delete(&Sparepart, "id = ?", id)

	c.JSON(http.StatusOK, gin.H{"message": "Sparepart deleted successfully"})
}
