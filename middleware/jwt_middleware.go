package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtAuthentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid token %v", t.Header["alg"])
			}
			secretKey := os.Getenv("JWT_SECRET_KEY")
			return []byte(secretKey), nil
		})
		if err != nil {
			log.Println(err.Error())
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		cliams := token.Claims.(jwt.MapClaims)
		ctx.Set("cliams", cliams["username"])
		ctx.Next()
	}
}
