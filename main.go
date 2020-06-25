package main

import (
	"log"
	"os"

	"github.com/alcjohn/rest_gin/controllers"
	"github.com/alcjohn/rest_gin/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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
	}

	books := r.Group("/api/books")
	{
		books.GET("/", controllers.FindBooks)
		books.POST("/", controllers.CreateBook)
		books.GET("/:id", controllers.FindBook)
		books.PATCH("/:id", controllers.UpdateBook)
		books.DELETE("/:id", controllers.DeleteBook)
	}

	r.Run(":" + port)
}
