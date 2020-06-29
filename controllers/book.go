package controllers

import (
	"net/http"

	"github.com/alcjohn/rest_gin/dto"
	"github.com/alcjohn/rest_gin/models"
	"github.com/gin-gonic/gin"
)

type BooksController struct{}

func BooksRoutes(r *gin.RouterGroup) {
	var controller BooksController
	var users []models.User
	var user models.User
	r.GET("/", Paginate(&users))
	r.GET("/:id", Show(&user))
	r.POST("/", controller.Create)
	r.PATCH("/:id", controller.Update)
	r.DELETE("/:id", Delete(&user))
}

func (controller *BooksController) Create(c *gin.Context) {
	// Validate input
	var input dto.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func (controller *BooksController) Update(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input dto.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
