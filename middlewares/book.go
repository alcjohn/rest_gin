package middlewares

import (
	"net/http"

	"github.com/alcjohn/rest_gin/models"
	"github.com/gin-gonic/gin"
)

func BookMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {

		var book models.Book
		if err := models.DB.Where("id = ?", c.Param("book_id")).First(&book).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
		c.Set("Book", book)
	}
}
