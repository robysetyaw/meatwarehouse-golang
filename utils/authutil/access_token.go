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

func GenerateToken(userName string) (string, error) {
	now := time.Now().UTC()
	end := now.Add(1 * time.Hour)
	claim := &JwtClaims{
		Username: userName,
	}
	claim.IssuedAt = now.Unix()
	claim.ExpiresAt = end.Unix()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := t.SignedString([]byte(KEY))
	if err != nil {
		return "", fmt.Errorf("error on GenerateToken : %w", err)
	}
	return token, nil
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
