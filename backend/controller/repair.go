package controller

import (
	"net/http"

	"github.com/tzsoulcap/ui-sa/entity"
	"github.com/gin-gonic/gin"
)

// POST /Repairs
func CreateRepair(c *gin.Context) {

	var repair entity.Repair
	var rental entity.Rental
	var tenant entity.Tenant
	var room_equipment entity.RoomEquipment

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 9 จะถูก bind เข้าตัวแปร repair
	if err := c.ShouldBindJSON(&repair); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 10: ค้นหา rental ด้วย id ผู้ใช้
	if tx := entity.DB().Where("tenant_id = ?", repair.TenantID).First(&rental); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rental not found"})
		return
	}

	// 11: ค้นหา tenant ด้วย id
	if tx := entity.DB().Where("id = ?", repair.TenantID).First(&tenant); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tenant not found"})
		return
	}

	// 12: ค้นหา room_equipment ด้วย id
	if tx := entity.DB().Where("id = ?", repair.RoomEquipmentID).First(&room_equipment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room equipment not found"})
		return
	}

	// 13: สร้าง Repair
	rp := entity.Repair{
		AddedTime:     repair.AddedTime, // ตั้งค่าฟิลด์ AddedTime
		Note:          repair.Note,      // ตั้งค่าฟิลด์ Note
		Tenant:        tenant,           // โยงความสัมพันธ์กับ Entity Tenant
		RoomEquipment: room_equipment,   // โยงความสัมพันธ์กับ Entity RoomEquipment
		Rental:        rental,           // โยงความสัมพันธ์กับ Entity Rental
	}

	// 14: บันทึก
	if err := entity.DB().Create(&rp).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rp})
}

// GET /Repair/:id
func GetRepair(c *gin.Context) {
	var repairs []entity.Repair
	id := c.Param("id")
	if err := entity.DB().Preload("Tenant").Preload("Rental").Preload("RoomEquipment").Preload("Rental.Room").Preload("RoomEquipment.Equipment").Raw("SELECT * FROM repairs WHERE tenant_id = ?", id).Find(&repairs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": repairs})
}

// GET /Repairs
func ListRepairs(c *gin.Context) {
	var repairs []entity.Repair
	if err := entity.DB().Preload("Tenant").Preload("Rental").Preload("RoomEquipment").Preload("Rental.Room").Preload("RoomEquipment.Equipment").Raw("SELECT * FROM repairs").Find(&repairs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": repairs})
}

// DELETE /Repairs/:id
func DeleteRepair(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM repairs WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "repair not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Repairs
func UpdateRepair(c *gin.Context) {
	var repair entity.Repair
	if err := c.ShouldBindJSON(&repair); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", repair.ID).First(&repair); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "repair not found"})
		return
	}

	if err := entity.DB().Save(&repair).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": repair})
}
