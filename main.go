// By Ahmad Fadhil Rizqi
package main

import (
	"jual_beli_barang_bekas/config"
	"jual_beli_barang_bekas/models"
	"jual_beli_barang_bekas/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
    // Inisialisasi Fiber
    app := fiber.New()

// Menghubungkan ke database
    config.ConnectDatabase()

    // Migrasi model Category
    config.DB.AutoMigrate(&models.Category{}) // GORM akan membuat tabel Category jika belum ada

	// Migrasi model Item
	config.DB.AutoMigrate(&models.Item{}) // GORM akan membuat tabel Item jika belum ada

    // Migrasi model User
    config.DB.AutoMigrate(&models.User{}) // GORM akan membuat tabel User jika belum ada

	// Migrasi model Transaction
	config.DB.AutoMigrate(&models.Transaction{}) // GORM akan membuat tabel Transaction jika belum ada

    // Menyiapkan routes
    routes.SetupRoutes(app)

    // Menjalankan server
    app.Listen(":8080")
}