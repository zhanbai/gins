// Package middlewares Gin 中间件
package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Cors 中间件，解决跨域问题
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-type, Accept")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, PUT, PATCH, POST, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
