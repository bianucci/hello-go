package main

import (
	"github.com/bianucci/hello-go/controllers"
	"github.com/bianucci/hello-go/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	public := router.Group("/api")
	public.POST("/login", controllers.Login)
	public.POST("/register", controllers.Register)

	protected := router.Group("/api/albums")
	protected.Use(middlewares.JwtAuthMiddleware())

	protected.GET("", controllers.GetAlbums)
	protected.POST("", controllers.PostAlbum)
	protected.GET("/:id", controllers.GetAlbum)

	router.Run("localhost:8080")
}
