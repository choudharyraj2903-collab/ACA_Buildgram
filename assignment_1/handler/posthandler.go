package handler 

import (
	"strings"
	"time"
	"net/http"
	"strconv"


	"github.com/gin-gonic/gin"
	"insta-clone/models"
	"insta-clone/response"
)

var posts = []models.Post{}
var PostIDCounter int = 0

func nextPostID() int {
	PostIDCounter++
	return PostIDCounter
}

func CreatePost( c *gin.Context){
	var req models.CreatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, "Invalid request data")
		return
	}

	req.ImageURL = strings.TrimSpace(req.ImageURL)
	req.Caption = strings.TrimSpace(req.Caption)

	if req.UserID == 0 || req.ImageURL == "" {
		response.SendErrorResponse(c, http.StatusBadRequest, "user id and image url are required")
		return
	}

	post := models.Post{
		ID:        nextPostID(),
		UserID:    req.UserID,
		ImageURL:  req.ImageURL,
		Caption:   req.Caption,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	posts = append(posts, post)
	response.SendSuccessResponse(c, http.StatusCreated, post)
}

func GetAllPosts(c *gin.Context) {
	response.SendSuccessResponse(c, http.StatusOK, posts)
}

func LikePost(c *gin.Context) {
	idParam := c.Param("id")
	postID, err := strconv.Atoi(idParam)
	if err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	for i, post := range posts {
		if post.ID == postID {
			posts[i].LikesCount++
			response.SendSuccessResponse(c, http.StatusOK, posts[i])
			return
		}
	}
	response.SendErrorResponse(c, http.StatusNotFound, "Post not found")
}