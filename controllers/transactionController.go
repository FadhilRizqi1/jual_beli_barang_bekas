//By Ahmad Fadhil Rizqi

package controllers

import (
	"jual_beli_barang_bekas/config"
	"jual_beli_barang_bekas/models"

	"github.com/gofiber/fiber/v2"
)

// CreateTransaction untuk mencatat transaksi
func CreateTransaction(c *fiber.Ctx) error {
    var transaction models.Transaction
    if err := c.BodyParser(&transaction); err != nil {
        return c.Status(400).SendString(err.Error())
    }
    config.DB.Create(&transaction)
    return c.Status(201).JSON(transaction)
}

// ReadTransaction untuk mendapatkan transaksi berdasarkan ID
func ReadTransaction(c *fiber.Ctx) error {
    id := c.Params("id")
    var transaction models.Transaction
    if err := config.DB.First(&transaction, id).Error; err != nil {
        return c.Status(404).SendString("Transaction not found")
    }
    return c.JSON(transaction)
}

// UpdateTransaction untuk memperbarui informasi transaksi
func UpdateTransaction(c *fiber.Ctx) error {
    id := c.Params("id")
    var transaction models.Transaction
    if err := config.DB.First(&transaction, id).Error; err != nil {
        return c.Status(404).SendString("Transaction not found")
    }
    if err := c.BodyParser(&transaction); err != nil {
        return c.Status(400).SendString(err.Error())
    }
    config.DB.Save(&transaction)
    return c.JSON(transaction)
}

// DeleteTransaction untuk menghapus transaksi
func DeleteTransaction(c *fiber.Ctx) error {
    id := c.Params("id")
    if err := config.DB.Delete(&models.Transaction{}, id).Error; err != nil {
        return c.Status(404).SendString("Transaction not found")
    }
    return c.SendStatus(204)
}