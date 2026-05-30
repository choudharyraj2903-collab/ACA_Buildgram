package models

import "time"


type User struct {
	ID        int       `json:"id"`
	Username string    `json:"username"`
	Email     string    `json:"email"`
	Bio	   string    `json:"bio"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Bio	   string `json:"bio"`
}