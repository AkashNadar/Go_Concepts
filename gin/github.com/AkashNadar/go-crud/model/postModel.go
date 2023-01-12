package model

import (
	"github.com/AkashNadar/go-crud/initializers"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string
	Body  string
}

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}
