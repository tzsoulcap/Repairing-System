package controller

import (
	"net/http"

	"github.com/tzsoulcap/ui-sa/entity"
	"github.com/gin-gonic/gin"
)

// POST /rentalstate
func CreateRentalstate(c *gin.Context) {
	var rentalstate entity.Rentalstate
	if err := c.ShouldBindJSON(&rentalstate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	if err := entity.DB().Create(&rentalstate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rentalstate})
}

// GET /rentalstate/:id
func GetRentalstate(c *gin.Context) {
	var rentalstate entity.Rentalstate
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM rentalstates WHERE id = ?", id).Scan(&rentalstate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rentalstate})
}

// GET /rentalstate
func ListRentalstate(c *gin.Context) {
	var rentalstate []entity.Rentalstate
	if err := entity.DB().Raw("SELECT * FROM rentalstates").Scan(&rentalstate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rentalstate})
}

// DELETE /rentalstate/:id
func DeleteRentalstate(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM rentalstates WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /rentalstate
func UpdateRentalstate(c *gin.Context) {
	var rentalstate entity.Rentalstate
	if err := c.ShouldBindJSON(&rentalstate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", rentalstate.ID).First(&rentalstate); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rentalstate not found"})
		return
	}
	if err := entity.DB().Save(&rentalstate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rentalstate})
}
