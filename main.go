package main

import (
	"log"
	"os"

	"github.com/alcjohn/rest_gin/controllers"
	"github.com/alcjohn/rest_gin/middlewares"
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

	r.Use(middlewares.AuthMiddleware())

	models.ConnectDatabase()

	controllers.AuthRoutes(r.Group("/api/auth"))
	controllers.BooksRoutes(r.Group("/api/books"))
	controllers.UsersRoutes(r.Group("api/users"))

	r.Run(":" + port)
}
