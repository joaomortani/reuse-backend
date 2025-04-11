package utils

import (
	"reuse-api/config"
	models "reuse-api/models/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokens(user *models.User) (string, string, error) {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	refreshClaim := jwt.MapClaims{
		"sub":     user.ID,
		"exp":     time.Now().Add(time.Hour * 100).Unix(),
		"refresh": true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim)
	secret := config.GetEnv("JWT_SECRET")

	accessToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", "", err
	}
	refreshTokenString, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshTokenString, nil
}
