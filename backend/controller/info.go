package controller

import (
	"github.com/tzsoulcap/ui-sa/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /infos

func CreateInfo(c *gin.Context) {

	var info entity.PromotionInfo

	if err := c.ShouldBindJSON(&info); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&info).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": info})

}

// GET /info/:id

func GetInfo(c *gin.Context) {

	var info entity.PromotionInfo

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM promotion_infos WHERE id = ?", id).Scan(&info).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": info})

}

// GET /infos

func ListInfo(c *gin.Context) {

	var infos []entity.PromotionInfo

	if err := entity.DB().Raw("SELECT * FROM promotion_infos").Scan(&infos).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": infos})

}

// DELETE /infos/:id

func DeleteInfo(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM promotion_infos WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "info not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /infos

func UpdateInfo(c *gin.Context) {

	var info entity.PromotionType

	if err := c.ShouldBindJSON(&info); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", info.ID).First(&info); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "info not found"})

		return

	}

	if err := entity.DB().Save(&info).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": info})

}
