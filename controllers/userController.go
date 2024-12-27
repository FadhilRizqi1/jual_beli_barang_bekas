//By Ahmad Fadhil Rizqi

package controllers

import (
	"fmt"
	"jual_beli_barang_bekas/config"
	"jual_beli_barang_bekas/models"

	"github.com/gofiber/fiber/v2"
)

// CreateUser untuk membuat pengguna baru
func CreateUser (c *fiber.Ctx) error {
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(400).SendString(err.Error())
    }
    user.Password = "hashed_password" 
    result := config.DB.Create(&user)
    if result.Error != nil {
        return c.Status(500).SendString("Gagal menambahkan pengguna")
    }
    return c.Status(201).JSON(user)
}

// ReadUser  untuk mendapatkan pengguna berdasarkan ID
func ReadUser (c *fiber.Ctx) error {
    id := c.Params("id")
    var user models.User
    if err := config.DB.First(&user, id).Error; err != nil {
        return c.Status(404).SendString("User  not found")
    }
    return c.JSON(user)
}

// UpdateUser  untuk memperbarui informasi pengguna
func UpdateUser (c *fiber.Ctx) error {
    id := c.Params("id")
    var user models.User
    if err := config.DB.First(&user, id).Error; err != nil {
        return c.Status(404).SendString("User  not found")
    }
    if err := c.BodyParser(&user); err != nil {
        return c.Status(400).SendString(err.Error())
    }
    config.DB.Save(&user)
    return c.JSON(user)
}

// DeleteUser  untuk menghapus pengguna
func DeleteUser (c *fiber.Ctx) error {
    id := c.Params("id")
    fmt.Println("Deleting user with ID:", id) // Log ID

    // Menghapus pengguna dari database
    if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
        fmt.Println("Error deleting user:", err) // Log kesalahan
        return c.Status(404).SendString("User  not found")
    }

    return c.SendStatus(204) // No Content
}