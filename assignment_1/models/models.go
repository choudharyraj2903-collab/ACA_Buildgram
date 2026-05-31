package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Bio      string `json:"bio"`
}

type Post struct {
	ID         int    `json:"id"`
	UserID     int    `json:"userID"`
	ImageURL   string `json:"imageURL"`
	Caption    string `json:"caption"`
	Timestamp  string `json:"timestamp"`
	LikesCount int    `json:"likesCount"`
}

type CreatePostRequest struct {
	UserID   int    `json:"userID" binding:"required"`
	ImageURL string `json:"imageURL" binding:"required,url"`
	Caption  string `json:"caption" binding:"required"`
}

type Comment struct {
	ID        int    `json:"id"`
	PostID    int    `json:"postID"`
	UserID    int    `json:"userID"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
}

type CreateCommentRequest struct {
	UserID int    `json:"userID" binding:"required"`
	Text   string `json:"text" binding:"required"`
}

type PostWithComments struct {
	Post     Post      `json:"post"`
	Comments []Comment `json:"comments"`
}

type LikePostResponse struct {
	ID         int `json:"id"`
	LikesCount int `json:"likesCount"`
}
