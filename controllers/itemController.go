//By Ahmad Fadhil Rizqi

package controllers

import (
	"jual_beli_barang_bekas/config"
	"jual_beli_barang_bekas/models"

	"github.com/gofiber/fiber/v2"
)

// CreateItem untuk menambah barang
func CreateItem(c *fiber.Ctx) error {
    var item models.Item
    if err := c.BodyParser(&item); err != nil {
        return c.Status(400).SendString(err.Error())
    }
    config.DB.Create(&item)
    return c.Status(201).JSON(item)
}

// ReadItem untuk mendapatkan item berdasarkan ID
func ReadItem(c *fiber.Ctx) error {
    id := c.Params("id")
    var item models.Item
    if err := config.DB.First(&item, id).Error; err != nil {
        return c.Status(404).SendString("Item not found")
    }
    return c.JSON(item)
}

// UpdateItem untuk memperbarui informasi item
func UpdateItem(c *fiber.Ctx) error {
    id := c.Params("id")
    var item models.Item
    if err := config.DB.First(&item, id).Error; err != nil {
        return c.Status(404).SendString("Item not found")
    }
    if err := c.BodyParser(&item); err != nil {
        return c.Status(400).SendString(err.Error())
    }
    config.DB.Save(&item)
    return c.JSON(item)
}

// DeleteItem untuk menghapus item
func DeleteItem(c *fiber.Ctx) error {
    id := c.Params("id")
    if err := config.DB.Delete(&models.Item{}, id).Error; err != nil {
        return c.Status(404).SendString("Item not found")
    }
    return c.SendStatus(204)
}