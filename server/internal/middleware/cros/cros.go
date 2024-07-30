package cros

import (
	"github.com/gin-gonic/gin"
)

// Cors 设置跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("origin")

		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers",
			"Content-Type, AccessToken, X-CSRF-Token, Authorization, Token, X-Token, X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET")
		c.Header("Access-Control-Expose-Headers",
			"Content-Length, Access-Control-Allow-Origin, "+
				"Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
		c.Header("Access-Control-Allow-Credentials", "true")

		c.Next()
	}
}
