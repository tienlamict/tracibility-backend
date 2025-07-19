package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string // Đã được hash
}

var jwtSecret = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(jwtSecret)
}
