package controller

import (
	"net/http"

	"github.com/tzsoulcap/ui-sa/entity"
	"github.com/gin-gonic/gin"
)

// POST /Equipments
func CreateEquipment(c *gin.Context) {
	var equipment entity.Equipment
	if err := c.ShouldBindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&equipment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": equipment})
}

// GET /Equipment/:id
func GetEquipment(c *gin.Context) {
	var equipment entity.Equipment

	id := c.Param("id")
	if err := entity.DB().Preload("TypeEquipment").Raw("SELECT * FROM equipment WHERE id = ?", id).Find(&equipment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": equipment})
}

// GET /Equipments
func ListEquipments(c *gin.Context) {
	var equipments []entity.Equipment
	if err := entity.DB().Preload("TypeEquipment").Raw("SELECT * FROM equipment").Find(&equipments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": equipments})
}

// DELETE /Equipments/:id
func DeleteEquipment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM equipment WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "equipment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Equipments
func UpdateEquipment(c *gin.Context) {
	var equipment entity.Equipment
	if err := c.ShouldBindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", equipment.ID).First(&equipment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "equipment not found"})
		return
	}

	if err := entity.DB().Save(&equipment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": equipment})
}
