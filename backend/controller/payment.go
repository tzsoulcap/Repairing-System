package controller

import (
	"github.com/tzsoulcap/ui-sa/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /payments
func CreatePayment(c *gin.Context) {

	var payment entity.Payment
	var tenant entity.Tenant
	var promotionActive entity.PromotionActive
	var rental entity.Rental

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร active
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา Tenant ด้วย id
	if tx := entity.DB().Where("id = ?", payment.TenantID).First(&tenant); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tenant not found"})
		return
	}
	// 11: ค้นหา PromotionActive ด้วย id
	if tx := entity.DB().Where("id = ?", payment.PromotionActiveID).First(&promotionActive); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PromotionActive not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", payment.RentalID).First(&rental); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rental not found"})
		return
	}
	// 12: สร้าง Payment
	pm := entity.Payment{
		Tenant:          tenant,
		PromotionActive: promotionActive,
		Rental:          rental,
		Paydate:         payment.Paydate,
	}

	// 13: บันทึก
	if err := entity.DB().Create(&pm).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": pm})
}

// GET /payment/:id
func GetPayment(c *gin.Context) {
	var payment entity.Payment
	id := c.Param("id")
	if err := entity.DB().Preload("Tenant").Preload("PromotionActive").Preload("Rental").Raw("SELECT * FROM payments WHERE id = ?", id).Find(&payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payment})
}

// GET /payments
func ListPayments(c *gin.Context) {
	var payment []entity.Payment
	if err := entity.DB().Preload("Tenant").Preload("PromotionActive").Preload("Rental").Preload("Rental.Room").Preload("PromotionActive.Info").Raw("SELECT * FROM payments").Find(&payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payment})
}

// DELETE /payments/:id
func DeletePayment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM payments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment pro not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /payments
func UpdatePayment(c *gin.Context) {
	var payment entity.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", payment.ID).First(&payment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment not found"})
		return
	}

	if err := entity.DB().Save(&payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payment})
}
