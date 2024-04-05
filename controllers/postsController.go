package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/satyam-jha-16/gin-api/initializers"
	"github.com/satyam-jha-16/gin-api/models"
)

func PostsCreate(c *gin.Context) {
	//get data from the input
	var body struct {
		Name string
		Body string
	}
	c.Bind(&body)
	// create a post
	post := models.Post{Name: body.Name, Body: body.Body}
	result := initializers.DB.Create(&post)
	//return it

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	// get id of URL
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	//getting parameters from request
	id := c.Param("id")
	// getting data of request body
	var body struct {
		Name string
		Body string
	}
	c.Bind(&body)
	//finding the data to be changed
	var post models.Post
	initializers.DB.First(&post, id)

	//changing -- actual code
	initializers.DB.Model(&post).Updates(models.Post{
		Name: body.Name,
		Body: body.Body,
	})

	// respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	//getting id from param
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)
}
