package routes

import (
	"github.com/Sinzxere/go-practice/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// กลุ่ม routes สำหรับ API v1
	v1 := r.Group("/api/v1")
	{
		// Routes สำหรับ items (ถ้ามี)
		items := v1.Group("/items")
		{
			items.GET("/", controllers.GetItems)
			items.POST("/", controllers.CreateItem)
			items.GET("/:id", controllers.GetItem)
			items.PUT("/:id", controllers.UpdateItem)
			items.DELETE("/:id", controllers.DeleteItem)
		}

		// Routes สำหรับ app migrations
		appMigrations := v1.Group("/app-migrations")
		{
			appMigrations.GET("/", controllers.GetAppMigrations)
			appMigrations.GET("/:app_name", controllers.GetAppMigrationByName)
			appMigrations.GET("/migrated", controllers.GetMigratedApps)
			appMigrations.GET("/recheck", controllers.GetAppsNeedingRecheck)
			appMigrations.POST("/", controllers.CreateAppMigration)
		}
	}

	return r
}
