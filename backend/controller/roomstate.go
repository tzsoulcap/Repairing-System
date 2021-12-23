package controller

import (
	"net/http"

	"github.com/tzsoulcap/ui-sa/entity"
	"github.com/gin-gonic/gin"
)

// POST /roomstate
func CreateRoomstate(c *gin.Context) {
	var roomstate entity.Roomstate
	if err := c.ShouldBindJSON(&roomstate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	if err := entity.DB().Create(&roomstate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roomstate})
}

// GET /roomstate/:id
func GetRoomstate(c *gin.Context) {
	var roomstate entity.Roomstate
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM roomstates WHERE id = ?", id).Scan(&roomstate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roomstate})
}

// GET /roomstate
func ListRoomstate(c *gin.Context) {
	var roomstate []entity.Roomstate
	if err := entity.DB().Raw("SELECT * FROM roomstates").Scan(&roomstate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roomstate})
}

// DELETE /roomstate/:id
func DeleteRoomstate(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM roomstates WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /roomstate
func UpdateRoomstate(c *gin.Context) {
	var roomstate entity.Roomstate
	if err := c.ShouldBindJSON(&roomstate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", roomstate.ID).First(&roomstate); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roomstate not found"})
		return
	}
	if err := entity.DB().Save(&roomstate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roomstate})
}
