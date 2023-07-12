package utils

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// ExtractTokenFromAuthHeader mengambil token JWT dari header Authorization dalam format "Bearer <token>"
func ExtractTokenFromAuthHeader(authHeader string) (string, error) {
	// Mengecek format header Authorization
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("invalid authorization header format")
	}

	// Mengambil token dari header Authorization
	token := strings.TrimPrefix(authHeader, "Bearer ")

	return token, nil
}

// VerifyJWTToken memverifikasi token JWT dan mengembalikan klaim JWT jika token valid
func VerifyJWTToken(tokenString string) (jwt.MapClaims, error) {
	// Menentukan fungsi validasi token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifikasi menggunakan HMAC, RSA, atau algoritma validasi lainnya
		// Pastikan kunci rahasia (secret key) sesuai dengan yang digunakan saat pembuatan token
		// Misalnya, untuk validasi menggunakan HMAC dengan algoritma HS256:
		secretKey := []byte("secret-key")
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Mengecek apakah token valid dan memiliki klaim
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
