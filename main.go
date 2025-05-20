package main

import (
	"github.com/Sinzxere/go-practice/config"
	"github.com/Sinzxere/go-practice/routes"
)

func main() {
	// เชื่อมต่อฐานข้อมูล
	config.ConnectDatabase()

	// ตั้งค่า router
	r := routes.SetupRouter()

	// เริ่มต้น server ที่พอร์ต 8080
	r.Run(":8080")
}
