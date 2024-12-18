package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sachin-gautam/gin-api/controller"
	"github.com/sachin-gautam/gin-api/database"
	"github.com/sachin-gautam/gin-api/middleware"
	"github.com/sachin-gautam/gin-api/model"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controller.AddEntry)
	protectedRoutes.GET("/entry", controller.GetAllEntries)
	protectedRoutes.GET("/entry/:id", controller.GetEntryByID)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})
}
