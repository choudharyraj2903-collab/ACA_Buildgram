package response

import "github.com/gin-gonic/gin"

func SendSuccessResponse(c *gin.Context, statuscode int, data any) {
	c.JSON(statuscode, gin.H{
		"status": "success",
		"data":   data,
	})
}

func SendErrorResponse(c *gin.Context, statuscode int, message string) {
	c.JSON(statuscode, gin.H{
		"status":  "error",
		"message": message,
	})
}
