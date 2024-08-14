package middlewares

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/satyam-jha-16/event-manager/models"
	"gorm.io/gorm"
)

func AuthProtected(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			log.Warnf("empty authorization error")
			return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status": "fail",
				"error":  "Unauthorized",
			})
		}
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			log.Warnf("empty Authorization Error")
			return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status": "fail",
				"error":  "Unauthorized",
			})
		}
		tokenStr := tokenParts[1]

		tokenSecret := []byte(os.Getenv("JWT_SECRET"))
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.GetSigningMethod("HS256").Alg() {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return tokenSecret, nil
		})

		if err != nil || !token.Valid {
			log.Warnf("invalid token")

			return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status": "fail",
				"error":  "Unauthorized",
			})
		}

		userId := token.Claims.(jwt.MapClaims)["id"]

		if err := db.Model(&models.User{}).Where("id = ?", userId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warnf("user not found in the db")

			return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "fail",
				"error": "Unauthorized",
			})
		}

		c.Locals("userId", userId)

		return c.Next()

	}

}
