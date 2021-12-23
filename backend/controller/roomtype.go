package controller

import (
	"net/http"

	"github.com/tzsoulcap/ui-sa/entity"
	"github.com/gin-gonic/gin"
)

// POST /roomtype
func CreateRoomtype(c *gin.Context) {
	var roomtype entity.Roomtype
	if err := c.ShouldBindJSON(&roomtype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	if err := entity.DB().Create(&roomtype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roomtype})
}

// GET /room/:id
func GetRoomtype(c *gin.Context) {
	var roomtype entity.Roomtype
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM roomtypes WHERE id = ?", id).Scan(&roomtype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roomtype})
}

// GET /roomtype
func ListRoomtype(c *gin.Context) {
	var roomtype []entity.Roomtype
	if err := entity.DB().Raw("SELECT * FROM roomtypes").Scan(&roomtype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roomtype})
}

// DELETE /roomtype/:id
func DeleteRoomtype(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM roomtypes WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /roomtype
func UpdateRoomtype(c *gin.Context) {
	var roomtype entity.Roomtype
	if err := c.ShouldBindJSON(&roomtype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", roomtype.ID).First(&roomtype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roomtype not found"})
		return
	}
	if err := entity.DB().Save(&roomtype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roomtype})
}
