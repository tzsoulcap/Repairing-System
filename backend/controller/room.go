package controller

import (
	"net/http"

	"github.com/tzsoulcap/ui-sa/entity"
	"github.com/gin-gonic/gin"
)

// POST /room
func CreateRoom(c *gin.Context) {
	var rooms entity.Room
	if err := c.ShouldBindJSON(&rooms); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	if err := entity.DB().Create(&rooms).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rooms})
}

// GET /room/:id
func GetRoom(c *gin.Context) {
	var rooms entity.Room
	id := c.Param("id")
	if err := entity.DB().Preload("Roomtype").Preload("Roomstate").Raw("SELECT * FROM rooms WHERE roomtype_id = ? and roomstate_id = 1 limit 1", id).Scan(&rooms).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rooms})
}

// GET /rooms
func ListRoom(c *gin.Context) {
	var rooms []entity.Room
	if err := entity.DB().Preload("Roomtype").Preload("Roomstate").Raw("SELECT * FROM rooms WHERE roomstate_id = 1").Scan(&rooms).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rooms})
}

// DELETE /rooms/:id
func DeleteRoom(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM rooms WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /rooms
func UpdateRoom(c *gin.Context) {
	var rooms entity.Room
	if err := c.ShouldBindJSON(&rooms); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", rooms.ID).First(&rooms); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rooms not found"})
		return
	}
	if err := entity.DB().Save(&rooms).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rooms})
}

//****************************************
//Repair //roomreturn
func GetRoomofRepairandroomreturn(c *gin.Context) {
	var room entity.Room
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM rooms WHERE id = ?", id).Scan(&room).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": room})
}

func ListRooms_RoomReturn(c *gin.Context) {
	var rooms []entity.Room
	if err := entity.DB().Raw("SELECT * FROM rooms").Scan(&rooms).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rooms})
}
