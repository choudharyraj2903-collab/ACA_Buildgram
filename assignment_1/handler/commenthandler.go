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
var CommentIDCounter int = 0

func nextCommentID() int {
	CommentIDCounter++
	return CommentIDCounter
}

func AddCommentToPost(c *gin.Context) {

	idParam := c.Param("id")
	postID, err := strconv.Atoi(idParam)
	if err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	var req models.CreateCommentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, "Invalid request data")
		return
	}

	req.Text = strings.TrimSpace(req.Text)

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
	PostExist := false
	for _, post := range posts {
		if post.ID == postID {
			PostExist = true
			break
		}
	}
	if !UserExist {
		response.SendErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}
	if !PostExist {
		response.SendErrorResponse(c, http.StatusNotFound, "Post not found")
		return
	}

	comment := models.Comment{
		ID:        nextCommentID(),
		PostID:    postID,
		UserID:    req.UserID,
		Text:      req.Text,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	comments = append(comments, comment)
	response.SendSuccessResponse(c, http.StatusCreated, comment)
}

func GetAllComments(c *gin.Context) {
	response.SendSuccessResponse(c, http.StatusOK, comments)
}
