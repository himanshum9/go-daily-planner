package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/himanshu/daily-planner/internal/auth"
)

func SessionAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip auth for login and register routes
		if isPublicRoute(c.Request.URL.Path) {
			c.Next()
			return
		}

		// Get token from cookie
		token, err := c.Cookie("auth_token")
		if err != nil {
			c.Redirect(http.StatusFound, "/auth/login")
			c.Abort()
			return
		}

		// Validate token and get user ID
		userID, err := auth.ValidateJWTToken(token)
		if err != nil {
			c.Redirect(http.StatusFound, "/auth/login")
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func isPublicRoute(path string) bool {
	publicRoutes := []string{
		"/auth/login",
		"/auth/register",
		"/auth/forgot-password",
		"/auth/reset-password",
		"/auth/google/login",
		"/auth/google/callback",
		"/static/",
	}

	for _, route := range publicRoutes {
		if len(route) <= len(path) && path[:len(route)] == route {
			return true
		}
	}
	return false
}
