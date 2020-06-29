package main

import (
	"log"
	"net/http"
	"os"

	"github.com/alcjohn/rest_gin/auth"
	"github.com/alcjohn/rest_gin/controllers"
	"github.com/alcjohn/rest_gin/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "StatusUnauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	r := gin.Default()

	models.ConnectDatabase()

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
		auth.GET("/me", AuthMiddleware(), controllers.Me)
	}

	var booksModel []models.Book
	books := r.Group("/api/books")
	books.Use(AuthMiddleware())
	{
		books.PATCH("/:id", controllers.UpdateBook)
		books.POST("/", controllers.CreateBook)
		books.GET("/", controllers.Paginate(&booksModel))
		books.DELETE("/:id", controllers.Delete(&booksModel))
		books.GET("/:id", controllers.Show(&booksModel))
	}

	var usersModel []models.User
	users := r.Group("/api/users")
	{
		users.GET("/", controllers.Paginate(&usersModel))
	}

	r.Run(":" + port)
}
