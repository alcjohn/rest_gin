package controllers

import (
	"net/http"

	"github.com/alcjohn/rest_gin/auth"
	"github.com/alcjohn/rest_gin/dto"
	"github.com/alcjohn/rest_gin/models"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func AuthRoutes(r *gin.RouterGroup) {
	var controller AuthController
	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	r.GET("/me", controller.Me)
}

func (controller *AuthController) Login(c *gin.Context) {
	var input dto.LoginInput
	var user models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Email"})
		return
	}

	if err := user.VerifyPassword(input.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Password"})
		return
	}
	token, err := auth.CreateToken(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func (controller *AuthController) Register(c *gin.Context) {
	var input dto.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{Email: input.Email, Password: input.Password}
	err := models.DB.Create(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	c.JSON(http.StatusOK, user.Format())
}

func (controller *AuthController) Me(c *gin.Context) {
	user := c.Keys["AuthUser"].(models.User)
	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	c.JSON(http.StatusOK, user.Format())
}
