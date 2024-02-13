package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mohammedfuta2000/jwt-project/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/api-1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for API 1"})
	})

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	

	router.GET("/api-2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for API 2"})
	})

	router.Run(":" + port)
}
