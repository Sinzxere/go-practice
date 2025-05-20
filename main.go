package main

import (
	"github.com/yourusername/go-api-project/config"
	"github.com/yourusername/go-api-project/routes"
)

func main() {
	// เชื่อมต่อฐานข้อมูล
	config.ConnectDatabase()

	// ตั้งค่า router
	r := routes.SetupRouter()

	// เริ่มต้น server ที่พอร์ต 8080
	r.Run(":8080")
}
