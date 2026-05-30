package handler 

import (
	"strings"
	"errors"
	"time"

	"insta-clone/models"
)

var posts = []models.Post{}
var PostIDCounter int = 0

func nextPostID() int {
	PostIDCounter++
	return PostIDCounter
}

func CreatePost(req models.CreatePostRequest) (models.Post, error) {

	req.ImageURL = strings.TrimSpace(req.ImageURL)
	req.Caption = strings.TrimSpace(req.Caption)

	if req.ImageURL == "" || req.UserID == 0 {
		return models.Post{}, errors.New("image url and user id are required")
	}

	post := models.Post{
		ID:        nextPostID(),
		UserID:    req.UserID,
		ImageURL:  req.ImageURL,
		Caption:   req.Caption,
		Timestamp: time.Now().Format(time.RFC3339),
		LikesCount: 0,
	}
	posts = append(posts, post)
	return post, nil

}

func GetAllPosts(userID int) ([]models.Post ,error) {
	var userPosts []models.Post	
	for _, post := range posts {
		if post.UserID == userID {
			userPosts = append(userPosts, post)
		}	
	}
	return userPosts, nil
}