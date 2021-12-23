package controller

import (
	"net/http"

	"github.com/tzsoulcap/ui-sa/entity"
	"github.com/gin-gonic/gin"
)

// POST /return_rooms
func CreateReturnRoom(c *gin.Context) {

	var returnroom entity.ReturnRoom
	var rental entity.Rental
	var room entity.Room
	var tenant entity.Tenant

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร returnRoom
	if err := c.ShouldBindJSON(&returnroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา tenant ด้วย id
	if tx := entity.DB().Where("id = ?", returnroom.TenantID).First(&tenant); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tenant not found"})
		return
	}

	// 10: ค้นหา room ด้วย id
	if tx := entity.DB().Where("id = ?", returnroom.RoomID).First(&room); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}

	// 11: ค้นหา rental ด้วย id
	if tx := entity.DB().Where("id = ?", returnroom.RentalID).First(&rental); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rental not found"})
		return
	}

	// 12: สร้าง ReturnRoom
	rt := entity.ReturnRoom{
		Tenant:      tenant,                 // โยงความสัมพันธ์กับ Entity Tenant
		Room:        room,                   // โยงความสัมพันธ์กับ Entity Room
		Rental:      rental,                 // โยงความสัมพันธ์กับ Entity Rental
		MoveOutTime: returnroom.MoveOutTime, // ตั้งค่าฟิลด์ MoveOutTime
	}

	// 13: บันทึก
	if err := entity.DB().Create(&rt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rt})

}

// GET /returnroom/:id
func GetReturnRoom(c *gin.Context) {
	var returnroom entity.ReturnRoom
	id := c.Param("id")
	if err := entity.DB().Preload("Tenant").Preload("Room").Preload("Rental").Raw("SELECT * FROM return_rooms WHERE id = ?", id).Find(&returnroom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": returnroom})
}

// GET /return_rooms
func ListReturnRooms(c *gin.Context) {
	var returnrooms []entity.ReturnRoom
	if err := entity.DB().Preload("Tenant").Preload("Room").Preload("Rental").Raw("SELECT * FROM return_rooms").Find(&returnrooms).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": returnrooms})
}

// DELETE /return_rooms/:id
func DeleteReturnRoom(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM return_rooms WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "returnroom not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /return_rooms
func UpdateReturnRoom(c *gin.Context) {
	var returnroom entity.ReturnRoom
	if err := c.ShouldBindJSON(&returnroom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", returnroom.ID).First(&returnroom); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "returnoom not found"})
		return
	}

	if err := entity.DB().Save(&returnroom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": returnroom})
}
