package controllers

import (
	"net/http"

	"github.com/alcjohn/rest_gin/middlewares"

	"github.com/alcjohn/rest_gin/dto"
	"github.com/alcjohn/rest_gin/models"
	"github.com/alcjohn/rest_gin/utils"
	"github.com/gin-gonic/gin"
)

type CommentController struct{}

func CommentsRoutes(r *gin.RouterGroup) {
	var controller CommentController
	r.GET("/", middlewares.PaginationMiddleware(), controller.Index)
	r.GET("/:comment_id", controller.Show)
	r.POST("/", controller.Create)
	r.PATCH("/:comment_id", controller.Update)
	r.DELETE("/:comment_id", controller.Delete)
}

func (controller *CommentController) Index(c *gin.Context) {

	var comments []models.Comment
	book := c.Keys["Book"].(models.Book)
	pagination := c.Keys["Pagination"].(utils.Pagination)
	db := models.DB.Where("book_id = ?", book.ID)
	c.JSON(http.StatusOK, pagination.Paginate(db, &comments))
}

func (controller *CommentController) Show(c *gin.Context) {
	book := c.Keys["Book"].(models.Book)
	var comment models.Comment
	if err := models.DB.Preload("User").Where("id = ?", c.Param("comment_id")).Where("book_id = ?", book.ID).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

func (controller *CommentController) Create(c *gin.Context) {
	user := c.Keys["AuthUser"].(models.User)
	book := c.Keys["Book"].(models.Book)
	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	var input dto.CreateComment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	comment := models.Comment{Content: input.Content, UserID: user.ID, BookID: book.ID}
	models.DB.Create(&comment).Preload("User")

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

func (controller *CommentController) Update(c *gin.Context) {

}

func (controller *CommentController) Delete(c *gin.Context) {

}
