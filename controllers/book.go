package controllers

import (
	"net/http"
	"strconv"

	"github.com/alcjohn/rest_gin/dto"
	"github.com/alcjohn/rest_gin/models"
	"github.com/alcjohn/rest_gin/utils"
	"github.com/gin-gonic/gin"
)

type BooksController struct{}

func BooksRoutes(r *gin.RouterGroup) {
	var controller BooksController
	r.GET("/", controller.Index)
	r.GET("/:id", controller.Show)
	r.POST("/", controller.Create)
	r.PATCH("/:id", controller.Update)
	r.DELETE("/:id", controller.Delete)
}

func (controller *BooksController) Index(c *gin.Context) {
	var books []models.Book
	db := models.DB.Where("id > 0")
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 30
	}
	pagination := &utils.Pagination{
		Page:    page,
		Limit:   limit,
		OrderBy: c.QueryArray("sort[]"),
	}
	c.JSON(http.StatusOK, pagination.Paginate(db, &books))
}

func (controller *BooksController) Show(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

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

func (controller *BooksController) Delete(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
