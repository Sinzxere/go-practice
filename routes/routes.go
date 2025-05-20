package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/go-api-project/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// กลุ่ม routes สำหรับ API v1
	v1 := r.Group("/api/v1")
	{
		items := v1.Group("/items")
		{
			items.GET("/", controllers.GetItems)
			items.POST("/", controllers.CreateItem)
			items.GET("/:id", controllers.GetItem)
			items.PUT("/:id", controllers.UpdateItem)
			items.DELETE("/:id", controllers.DeleteItem)
		}
	}

	return r
}
