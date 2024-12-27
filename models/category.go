//By Ahmad Fadhil Rizqi

package models

import "gorm.io/gorm"

// Category mewakili kategori barang
type Category struct {
    gorm.Model // Ini sudah mencakup kolom ID, CreatedAt, UpdatedAt, DeletedAt
    CategoryName string `json:"category_name"`
}
