// pkg/middleware/cors.go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cors 跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 允许的源（生产环境建议指定具体域名，如"http://localhost:5173"）
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin) // 动态允许请求源
		} else {
			c.Header("Access-Control-Allow-Origin", "*")
		}

		// 允许的请求头（包含前端可能传递的所有头）
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Origin, Accept")
		// 允许的请求方法
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		// 允许前端获取的头
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, X-Token")
		// 允许携带cookie（如果需要）
		c.Header("Access-Control-Allow-Credentials", "true")

		// 处理预检请求（OPTIONS）
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
