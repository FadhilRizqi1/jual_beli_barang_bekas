//By Ahmad Fadhil Rizqi

package controllers

import (
	"jual_beli_barang_bekas/config"
	"jual_beli_barang_bekas/models"

	"github.com/gofiber/fiber/v2"
)

// CreateCategory untuk menambah kategori
func CreateCategory(c *fiber.Ctx) error {
    var category models.Category
    if err := c.BodyParser(&category); err != nil {
        return c.Status(400).SendString(err.Error())
    }

    // Simpan kategori ke database
    if err := config.DB.Create(&category).Error; err != nil {
        return c.Status(500).SendString("Gagal menambahkan kategori")
    }

    return c.Status(201).JSON(category)
}

// ReadCategory untuk mendapatkan kategori berdasarkan ID
func ReadCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return c.Status(404).SendString("Category not found")
	}
	return c.JSON(category)
}

// UpdateCategory untuk memperbarui informasi kategori
func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return c.Status(404).SendString("Category not found")
	}
	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	config.DB.Save(&category)
	return c.JSON(category)
}

// DeleteCategory untuk menghapus kategori
func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Category{}, id).Error; err != nil {
		return c.Status(404).SendString("Category not found")
	}
	return c.SendStatus(204)
}