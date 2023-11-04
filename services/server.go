package services

import (
	"github.com/akufarish/gin-crud/controllers"
	"github.com/gin-gonic/gin"
)

func Init() {
	gin := gin.Default()

	v1 := gin.Group("/api/v1")

	{
		v1.POST("/auth/register", controllers.Register)
		v1.POST("/auth/login", controllers.Login)
	}

	gin.Run()
}