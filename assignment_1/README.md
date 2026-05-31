# BuildGram REST API

BuildGram is a small REST API built with Go and Gin for a simplified Instagram-like app. It supports creating users, creating posts, liking posts, adding comments, and fetching a post with its comments.

This assignment uses in-memory storage with Go slices. That means all users, posts, and comments are reset when the server restarts. That is expected for this project.

## How to Run

Prerequisites:

- Go installed on your system
- A terminal opened inside the `assignment_1` folder

Install dependencies:

```bash
go mod tidy
```

Start the server:

```bash
go run .
```

The API runs on:

```text
http://localhost:8080
```

## Project Structure

```text
assignment_1/
  main.go                 Starts the Gin server and registers routes
  models/                 Request and response structs
  handler/                User, post, and comment route handlers
  middleware/             Custom request logger middleware
  response/               Helper functions for standard JSON responses
```

## API Reference

All routes are under:

```text
/api/v1
```

### Create User

```http
POST /api/v1/users
```

Required body:

```json
{
  "username": "harshit_is_sleeping",
  "email": "harshit@example.com",
  "bio": "Bio of Harshit"
}
```

### Get User by ID

```http
GET /api/v1/users/:id
```

### Create Post

```http
POST /api/v1/posts
```

Required body:

```json
{
  "userID": 1,
  "imageURL": "https://example.com/image.jpg",
  "caption": "My first BuildGram post"
}
```

### Get All Posts

```http
GET /api/v1/posts
```

### Get Post with Comments

```http
GET /api/v1/posts/:id
```

This returns the selected post and all comments linked to that post.

### Like a Post

```http
POST /api/v1/posts/:id/like
```

### Add Comment to Post

```http
POST /api/v1/posts/:id/comments
```

Required body:

```json
{
  "userID": 2,
  "text": "This is stunning!"
}
```

## Response Format

Successful responses use:

```json
{
  "status": "success",
  "data": {}
}
```

Error responses use:

```json
{
  "status": "error",
  "message": "human readable error message"
}
```
