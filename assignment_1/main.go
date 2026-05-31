package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"insta-clone/handler"
	"insta-clone/models"
)


func main() {
	router := gin.Default()

	router.POST("/users", handler.CreateUser)

	router.GET("/users/:id", handler.GetUserbyID)


	router.POST("/posts", func(c *gin.Context) {
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


	router.GET("/posts/:id", func(c *gin.Context) {

		idParam := c.Param("id")
		idParamInt, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
			return
		}

		posts ,err:= handler.GetPostByUserID(idParamInt)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, posts)
	})

	router.GET("/posts", func(c *gin.Context) {
		posts := handler.GetAllPosts()
		c.JSON(http.StatusOK, posts)
	})

	router.Run(":8080")
}
