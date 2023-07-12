package authutil

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const KEY = "asd132sd"

type JwtClaims struct {
	jwt.StandardClaims
	Username        string `json:"Username"`
	ApplicationName string
}

func GenerateJWTToken(userID, username, role string) (string, error) {
	// Membuat claim JWT
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token berlaku selama 1 hari
	}

	// Membuat token JWT dengan menggunakan secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte("secret-key") // Ganti dengan secret key Anda sendiri
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyAccessToken(tokenString string) (string, error) {
	claim := &JwtClaims{}
	t, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(KEY), nil
	})
	if err != nil {
		return "", fmt.Errorf("error on VerifyAccessToken : %w", err)
	}
	if !t.Valid {
		return "", fmt.Errorf("error on VerifyAccessToken : invalid token")
	}
	return claim.Username, nil
}
