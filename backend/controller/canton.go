package controller

import (
	"net/http"

	"github.com/tzsoulcap/ui-sa/entity"
	"github.com/gin-gonic/gin"
)

// POST /Canton
func CreateCanton(c *gin.Context) {
	var canton entity.Canton
	if err := c.ShouldBindJSON(&canton); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&canton).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": canton})
}

// GET /Canton/:id
func GetCanton(c *gin.Context) {
	var canton []entity.Canton
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM cantons WHERE prefecture_id = ?", id).Scan(&canton).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": canton})
}

// GET /Canton
func ListCanton(c *gin.Context) {
	var canton []entity.Canton
	if err := entity.DB().Raw("SELECT * FROM cantons").Scan(&canton).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": canton})
}

// DELETE /Canton/:id
func DeleteCanton(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM cantons WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "canton not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Canton
func UpdateCanton(c *gin.Context) {
	var canton entity.Canton
	if err := c.ShouldBindJSON(&canton); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", canton.ID).First(&canton); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "canton not found"})
		return
	}

	if err := entity.DB().Save(&canton).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": canton})
}
