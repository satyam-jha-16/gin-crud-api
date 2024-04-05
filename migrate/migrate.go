package main

import (
	"github.com/satyam-jha-16/gin-api/initializers"
	"github.com/satyam-jha-16/gin-api/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
