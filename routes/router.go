//By Ahmad Fadhil Rizqi

package routes

import (
	"jual_beli_barang_bekas/controllers"
	"jual_beli_barang_bekas/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    // Routes untuk User
    app.Post("/users", controllers.CreateUser ) // Endpoint untuk membuat pengguna
    app.Get("/users/:id", controllers.ReadUser ) // Endpoint untuk membaca informasi pengguna
    app.Put("/users/:id", controllers.UpdateUser ) // Endpoint untuk memperbarui informasi pengguna
    app.Delete("/users/:id", controllers.DeleteUser ) // Endpoint untuk menghapus pengguna

    // Routes untuk Auth
    app.Post("/register", controllers.Register) // Endpoint untuk registrasi
    app.Post("/login", controllers.Login)       // Endpoint untuk login
    app.Get("/validate", middleware.AuthMiddleware, controllers.ValidateToken) // Endpoint untuk memverifikasi token

    // Routes untuk Item
    app.Post("/items", middleware.AuthMiddleware, controllers.CreateItem) // Endpoint untuk membuat item
    app.Get("/items/:id", controllers.ReadItem) // Endpoint untuk membaca item
    app.Put("/items/:id", middleware.AuthMiddleware, controllers.UpdateItem) // Endpoint untuk memperbarui item
    app.Delete("/items/:id", middleware.AuthMiddleware, controllers.DeleteItem) // Endpoint untuk menghapus item

    // Routes untuk Transaction
    app.Post("/transactions", middleware.AuthMiddleware, controllers.CreateTransaction) // Endpoint untuk membuat transaksi
    app.Get("/transactions/:id", controllers.ReadTransaction) // Endpoint untuk membaca transaksi
    app.Put("/transactions/:id", middleware.AuthMiddleware, controllers.UpdateTransaction) // Endpoint untuk memperbarui transaksi
    app.Delete("/transactions/:id", middleware.AuthMiddleware, controllers.DeleteTransaction) // Endpoint untuk menghapus transaksi

    // Routes untuk Category
    app.Post("/categories", middleware.AuthMiddleware, controllers.CreateCategory) // Endpoint untuk membuat kategori
    app.Get("/categories/:id", controllers.ReadCategory) // Endpoint untuk membaca kategori
    app.Put("/categories/:id", middleware.AuthMiddleware, controllers.UpdateCategory) // Endpoint untuk memperbarui kategori
    app.Delete("/categories/:id", middleware.AuthMiddleware, controllers.DeleteCategory) // Endpoint untuk menghapus kategori
}