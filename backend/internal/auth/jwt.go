package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(userID string) (string, error) {
	claims := &JwtCustomClaims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(JwtSecret)
}

func ParseToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JwtSecret), nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims["user_id"].(string), nil
	} else {
		return "", err
	}
}
