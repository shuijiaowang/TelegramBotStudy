package routes

import (
	"SService/module/api"
	"SService/pkg/middleware" // 导入中间件包

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 跨域中间件（放在最前面）
	r.Use(middleware.Cors())
	// 注册全局异常处理中间件
	r.Use(middleware.ErrorHandler())

	userApi := api.UserApi{}
	// 用户路由
	userGroup := r.Group("/api/user")
	{
		userGroup.POST("/login", userApi.Login)
		userGroup.POST("/register", userApi.Register)
		userGroup.GET("/test", userApi.Test)
	}

	exampleApi := api.ExampleApi{}
	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.JWTInterceptor()) // 应用JWT拦截器
	{
		// 消费记录路由（需要认证）
		exampleGroup := apiGroup.Group("/example")
		{
			exampleGroup.POST("/test", exampleApi.Test) // 添加消费记录
		}
		// 消费拓展路由（需要认证）

	}

	return r
}
