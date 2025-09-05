package utils

import (
	"mentors/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID uint, cfg *config.Config) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}
