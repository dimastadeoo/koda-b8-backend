package lib

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type JwtToken struct {
	UserId uint `json:"userId"`
	jwt.RegisteredClaims
}

// GeneratedToken membuat JWT token untuk user id
func GeneratedToken(userId uint) (string, error) {
	_ = godotenv.Load() // ignore error

	secret := os.Getenv("JWT_KEY")
	if secret == "" {
		return "", fmt.Errorf("JWT_KEY not set in environment")
	}

	expiryStr := os.Getenv("JWT_EXPIRY")
	expiry, err := time.ParseDuration(expiryStr)
	if err != nil {
		expiry = 15 * time.Minute // default jika tidak ada atau salah format
	}

	claims := JwtToken{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// VerifyToken memverifikasi token dan mengembalikan userId jika valid
func VerifyToken(tokenString string) (bool, uint, error) {
	secret := os.Getenv("JWT_KEY")
	if secret == "" {
		return false, 0, fmt.Errorf("JWT_KEY not set")
	}

	claims := &JwtToken{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return false, 0, err
	}
	if !token.Valid {
		return false, 0, fmt.Errorf("invalid token")
	}
	return true, claims.UserId, nil
}
