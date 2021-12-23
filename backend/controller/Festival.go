package controller

import (
	"github.com/tzsoulcap/ui-sa/entity"
	"github.com/gin-gonic/gin"

	"net/http"
)

// GET /fests

func ListFests(c *gin.Context) {

	var fests []entity.Festival

	if err := entity.DB().Raw("SELECT * FROM festivals").Scan(&fests).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": fests})

}

// GET /fest/:id

func GetFest(c *gin.Context) {

	var fest entity.Festival

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM festivals WHERE id = ?", id).Scan(&fest).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": fest})

}

// POST /fests

func CreateFest(c *gin.Context) {

	var fest entity.Festival

	if err := c.ShouldBindJSON(&fest); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&fest).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": fest})

}

// PATCH /fests

func UpdateFest(c *gin.Context) {

	var fest entity.Festival

	if err := c.ShouldBindJSON(&fest); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", fest.ID).First(&fest); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "festival not found"})

		return

	}

	if err := entity.DB().Save(&fest).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": fest})

}

// DELETE /fests/:id

func DeleteFest(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM festivals WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "festival not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}
