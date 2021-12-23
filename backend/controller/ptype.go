package controller

import (
	"github.com/tzsoulcap/ui-sa/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /ptypes

func CreatePType(c *gin.Context) {

	var ptype entity.PromotionType

	if err := c.ShouldBindJSON(&ptype); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&ptype).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": ptype})

}

// GET /ptype/:id

func GetPtype(c *gin.Context) {

	var ptype entity.PromotionType

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM promotion_types WHERE id = ?", id).Scan(&ptype).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": ptype})

}

// GET /ptypes

func ListPtype(c *gin.Context) {

	var ptypes []entity.PromotionType

	if err := entity.DB().Raw("SELECT * FROM promotion_types").Scan(&ptypes).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": ptypes})

}

// DELETE /ptypes/:id

func DeletePtype(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM promotion_types WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "type not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /ptypes

func UpdatePtype(c *gin.Context) {

	var ptype entity.PromotionType

	if err := c.ShouldBindJSON(&ptype); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", ptype.ID).First(&ptype); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "type not found"})

		return

	}

	if err := entity.DB().Save(&ptype).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": ptype})

}
