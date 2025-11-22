package middleware

import (
	"github.com/gin-gonic/gin"
)

// RequestHandler is a middleware that sets session information
func RequestHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("session", "user1")
		context.Next()
	}
}