//By Ahmad Fadhil Rizqi

package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB adalah variabel global untuk koneksi database
var DB *gorm.DB

// ConnectDatabase menghubungkan ke database MySQL
func ConnectDatabase() {
    dsn := "root:@tcp(127.0.0.1:3306)/jual_beli_barang_bekas?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Gagal terhubung ke database:", err)
    }
    fmt.Println("Berhasil terhubung ke database")
}