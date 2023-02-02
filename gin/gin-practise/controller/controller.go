package controller

import (
	"log"
	"net/http"
	"practise_gin/model"

	"github.com/gin-gonic/gin"
)

var data []model.Album

func init() {
	data = []model.Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
}

func GetAlbums(c *gin.Context) {

	c.JSON(http.StatusOK, data)
}

func CreateAlbum(c *gin.Context) {
	var newAlbum model.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		log.Fatalf("Binding error: %v", err)
		return
	}
	data = append(data, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, val := range data {
		if val.ID == id {
			c.JSON(http.StatusOK, val)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "album not found :(",
	})
}
