package controllers

import (
	"github.com/alcjohn/rest_gin/models"
	"github.com/gin-gonic/gin"
)

func UsersRoutes(r *gin.RouterGroup) {
	var users []models.User
	r.GET("/", Paginate(&users))
}
