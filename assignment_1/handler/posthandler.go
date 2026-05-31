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

var posts = []models.Post{}
var PostIDCounter int = 0

func nextPostID() int {
	PostIDCounter++
	return PostIDCounter
}

func CreatePost(c *gin.Context) {
	var req models.CreatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, "Invalid request data")
		return
	}

	req.ImageURL = strings.TrimSpace(req.ImageURL)
	req.Caption = strings.TrimSpace(req.Caption)

	if req.UserID == 0 || req.ImageURL == "" || req.Caption == "" {
		response.SendErrorResponse(c, http.StatusBadRequest, "user id, image url and caption are required")
		return
	}

	UserExist := false
	for _, user := range users {
		if user.ID == req.UserID {
			UserExist = true
			break
		}
	}

	if !UserExist {
		response.SendErrorResponse(c, http.StatusNotFound, "User not found")
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

func GetPostByID(c *gin.Context) {
	idParam := c.Param("id")
	postID, err := strconv.Atoi(idParam)
	if err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, "Invalid post ID")
		return
	}

	var selectedPost models.Post
	postFound := false
	for _, post := range posts {
		if post.ID == postID {
			selectedPost = post
			postFound = true
			break
		}
	}

	if !postFound {
		response.SendErrorResponse(c, http.StatusNotFound, "Post not found")
		return
	}

	postComments := []models.Comment{}
	for _, comment := range comments {
		if comment.PostID == postID {
			postComments = append(postComments, comment)
		}
	}

	response.SendSuccessResponse(c, http.StatusOK, models.PostWithComments{
		Post:     selectedPost,
		Comments: postComments,
	})
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
			response.SendSuccessResponse(c, http.StatusOK, models.LikePostResponse{
				ID:         posts[i].ID,
				LikesCount: posts[i].LikesCount,
			})
			return
		}
	}
	response.SendErrorResponse(c, http.StatusNotFound, "Post not found")
}
