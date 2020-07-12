package main

import (
	"log"
	"os"

	"github.com/alcjohn/rest_gin/controllers"
	"github.com/alcjohn/rest_gin/middlewares"
	"github.com/alcjohn/rest_gin/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	models.ConnectDatabase()
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}

	r.Use(cors.New(config))
	r.Use(middlewares.AuthMiddleware())

	controllers.AuthRoutes(r.Group("/api/auth"))
	controllers.BooksRoutes(r.Group("/api/books"))

	r.Run(":" + port)
}
