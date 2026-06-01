package models
// the first step is to define the format in which we are storing , requesting and getting the data so that we have a welll defined format
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}
// User,Post and Comment are the main struct storing the data temporary we also have some 
// CreateUserRequest,CreatePostRequest and CreateCommentRequest struct which are used to bind the incoming request data to the struct and validate it using the binding tags
// PostWithComments struct is used to return the post along with its comments in a single response and LikePostResponse struct is used to return the updated likes count after liking a post
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
