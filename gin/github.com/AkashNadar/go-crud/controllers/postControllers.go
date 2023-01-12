package controllers

import (
	"fmt"
	"log"

	"github.com/AkashNadar/go-crud/initializers"
	"github.com/AkashNadar/go-crud/model"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {

	data := model.Post{
		Title: "Hello",
		Body:  "Hello Body",
	}
	// fmt.Println(initializers.DB)
	res := initializers.DB.Create(&data)
	fmt.Println(res)
	if res.Error != nil {
		log.Fatalf("Insertion failed")
		c.Status(400)
	}

	c.JSON(200, gin.H{
		"post": data,
	})

}
