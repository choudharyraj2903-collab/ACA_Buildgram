package handler 

import (
	"strings"
	"errors"

	"insta-clone/models"
)

var users []models.User

var UserIDCounter int = 0

func nextUserID() int {
	UserIDCounter++
	return UserIDCounter
}

func CreateUser(req models.CreateUserRequest) (models.User, error) {

	req.Username = strings.TrimSpace(req.Username)
	req.Email = strings.TrimSpace(req.Email)

	if req.Username == "" || req.Email == "" {
		return models.User{},  errors.New("username and email are required")
	}

	user := models.User{
		ID:        nextUserID(),
		Username:  strings.ToLower(req.Username),
		Email:     strings.ToLower(req.Email),
		Bio:	   req.Bio,
	}
	users = append(users, user)
	return user, nil

}