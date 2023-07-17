package middleware

import (
	"net/http"

	"enigmacamp.com/final-project/team-4/track-prosto/delivery/utils"
	"github.com/gin-gonic/gin"
)



func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mendapatkan token dari header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		token, err := utils.ExtractTokenFromAuthHeader(authHeader)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		claims, err := utils.VerifyJWTToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token or expired"})
			c.Abort()
			return
		}

		// Memeriksa peran pengguna
		role, ok := claims["role"].(string)
		if !ok || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. Admin role required"})
			c.Abort()
			return
		}

		// Menyimpan klaim dalam konteks untuk digunakan oleh handler API
		c.Set("claims", claims)

		c.Next()
	}
}

func JSONMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Content-Type", "application/json")
        c.Next()
    }
}