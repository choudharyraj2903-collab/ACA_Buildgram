package main

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"insta-clone/handler"
	"insta-clone/models"
)

func main() {
	router := gin.Default()

	router.POST("/users", func(c *gin.Context) {
		var req models.CreateUserRequest
		
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		user, err := handler.CreateUser(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	router.POST("/posts",func(c *gin.Context) {
		var req models.CreatePostRequest
		
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		post, err := handler.CreatePost(req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, post)
	})

	router.Run(":8080")

}

