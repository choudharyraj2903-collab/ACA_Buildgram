package handler

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"insta-clone/models"
	"insta-clone/response"
)

var comments = []models.Comment{}

// THis function is responsible for genrating the unique comment ID for each comment  bcz as in database we do not have the concept of unique id's
var CommentIDCounter int = 0
func nextCommentID() int {
	CommentIDCounter++
	return CommentIDCounter
}
//resets after every run of the program but in real world we use database 

// AddCommentToPost function is responsible for adding a comment to a specific post . It takes the post ID from the path parameter and the comment data from the request body . It validates the request data and checks if the user and post exist before adding the comment to the comments slice and sending the response back to the client . We are also using the response package to send the responses in a standard format as defined in the PS .
func AddCommentToPost(c *gin.Context) {

	//getting ID from path parameter
	// if the id is not valid then error : invalid post ID and return
	idParam := c.Param("id")
	// converting ID to int bcz path parameters are always string and we need to convert it to int for further processing and validation
	postID, err := strconv.Atoi(idParam)
	if err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	//using the model struct to bind the incoming request data to the struct
	var req models.CreateCommentRequest

	//validating the request data
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, "Invalid request data")
		return
	}

	//trimming the request data
	req.Text = strings.TrimSpace(req.Text)
	// validating the request data after trimming
	if req.UserID == 0 || req.Text == "" {
		response.SendErrorResponse(c, http.StatusBadRequest, "user id and text are required")
		return
	}

	//verifying if user exist
	UserExist := false
	for _, user := range users {
		if user.ID == req.UserID {
			UserExist = true
			break
		}
	}
	//verifying if post exist
	PostExist := false
	for _, post := range posts {
		if post.ID == postID {
			PostExist = true
			break
		}
	}
	// sending error response if user or post does not exist in the system as we cannot add comment to a post which does not exist and we cannot add comment by a user who does not exist in the system .
	if !UserExist {
		response.SendErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}
	if !PostExist {
		response.SendErrorResponse(c, http.StatusNotFound, "Post not found")
		return
	}

	// adding the comment to the comments slice

	comment := models.Comment{
		ID:        nextCommentID(),
		PostID:    postID,
		UserID:    req.UserID,
		Text:      req.Text,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	comments = append(comments, comment)
	//201 if successful request and comment is added successfully to the comments slice and sending the response back to the client in a standard format as defined in the PS using the response package .
	response.SendSuccessResponse(c, http.StatusCreated, comment)
}

