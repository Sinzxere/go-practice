package controllers

import (
	"net/http"

	"github.com/Sinzxere/go-practice/config"
	"github.com/Sinzxere/go-practice/models"
	"github.com/gin-gonic/gin"
)

// GetAppMigrations godoc
// @Summary Get all app migrations
// @Description Get a list of all app migrations
// @Tags app-migrations
// @Accept json
// @Produce json
// @Success 200 {array} models.AppMigration
// @Router /app-migrations [get]
func GetAppMigrations(c *gin.Context) {
	var appMigrations []models.AppMigration
	result := config.DB.Find(&appMigrations)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch app migrations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": appMigrations})
}

// GetAppMigrationByName godoc
// @Summary Get app migration by name
// @Description Get an app migration by its name
// @Tags app-migrations
// @Accept json
// @Produce json
// @Param app_name path string true "App Name"
// @Success 200 {object} models.AppMigration
// @Failure 404 {object} string
// @Router /app-migrations/{app_name} [get]
func GetAppMigrationByName(c *gin.Context) {
	var appMigration models.AppMigration
	appName := c.Param("app_name")

	result := config.DB.Where("app_name = ?", appName).First(&appMigration)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "App migration not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": appMigration})
}

// GetMigratedApps godoc
// @Summary Get all migrated apps
// @Description Get a list of all apps that have been migrated
// @Tags app-migrations
// @Accept json
// @Produce json
// @Success 200 {array} models.AppMigration
// @Router /app-migrations/migrated [get]
func GetMigratedApps(c *gin.Context) {
	var appMigrations []models.AppMigration
	result := config.DB.Where("migrated = ?", true).Find(&appMigrations)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch migrated apps"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": appMigrations})
}

// GetAppsNeedingRecheck godoc
// @Summary Get all apps needing recheck
// @Description Get a list of all apps that need rechecking
// @Tags app-migrations
// @Accept json
// @Produce json
// @Success 200 {array} models.AppMigration
// @Router /app-migrations/recheck [get]
func GetAppsNeedingRecheck(c *gin.Context) {
	var appMigrations []models.AppMigration
	result := config.DB.Where("recheck = ?", true).Find(&appMigrations)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch apps needing recheck"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": appMigrations})
}

// CreateAppMigration godoc
// @Summary Create a new app migration
// @Description Create a new app migration with the input payload
// @Tags app-migrations
// @Accept json
// @Produce json
// @Param app-migration body models.AppMigration true "Create app migration"
// @Success 201 {object} models.AppMigration
// @Failure 400 {object} string
// @Router /app-migrations [post]
func CreateAppMigration(c *gin.Context) {
	var input models.AppMigration

	// ตรวจสอบความถูกต้องของข้อมูลที่รับเข้ามา
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบว่ามีข้อมูลในตารางหรือไม่
	var count int64
	config.DB.Model(&models.AppMigration{}).Count(&count)

	// ถ้าไม่มีข้อมูลในตาราง (count = 0) ให้รีเซ็ต sequence
	if count == 0 {
		// สำหรับ PostgreSQL ใช้คำสั่ง ALTER SEQUENCE
		// หมายเหตุ: ชื่อ sequence อาจต้องปรับเปลี่ยนตาม schema ของคุณ
		config.DB.Exec("ALTER SEQUENCE app_migrations_id_seq RESTART WITH 1")
	}

	// สร้าง app migration ใหม่
	appMigration := models.AppMigration{
		AppName:  input.AppName,
		Migrated: input.Migrated,
		Restored: input.Restored,
		Recheck:  input.Recheck,
	}

	// บันทึกลงฐานข้อมูล
	result := config.DB.Create(&appMigration)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": appMigration})
}
