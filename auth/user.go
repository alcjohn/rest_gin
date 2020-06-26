package auth

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alcjohn/rest_gin/models"
	"github.com/dgrijalva/jwt-go"
)

func User(r *http.Request) (*models.User, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		var user models.User

		if err := models.DB.Where("id = ?", userID).First(&user).Error; err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, err
}
