package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"insta-clone/models"
	"insta-clone/response"
)

var users = []models.User{}

var UserIDCounter int = 0

func nextUserID() int {
	UserIDCounter++
	return UserIDCounter
}

func CreateUser(c *gin.Context) {
	var req models.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, "Invalid request data")
		return
	}

	req.Username = strings.TrimSpace(req.Username)
	req.Email = strings.TrimSpace(req.Email)
	req.Bio = strings.TrimSpace(req.Bio)

	if req.Username == "" || req.Email == "" {
		response.SendErrorResponse(c, http.StatusBadRequest, "username and email are required")
		return
	}

	user := models.User{
		ID:       nextUserID(),
		Username: req.Username,
		Email:    req.Email,
		Bio:      req.Bio,
	}
	users = append(users, user)
	response.SendSuccessResponse(c, http.StatusCreated, user)
}

func GetUserbyID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	for _, user := range users {
		if user.ID == userID {
			response.SendSuccessResponse(c, http.StatusOK, user)
			return
		}
	}
	response.SendErrorResponse(c, http.StatusNotFound, "User not found")
}
