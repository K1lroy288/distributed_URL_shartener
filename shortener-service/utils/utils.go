package utils

import (
	"crypto/rand"
	"fmt"
	"shortener-service/config"
	"strings"

	"github.com/golang-jwt/jwt"
)

func GenerateShortCode() (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const idLength = 6

	bytes := make([]byte, idLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	var res strings.Builder
	res.Grow(idLength)
	for i := range idLength {
		res.WriteByte(charset[bytes[i]%byte(len(charset))])
	}

	return res.String(), nil
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	cfg := config.GetConfig()
	jwtKey := []byte(cfg.JwtSecret)

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("wrong authentication method")
		}

		return jwtKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error jwt token parse: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("wronk token")
}
