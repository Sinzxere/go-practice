package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Sinzxere/go-practice/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// โหลดไฟล์ .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}
}

func ConnectDatabase() {
	// ข้อมูลการเชื่อมต่อฐานข้อมูล
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// ถ้าไม่มีค่า environment variables ให้ใช้ค่าเริ่มต้น
	if username == "" {
		username = "postgres"
	}
	if password == "" {
		password = "postgres"
	}
	if dbName == "" {
		dbName = "go_api_db"
	}
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "5432"
	}

	// สร้าง connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		dbHost, username, password, dbName, dbPort)

	// เชื่อมต่อกับฐานข้อมูล
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migrate โมเดล
	db.AutoMigrate(&models.Item{})

	DB = db
	fmt.Println("Database connected successfully")
}
