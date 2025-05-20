package main

import (
	"https://github.com/Sinzxere/go-practice/tree/main/config"
	"https://github.com/Sinzxere/go-practice/tree/main/routes"
)

func main() {
	// เชื่อมต่อฐานข้อมูล
	config.ConnectDatabase()

	// ตั้งค่า router
	r := routes.SetupRouter()

	// เริ่มต้น server ที่พอร์ต 8080
	r.Run(":8080")
}
