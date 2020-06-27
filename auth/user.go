package auth

import (
	"fmt"
	"strconv"

	"github.com/alcjohn/rest_gin/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) (*models.User, error) {
	token, err := VerifyToken(c.Request)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		tugudu := claims["exp"]
		fmt.Println(tugudu)
		var user models.User

		if err := models.DB.Where("id = ?", userID).First(&user).Error; err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, err
}
