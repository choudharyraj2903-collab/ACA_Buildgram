package response

import "github.com/gin-gonic/gin"

//Bsic purpose of this file and functions is to give a standard format input and output as defined in the 
// PS .We are using a folder for it for orgainizing the things .


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
