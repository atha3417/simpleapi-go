package main

import (
	"log"
	"os"
	db "simpleapi/database"
	"simpleapi/database/seeders"
	"simpleapi/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Koneksi ke database
	db.ConnectDatabase()

	// Menjalankan migrasi
	db.Migrate()

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
