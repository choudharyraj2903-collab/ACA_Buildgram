package main

import (
	"github.com/gin-gonic/gin"

	"insta-clone/handler"
)


func main() {
	router := gin.Default()

	router.POST("/users", handler.CreateUser)

	router.GET("/users/:id", handler.GetUserbyID)

	router.POST("/posts", handler.CreatePost)

	router.GET("/posts", handler.GetAllPosts)

	router.POST("/posts/:id/comments", handler.AddCommentToPost)

	router.POST("/posts/:id/like", handler.LikePost)

	router.GET("posts/:id/comments", handler.GetAllComments)

	router.Run(":8080")
}
