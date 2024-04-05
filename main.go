package main

import (
	"github.com/gin-gonic/gin"
	"github.com/satyam-jha-16/gin-api/controllers"
	"github.com/satyam-jha-16/gin-api/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/all", controllers.PostsIndex)
	r.GET("/all/:id", controllers.PostShow)
	r.PUT("/all/:id", controllers.PostUpdate)
	r.DELETE("/all/:id", controllers.PostDelete)
	r.Run()
	// fmt.Println("hello world")
}
