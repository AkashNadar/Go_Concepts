package main

import (
	"github.com/AkashNadar/go-crud/controllers"
	"github.com/AkashNadar/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/post", controllers.CreatePost)
	r.Run()
}
