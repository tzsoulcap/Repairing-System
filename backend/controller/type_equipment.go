package controller

import (
	"net/http"

	"github.com/tzsoulcap/ui-sa/entity"
	"github.com/gin-gonic/gin"
)

// POST /TypeEquipments
func CreateTypeEquipment(c *gin.Context) {
	var type_equipment entity.TypeEquipment
	if err := c.ShouldBindJSON(&type_equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&type_equipment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": type_equipment})
}

// GET /TypeEquipment/:id
func GetTypeEquipment(c *gin.Context) {
	var type_equipment entity.TypeEquipment
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM type_equipments WHERE id = ?", id).Scan(&type_equipment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": type_equipment})
}

// GET /TypeEquipments
func ListTypeEquipments(c *gin.Context) {
	var type_equipments []entity.TypeEquipment
	if err := entity.DB().Raw("SELECT * FROM type_equipments").Scan(&type_equipments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": type_equipments})
}

// DELETE /TypeEquipments/:id
func DeleteTypeEquipment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM type_equipments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "TypeEquipments not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /TypeEquipments
func UpdateTypeEquipment(c *gin.Context) {
	var type_equipment entity.TypeEquipment
	if err := c.ShouldBindJSON(&type_equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", type_equipment.ID).First(&type_equipment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "TypeEquipments not found"})
		return
	}

	if err := entity.DB().Save(&type_equipment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": type_equipment})
}
