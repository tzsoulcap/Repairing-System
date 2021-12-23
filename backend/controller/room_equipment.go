package controller

import (
	"net/http"

	"github.com/tzsoulcap/ui-sa/entity"
	"github.com/gin-gonic/gin"
)

// POST /RoomEquipments
func CreateRoomEquipment(c *gin.Context) {
	var room_equipment entity.RoomEquipment
	var equipment entity.Equipment
	var room entity.Room
	if err := c.ShouldBindJSON(&room_equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", room_equipment.RoomID).First(&room); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", room_equipment.EquipmentID).First(&equipment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "equipment not found"})
		return
	}
	re := entity.RoomEquipment{
		Room:      room,
		Equipment: equipment,
	}
	if err := entity.DB().Create(&re).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": re})
}

// GET /RoomEquipment/:id
func GetRoomEquipment(c *gin.Context) {
	var room_equipment entity.RoomEquipment
	id := c.Param("id")
	if err := entity.DB().Preload("Room").Preload("Equipment").Raw("SELECT * FROM room_equipments WHERE id = ?", id).Find(&room_equipment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": room_equipment})
}

// GET /RoomEquipments
func ListRoomEquipments(c *gin.Context) {
	var room_equipments []entity.RoomEquipment
	number := c.Param("number")
	typeid := c.Param("typeid")
	if err := entity.DB().Preload("Room").Preload("Equipment").Raw("SELECT room_equipments.* FROM room_equipments INNER JOIN equipment  INNER JOIN rooms ON room_equipments.equipment_id = equipment.id AND room_equipments.room_id = rooms.id WHERE rooms.room_number = ? AND equipment.type_equipment_id = ?", number, typeid).Find(&room_equipments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": room_equipments})
}

// DELETE /RoomEquipments/:id
func DeleteRoomEquipment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM room_equipments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room equipment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /RoomEquipments
func UpdateRoomEquipment(c *gin.Context) {
	var room_equipment entity.RoomEquipment
	if err := c.ShouldBindJSON(&room_equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", room_equipment.ID).First(&room_equipment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room equipment not found"})
		return
	}

	if err := entity.DB().Save(&room_equipment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": room_equipment})
}
