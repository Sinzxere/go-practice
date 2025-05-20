package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/go-api-project/config"
	"github.com/yourusername/go-api-project/models"
)

// GetItems godoc
// @Summary Get all items
// @Description Get a list of all items
// @Tags items
// @Accept json
// @Produce json
// @Success 200 {array} models.Item
// @Router /items [get]
func GetItems(c *gin.Context) {
	var items []models.Item
	config.DB.Find(&items)

	c.JSON(http.StatusOK, gin.H{"data": items})
}

// CreateItem godoc
// @Summary Create a new item
// @Description Create a new item with the input payload
// @Tags items
// @Accept json
// @Produce json
// @Param item body models.Item true "Create item"
// @Success 201 {object} models.Item
// @Failure 400 {object} string
// @Router /items [post]
func CreateItem(c *gin.Context) {
	var input models.Item

	// ตรวจสอบความถูกต้องของข้อมูลที่รับเข้ามา
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// สร้างรายการใหม่
	item := models.Item{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}

	// บันทึกลงฐานข้อมูล
	result := config.DB.Create(&item)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": item})
}

// GetItem godoc
// @Summary Get an item by ID
// @Description Get an item by its ID
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} models.Item
// @Failure 404 {object} string
// @Router /items/{id} [get]
func GetItem(c *gin.Context) {
	var item models.Item

	// ค้นหารายการตาม ID
	if err := config.DB.First(&item, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": item})
}

// UpdateItem godoc
// @Summary Update an item
// @Description Update an item with the input payload
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param item body models.Item true "Update item"
// @Success 200 {object} models.Item
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /items/{id} [put]
func UpdateItem(c *gin.Context) {
	var item models.Item

	// ค้นหารายการตาม ID
	if err := config.DB.First(&item, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	// ตรวจสอบความถูกต้องของข้อมูลที่รับเข้ามา
	var input models.Item
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// อัปเดตข้อมูล
	config.DB.Model(&item).Updates(models.Item{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	})

	c.JSON(http.StatusOK, gin.H{"data": item})
}

// DeleteItem godoc
// @Summary Delete an item
// @Description Delete an item by its ID
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} string
// @Failure 404 {object} string
// @Router /items/{id} [delete]
func DeleteItem(c *gin.Context) {
	var item models.Item

	// ค้นหารายการตาม ID
	if err := config.DB.First(&item, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	// ลบรายการ
	config.DB.Delete(&item)

	c.JSON(http.StatusOK, gin.H{"data": "Item deleted successfully"})
}
