//By Ahmad Fadhil Rizqi

package models

import "gorm.io/gorm"

// Item mewakili barang yang dijual
type Item struct {
    gorm.Model
    UserID      uint    `json:"user_id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    CategoryID  uint    `json:"category_id"`
    Price       float64 `json:"price"`
    Stock       int     `json:"stock"`

    // Relasi dengan tabel User dan Category
    User        User    `gorm:"foreignKey:UserID;references:ID"`
    Category    Category `gorm:"foreignKey:CategoryID;references:ID"`
}
