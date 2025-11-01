package tgbot

import (
	"log/slog"
	"time"
)

func InitBot() {
	slog.Info("初始化机器人")
	// 1. 创建机器人管理器
	botManager := NewBotManager()
	//storage := NewSimpleStorage()

	// 2. 启动机器人（替换为你的BotID和Token）
	botID := int64(8295682465)                                // 机器人ID（自定义）
	token := "8295682465:AAE2g8IWgK1v2xPLeLffR3hUV5aa_BKezUE" // TgBot的Token
	_, err := botManager.StartBot(botID, token)
	if err != nil {
		panic(err)
	}

	// 3. 示例：主动发送消息
	go func() {
		// 创建一个5秒间隔的定时器
		ticker := time.NewTicker(60 * time.Second)
		defer ticker.Stop() // 退出时停止定时器，避免资源泄漏

		// 循环等待定时器触发
		for range ticker.C {
			// 每次触发时发送消息
			err := botManager.ReplyToUser(botID, 7600284259, "Hello! 我是机器人") // 7600284259是用户ID
			if err != nil {
				println("发送消息失败:", err.Error())
			}
		}
	}()

	// 4. 保持程序运行
	select {}
}
