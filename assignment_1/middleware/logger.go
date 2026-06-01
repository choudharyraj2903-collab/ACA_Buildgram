package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// This file is basically giving the stats of the requets like the method , path and latency of the request . We are using it as a middleware in the main.go file so that it gets executed for every request and gives us the stats of the request in the console . It is a good practice to have a logger middleware in our application to keep track of the requests and their performance .
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next() 

		latency := time.Since(start)
		fmt.Printf("[BuildGram] %s %s | %v\n", c.Request.Method, c.Request.URL.Path, latency)
		// Method , Path and Latency of the request is printed in the console in a specific format for better readability and understanding of the logs .
	}
}
