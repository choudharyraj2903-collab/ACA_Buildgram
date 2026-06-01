package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"insta-clone/models"
	"insta-clone/response"
)

//defining the user slice 
var users = []models.User{}

// for generating the next id for each user
var UserIDCounter int = 0

func nextUserID() int {
	UserIDCounter++
	return UserIDCounter
}

func CreateUser(c *gin.Context) {
	// models struct for giving the standard format for the incoming request data and for validating the request data as per the PS . We are using the ShouldBindJSON method of gin to bind the incoming request data to the struct and validate it at the same time . We are also using the response package to send the responses in a standard format as defined in the PS .
	var req models.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, "Invalid request data")
		return
	}

	// trimming 
	req.Username = strings.TrimSpace(req.Username)
	req.Email = strings.TrimSpace(req.Email)
	req.Bio = strings.TrimSpace(req.Bio)

	// checking the required fields are not empty after trimming and sending error response if any of the required field is empty as per the PS we have to send error response if any of the required field is empty in the request body and we are also using the response package to send the responses in a standard format as defined in the PS .

	if req.Username == "" || req.Email == "" {
		response.SendErrorResponse(c, http.StatusBadRequest, "username and email are required")
		return
	}

	//creating the user and adding the slice 
	user := models.User{
		ID:       nextUserID(),
		Username: req.Username,
		Email:    req.Email,
		Bio:      req.Bio,
	}
	users = append(users, user)
	response.SendSuccessResponse(c, http.StatusCreated, user)
}

// getting the user by its ID fro path parameter

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
