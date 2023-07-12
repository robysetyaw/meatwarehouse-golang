package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"enigmacamp.com/final-project/team-4/track-prosto/utils/authutil"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func RequireToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// check exist token
		h := &authHeader{}
		if err := ctx.ShouldBindHeader(&h); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Unauthorize",
			})
			ctx.Abort()
			return
		}

		tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)

		// check token kosong
		fmt.Println(tokenString)
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Unauthorize",
			})
			ctx.Abort()
			return
		}

		// check verify token
		token, err := authutil.VerifyAccessToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Unauthorize",
			})
			ctx.Abort()
			return
		}

		fmt.Println(token)

		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "Unauthorize",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
