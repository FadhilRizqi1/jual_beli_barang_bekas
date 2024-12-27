//By Ahmad Fadhil Rizqi

package models

import "gorm.io/gorm"

// User mewakili pengguna di sistem
type User struct {
    gorm.Model // Ini sudah mencakup kolom ID, CreatedAt, UpdatedAt
    Name     string `json:"name"`
    Email    string `json:"email" gorm:"unique"`
    Password string `json:"password"`
}