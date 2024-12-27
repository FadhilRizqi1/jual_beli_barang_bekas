//By Ahmad Fadhil Rizqi

package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// AuthMiddleware untuk memeriksa token JWT
func AuthMiddleware(c *fiber.Ctx) error {
    token := c.Get("Authorization")
    if token == "" {
        return c.Status(401).SendString("Unauthorized")
    }

    claims := jwt.MapClaims{}
    _, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte("secret_key"), nil // Ganti dengan kunci rahasia Anda
    })

    if err != nil {
        return c.Status(401).SendString("Unauthorized")
    }

    c.Locals("user_id", claims["user_id"])
    return c.Next()
}