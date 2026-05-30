package models


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

type Post struct {
	ID 	 int    `json:"id"`
	UserID int    `json:"user_id"`
	ImageURL string `json:"image_url"`
	Caption  string `json:"caption"`
	Timestamp string `json:"timestamp"`
	LikesCount int    `json:"likes_count"`
}

type CreatePostRequest struct {
	UserID   int    `json:"user_id" binding:"required"`
	ImageURL string `json:"image_url" binding:"required,url"`
	Caption  string `json:"caption"`
}