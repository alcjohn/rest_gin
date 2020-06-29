package middlewares

import (
	"fmt"
	"os"
	"strings"

	"github.com/alcjohn/rest_gin/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		if bearer == "" {
			c.Next()
			return
		}
		jwtParts := strings.Split(bearer, " ")
		if len(jwtParts) != 2 {
			c.Next()
			return
		}
		jwtEncoded := jwtParts[1]
		token, err := jwt.Parse(jwtEncoded, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("ACCESS_SECRET")), nil
		})
		if err != nil {
			println(err.Error())
			c.Next()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := uint(claims["user_id"].(float64))
			fmt.Printf("[+] Authenticated request, authenticated user id is %d\n", userID)

			var user models.User
			if userID != 0 {
				models.DB.First(&user, userID)
			}
			c.Set("AuthUser", user)
			c.Set("AuthUserID", userID)
		}

	}
}
