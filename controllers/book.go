package controllers

import (
	"net/http"

	"github.com/alcjohn/rest_gin/middlewares"
	"github.com/alcjohn/rest_gin/utils"

	"github.com/alcjohn/rest_gin/dto"
	"github.com/alcjohn/rest_gin/models"
	"github.com/gin-gonic/gin"
)

type BooksController struct{}

func BooksRoutes(r *gin.RouterGroup) {
	var controller BooksController

	r.GET("/", middlewares.PaginationMiddleware(), controller.Index)
	r.POST("/", controller.Create)
	r.Use(middlewares.BookMiddlewares())
	{
		r.GET("/:book_id", controller.Show)
		r.PATCH("/:book_id", controller.Update)
		r.DELETE("/:book_id", controller.Delete)
		CommentsRoutes(r.Group("/:book_id/comments"))
	}
}

func (controller *BooksController) Index(c *gin.Context) {
	var books []models.Book

	pagination := c.Keys["Pagination"].(utils.Pagination)
	db := models.DB.Where("id > 0")
	c.JSON(http.StatusOK, pagination.Paginate(db, &books))
}

func (controller *BooksController) Show(c *gin.Context) {
	book := c.Keys["Book"].(models.Book)
	c.JSON(http.StatusOK, gin.H{"data": book})

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

	book := c.Keys["Book"].(models.Book)

	// Validate input
	var input dto.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func (controller *BooksController) Delete(c *gin.Context) {

	book := c.Keys["Book"].(models.Book)

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
