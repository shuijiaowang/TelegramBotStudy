package middleware

import (
	"SService/pkg/util"
	"SService/pkg/util/response"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// JWTInterceptor 验证JWT令牌的中间件
func JWTInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取Authorization
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.FailWithMessage("未提供token", c)
			c.Abort() //终止后续处理
			return    //返回
		}

		// 检查格式是否为Bearer <token>
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.FailWithMessage("token格式错误", c)
			c.Abort()
			return
		}

		// 解析token
		claims, err := util.ParseToken(parts[1])
		if err != nil {
			response.FailWithMessage("无效的token", c)
			c.Abort()
			return
		}
		// 【新增】判断 Token 是否即将过期（例如：剩余时间 < 30分钟），如果是则生成新 Token
		if claims.ExpiresAt != nil && claims.ExpiresAt.Unix()-time.Now().Unix() < 30*60 {
			newToken, err := util.GenerateToken(claims.UserID, claims.Username, claims.UserUUID)
			if err == nil {
				c.Header("new-token", newToken)
			}
		}
		// 将用户ID存入上下文
		//c.Set("userID", claims.UserID)
		//c.Set("username", claims.Username)
		//c.Set("userUUID", claims.UserUUID)
		c.Set("claims", claims)
		c.Next()
	}
}
