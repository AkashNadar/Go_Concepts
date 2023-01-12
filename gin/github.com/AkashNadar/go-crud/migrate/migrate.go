package main

import (
	"github.com/AkashNadar/go-crud/initializers"
	"github.com/AkashNadar/go-crud/model"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&model.Post{})
}
