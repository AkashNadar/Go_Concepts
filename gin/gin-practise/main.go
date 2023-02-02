package main

import (
	"practise_gin/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/album", controller.GetAlbums)
	r.POST("/album", controller.CreateAlbum)
	r.GET("/album/:id", controller.GetAlbumById)
	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}
