package controller

import (
	"net/http"

	"github.com/tzsoulcap/ui-sa/entity"
	"github.com/gin-gonic/gin"
)

// POST /Prefecture
func CreatePrefecture(c *gin.Context) {
	var prefecture entity.Prefecture
	if err := c.ShouldBindJSON(&prefecture); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&prefecture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": prefecture})
}

// GET /Prefecture/:id
func GetPrefecturebyproviceid(c *gin.Context) {
	var prefecture []entity.Prefecture
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM prefectures WHERE province_id = ?", id).Scan(&prefecture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": prefecture})
}

// GET /Prefecture
func ListPrefecture(c *gin.Context) {
	var prefecture []entity.Prefecture
	if err := entity.DB().Raw("SELECT * FROM prefectures").Scan(&prefecture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": prefecture})
}

// DELETE /Prefecture/:id
func DeletePrefecture(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM prefectures WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prefectures not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Prefecture
func UpdatePrefecture(c *gin.Context) {
	var prefecture entity.Prefecture
	if err := c.ShouldBindJSON(&prefecture); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", prefecture.ID).First(&prefecture); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prefecture not found"})
		return
	}

	if err := entity.DB().Save(&prefecture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": prefecture})
}
