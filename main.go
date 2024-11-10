package main

import (
	"log"
	"os"
	"simpleapi/config"
	"simpleapi/models"
	"simpleapi/routes"
	"simpleapi/seeders"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Koneksi ke database
	config.ConnectDatabase()

	// Menjalankan migrasi
	config.DB.AutoMigrate(&models.Name{})
	config.DB.AutoMigrate(&models.Geolocation{})
	config.DB.AutoMigrate(&models.Address{})
	config.DB.AutoMigrate(&models.User{})

	// Menjalankan seeders
	seeders.Seed()

	// Membuat instance Fiber dan mengatur route
	app := fiber.New()
	routes.Routes(app)

	// Menjalankan server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
