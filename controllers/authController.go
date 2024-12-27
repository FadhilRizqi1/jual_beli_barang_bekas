//By Ahmad Fadhil Rizqi

package controllers

import (
	"jual_beli_barang_bekas/config"
	"jual_beli_barang_bekas/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Register untuk registrasi pengguna baru
func Register(c *fiber.Ctx) error {
    var user models.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(400).SendString(err.Error())
    }

    // Hash password
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
    user.Password = string(hashedPassword)

    result := config.DB.Create(&user)
    if result.Error != nil {
        return c.Status(500).SendString("Gagal menambahkan pengguna")
    }
    return c.Status(201).JSON(user)
}

// Login untuk autentikasi pengguna
func Login(c *fiber.Ctx) error {
    var user models.User
    var foundUser  models.User

    if err := c.BodyParser(&user); err != nil {
        return c.Status(400).SendString(err.Error())
    }

    // Mencari pengguna berdasarkan email
    if err := config.DB.Where("email = ?", user.Email).First(&foundUser ).Error; err != nil {
        return c.Status(404).SendString("User  not found")
    }

    // Memeriksa password
    if err := bcrypt.CompareHashAndPassword([]byte(foundUser .Password), []byte(user.Password)); err != nil {
        return c.Status(401).SendString("Invalid password")
    }

    // Generate JWT
    token, err := GenerateJWT(foundUser .ID)
    if err != nil {
        return c.Status(500).SendString("Gagal membuat token")
    }

    return c.JSON(fiber.Map{
        "token": token,
    })
}

// GenerateJWT untuk menghasilkan token JWT
func GenerateJWT(userID uint) (string, error) {
    claims := jwt.MapClaims{}
    claims["user_id"] = userID
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
     return token.SignedString([]byte("secret_key"))
}

// ValidateToken untuk memvalidasi token JWT
func ValidateToken(c *fiber.Ctx) error {
    userID := c.Locals("user_id")
    if userID == nil {
        return c.Status(401).JSON(fiber.Map{
            "valid": false,
            "message": "Token is invalid",
        })
    }
    return c.JSON(fiber.Map{
        "valid": true,
        "user_id": userID,
        "message": "Token is valid",
    })
}