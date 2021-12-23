package controller

import (
	"net/http"

	"github.com/tzsoulcap/ui-sa/entity"
	"github.com/gin-gonic/gin"
)

// POST /Career
func CreateCareer(c *gin.Context) {
	var career entity.Career
	if err := c.ShouldBindJSON(&career); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&career).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": career})
}

// GET /Career/:id
func GetCareer(c *gin.Context) {
	var career entity.Career
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM careers WHERE id = ?", id).Scan(&career).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": career})
}

// GET /Career
func ListCareer(c *gin.Context) {
	var career []entity.Career
	if err := entity.DB().Raw("SELECT * FROM careers").Scan(&career).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": career})
}

// DELETE /Career/:id
func DeleteCareer(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM careers WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "careers not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Career
func UpdateCareer(c *gin.Context) {
	var career entity.Career
	if err := c.ShouldBindJSON(&career); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", career.ID).First(&career); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "careers not found"})
		return
	}

	if err := entity.DB().Save(&career).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": career})
}
