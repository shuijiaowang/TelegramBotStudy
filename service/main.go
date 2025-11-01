package main

import (
	cron "SService/pkg/corn"
	"SService/pkg/database"
	"SService/routes"
	"log"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	// 初始化数据库
	if err := database.InitDB(); err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}

	// 初始化定时任务
	c, err := cron.Init()
	if err != nil {
		log.Fatalf("定时任务初始化失败: %v", err)
	}
	defer c.Stop() // 程序退出时停止

	// 创建路由
	r := routes.SetupRouter()

	// 启动服务
	if err := r.Run(":7789"); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
