package main

import (
	"event-booking-api/internal/auth"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("register", auth.Register)
		authRoutes.POST("login", auth.Login)
	}
	//bjnk
	log.Println("Server running on port 8080")
	r.Run(":8080")
}
