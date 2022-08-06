// Package middlewares Gin 中间件
package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Accept 中间件，设置默认响应格式
func Accept() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Accept", "application/json")

		c.Next()
	}
}
