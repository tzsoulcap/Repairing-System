package controller

import (
	"github.com/tzsoulcap/ui-sa/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /actives
func CreateActive(c *gin.Context) {

	var active entity.PromotionActive
	var fest entity.Festival
	var info entity.PromotionInfo
	var employer entity.Employer

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร active
	if err := c.ShouldBindJSON(&active); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา Festival ด้วย id
	if tx := entity.DB().Where("id = ?", active.FestID).First(&fest); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "festival not found"})
		return
	}

	// 10: ค้นหา PromotionInfo ด้วย id
	if tx := entity.DB().Where("id = ?", active.InfoID).First(&info); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Infomation not found"})
		return
	}

	// 11: ค้นหา user ด้วย id
	if tx := entity.DB().Where("id = ?", active.EmployerID).First(&employer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}
	// 12: สร้าง PromotionActive
	ac := entity.PromotionActive{
		Employer:      employer,
		Fest:          fest,
		Info:          info,
		DurationStart: active.DurationStart,
		DurationEnd:   active.DurationEnd,
	}

	// 13: บันทึก
	if err := entity.DB().Create(&ac).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ac})
}

// GET /active/:id
func GetActive(c *gin.Context) {
	var active entity.PromotionActive
	id := c.Param("id")
	if err := entity.DB().Preload("Fest").Preload("Info").Preload("Employer").Raw("SELECT * FROM promotion_actives WHERE id = ?", id).Find(&active).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": active})
}

// GET /actives
func ListActives(c *gin.Context) {
	var active []entity.PromotionActive
	if err := entity.DB().Preload("Fest").Preload("Info").Preload("Employer").Raw("SELECT * FROM promotion_actives").Find(&active).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": active})
}

// DELETE /actives/:id
func DeleteActive(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM promotion_actives WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "active pro not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /actives
func UpdateActive(c *gin.Context) {
	var active entity.PromotionActive
	if err := c.ShouldBindJSON(&active); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", active.ID).First(&active); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "active not found"})
		return
	}

	if err := entity.DB().Save(&active).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": active})
}
