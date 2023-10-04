package main

import (
	"Go_REST_API_wit_JWT_Authentication_in_Gin/controllers"
	"Go_REST_API_wit_JWT_Authentication_in_Gin/middlewares"
	"Go_REST_API_wit_JWT_Authentication_in_Gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()
	router := gin.Default()

	public := router.Group("/api")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := router.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)
	router.Run(":8080")
}
