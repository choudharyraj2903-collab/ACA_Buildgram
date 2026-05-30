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

	router.GET("/users", func(c *gin.Context) {
		users := handler.GetAllUsers()
		c.JSON(http.StatusOK, users)
	})

	router.GET("/users/:id", func(c *gin.Context) {

		idParam := c.Param("id")
		idParamInt, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
			return
		}

		user, err := handler.GetUserByID(idParamInt)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	router.Run(":8080")
}
