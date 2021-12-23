package controller

import (
	"net/http"

	"github.com/tzsoulcap/ui-sa/entity"
	"github.com/gin-gonic/gin"
)

// Post /Tenant
// List all Tenant
func CreateTenant(c *gin.Context) {

	var tenant entity.Tenant
	var prefix entity.Prefix
	var career entity.Career
	var gender entity.Gender
	var canton entity.Canton
	var employer entity.Employer

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 10 จะถูก bind เข้าตัวแปร tenant
	if err := c.ShouldBindJSON(&tenant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", tenant.OwnerID).First(&employer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employer not found"})
		return
	}

	// 11: ค้นหา prefix ด้วย id
	if tx := entity.DB().Where("id = ?", tenant.PrefixID).First(&prefix); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prefix not found"})
		return
	}

	// 12: ค้นหา career ด้วย id
	if tx := entity.DB().Where("id = ?", tenant.CareerID).First(&career); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "career not found"})
		return
	}

	// 13: ค้นหา gender ด้วย id
	if tx := entity.DB().Where("id = ?", tenant.GenderID).First(&gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}

	// 14: ค้นหา canton ด้วย id
	if tx := entity.DB().Where("id = ?", tenant.CantonID).First(&canton); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "canton not found"})
		return
	}

	// 15: สร้าง tenant
	tn := entity.Tenant{
		Owner:    employer,
		Prefix:   prefix,
		Career:   career,
		Gender:   gender,
		Canton:   canton,
		Name:     tenant.Name,
		Idcard:   tenant.Idcard,
		BirthDay: tenant.BirthDay,
		Age:      tenant.Age,
		Address:  tenant.Address,
		Village:  tenant.Village,
		Tel:      tenant.Tel,
		Email:    tenant.Email,
		Password: tenant.Password,
	}

	// 16: บันทึก
	if err := entity.DB().Create(&tn).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tn})

}

// GET /Tenant/:id
func GetTenant(c *gin.Context) {
	var tenant entity.Tenant
	id := c.Param("id")
	if err := entity.DB().Preload("Owner").Preload("Prefix").Preload("Career").Preload("Gender").Preload("Canton").Preload("Canton.Prefecture").Preload("Canton.Prefecture.Province").Find(&tenant).Raw("SELECT * FROM tenants WHERE id = ?", id).Find(&tenant).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tenant})
}

// GET /Tenant
func ListTenant(c *gin.Context) {
	var tenant []entity.Tenant
	if err := entity.DB().Preload("Owner").Preload("Prefix").Preload("Career").Preload("Gender").Preload("Canton").Preload("Canton.Prefecture").Preload("Canton.Prefecture.Province").Find(&tenant).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tenant})
}

// DELETE /Tenant/:id
func DeleteTenant(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM tenants WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tenant not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Tenant
func UpdateTenant(c *gin.Context) {
	var tenant entity.Tenant
	if err := c.ShouldBindJSON(&tenant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", tenant.ID).First(&tenant); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tenant not found"})
		return
	}

	if err := entity.DB().Save(&tenant).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tenant})
}
