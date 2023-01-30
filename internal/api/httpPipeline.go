package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func AddHttpMiddleware(api *gin.Engine) {
	api.Use(tokenAuthorization())
	api.Use(customLogMiddleware())
}

func errorResponse(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func tokenAuthorization() gin.HandlerFunc {
	requiredToken := os.Getenv("API_TOKEN")

	if requiredToken == "" {
		log.Fatal("please set API_TOKEN environment variable")
	}

	return func(c *gin.Context) {
		token := c.Request.Header.Get("api_token")

		if token == "" {
			errorResponse(c, 401, "API token required")
			return
		}

		if token != requiredToken {
			errorResponse(c, 401, "invalid API token")
			return
		}

		c.Next()
	}
}

func customLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, exists := c.Get("resp")
		if exists {
			log.Printf("custom log info %v", resp)
		}

		c.Next()
	}
}
