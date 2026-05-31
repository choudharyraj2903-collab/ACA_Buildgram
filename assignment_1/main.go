package main

import (
	"github.com/gin-gonic/gin"

	"insta-clone/handler"
	"insta-clone/middleware"
)

func main(){
	router := gin.Default()

	router.Use(middleware.RequestLogger())


	v1:= router.Group("/api/v1")
	{
		v1.POST("/users", handler.CreateUser)
		v1.GET("/users/:id", handler.GetUserbyID)
		v1.POST("/posts", handler.CreatePost)
		v1.GET("/posts", handler.GetAllPosts)
		v1.POST("/posts/:id/comments", handler.AddCommentToPost)
		v1.POST("/posts/:id/like", handler.LikePost)
		v1.GET("posts/:id/comments", handler.GetAllComments)
	}

	router.Run(":8080")
}
