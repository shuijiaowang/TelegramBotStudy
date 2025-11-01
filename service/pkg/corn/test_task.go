package cron

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

// 注册测试相关的定时任务
func registerTestTasks(c *cron.Cron) error {
	// 测试任务1：每10秒打印时间
	_, err := c.AddFunc("1 * * * * *", testPrintTime)
	if err != nil {
		return err
	}

	// 测试任务2：每30秒打印日志（示例）
	_, err = c.AddFunc("*/5 */2 * * * *", testLog)
	if err != nil {
		return err
	}

	return nil
}

// 测试任务1：打印当前时间
func testPrintTime() {
	log.Printf("测试任务执行：当前时间 %s\n", time.Now().Format("2006-01-02 15:04:05"))
}

// 测试任务2：打印固定日志
func testLog() {
	log.Println("测试任务2执行：这是30秒一次的日志")
}
