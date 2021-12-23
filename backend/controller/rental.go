package controller

import (
	"net/http"

	"github.com/tzsoulcap/ui-sa/entity"
	"github.com/gin-gonic/gin"
)

// POST /rental
func CreateRental(c *gin.Context) {

	var rental entity.Rental
	var tenant entity.Tenant
	var room entity.Room
	var rentalstate entity.Rentalstate

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร rental
	if err := c.ShouldBindJSON(&rental); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//10: ค้นหา User ด้วย id
	if tx := entity.DB().Where("id = ?", rental.TenantID).First(&tenant); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// 11: ค้นหา rentalstate ด้วย id
	if tx := entity.DB().Where("id = ?", rental.RentalstateID).First(&rentalstate); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rentalstate not found"})
		return
	}

	// 12: ค้นหา room ด้วย id
	if tx := entity.DB().Where("id = ? and roomstate_id = 1 ", rental.RoomID).First(&room); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}
	if tx := entity.DB().Model(&room).Where(rental.RoomID).Update("roomstate_id", 2); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}

	// 13: สร้าง rental
	wv := entity.Rental{
		DAYTIME:     rental.DAYTIME, //ตั้งค่าฟิลด์ DAYTIME
		CHACKIN:     rental.CHACKIN, //ตั้งค่าฟิลด์ CHACKIN
		Tenant:      tenant,         // โยงความสัมพันธ์กับ Entity user
		Room:        room,           // โยงความสัมพันธ์กับ Entity room
		Rentalstate: rentalstate,    // โยงความสัมพันธ์กับ Entity rentalstate
	}

	// 14: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})

}

// GET /Rental/:id
func GetRental(c *gin.Context) {
	var rental entity.Rental
	id := c.Param("id")
	if err := entity.DB().Preload("Tenant").Preload("Room").Preload("Rentalstate").Raw("SELECT * FROM rentals WHERE id = ?", id).Find(&rental).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rental})
}

// GET /Rental
func ListRental(c *gin.Context) {
	var rentals []entity.Rental
	if err := entity.DB().Preload("Tenant").Preload("Room").Preload("Rentalstate").Preload("Room.Roomtype").Raw("SELECT * FROM rentals").Find(&rentals).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rentals})
}

// DELETE /Rental/:id
func DeleteRental(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM rentals WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rental not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Rental
func UpdateRental(c *gin.Context) {
	var rental entity.Rental
	if err := c.ShouldBindJSON(&rental); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", rental.ID).First(&rental); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rental not found"})
		return
	}

	if err := entity.DB().Save(&rental).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rental})
}

//Repair
// GET /:id
func GetRentalofRepair(c *gin.Context) {
	var rental entity.Rental
	uid := c.Param("id")
	if err := entity.DB().Preload("Room").Raw("SELECT * FROM rentals WHERE tenant_id = ?", uid).Find(&rental).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rental})
}

//Roomretrun
func GetRentalofRoomretrun(c *gin.Context) {

	var rental []entity.Rental

	id := c.Param("id")

	if err := entity.DB().Preload("Room").Raw("SELECT * FROM rentals WHERE id = ?", id).Find(&rental).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": rental})

}
