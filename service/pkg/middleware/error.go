// pkg/middleware/error.go
package middleware

import (
	"SService/pkg/util"
	"SService/pkg/util/response"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// ErrorHandler 全局异常处理中间件
// 支持捕获自定义AppError和其他未知错误，区分处理并统一响应格式
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 1. 区分错误类型：自定义AppError还是未知错误
				switch e := err.(type) {
				case *util.AppError:
					// 2. 处理自定义业务错误
					// 打印原始错误（内部调试用）
					if e.Err != nil {
						fmt.Printf("业务错误: %v (原始错误: %v)\n", e.Message, e.Err)
					} else {
						fmt.Printf("业务错误: %v\n", e.Message)
					}
					// 向前端返回自定义错误信息（使用指定的Code和Message）
					response.Result(e.Code, nil, e.Message, c)

				default:
					// 3. 处理未知错误（如空指针、数组越界等）
					// 打印完整堆栈信息，方便排查
					fmt.Printf("未知错误: %v\n堆栈信息: %s\n", err, debug.Stack())
					// 向前端返回通用错误信息（避免暴露敏感信息）
					response.Result(http.StatusInternalServerError, nil, "服务器内部错误", c)
				}

				// 终止后续请求处理
				c.Abort()
			}
		}()

		// 继续执行后续中间件和处理器
		c.Next()
	}
}
