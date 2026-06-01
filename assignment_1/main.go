package main

import (
	"github.com/gin-gonic/gin"

	"insta-clone/handler"
	"insta-clone/middleware"
)

func main() {
	router := gin.Default()
	// defining a gin's router and using the middleware for logging the request and its performance in the console .
	router.Use(middleware.RequestLogger())

	//Defining the v1 engine group for the API and defining the routes for the API as per the PS . We are using the handler functions defined in the handler package to handle the requests and send the responses . We are also using the response package to send the responses in a standard format as defined in the PS .
	// POST - sending the request
	// GET - getting the data from the server and sending it to the client
	// we are using the path parameters for getting the specific user or post by its ID and we are using the request body for creating a new user or post or comment and liking a post .
	v1 := router.Group("/api/v1")
	{
		// 7 endpoints as in PS 
		v1.POST("/users", handler.CreateUser)
		v1.GET("/users/:id", handler.GetUserbyID)
		v1.POST("/posts", handler.CreatePost)
		v1.GET("/posts", handler.GetAllPosts)
		v1.GET("/posts/:id", handler.GetPostByID)
		v1.POST("/posts/:id/comments", handler.AddCommentToPost)
		v1.POST("/posts/:id/like", handler.LikePost)
	}

	router.Run(":8080")
	// in case we do not want to show the port we can use the .env file to set the port
	// but we have to add the config for loading it .
}
