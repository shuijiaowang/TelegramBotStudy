package util

import (
	"log"
	"time"
)

// 测试定时任务：每10秒打印一次当前时间
func TestCronTask() {
	log.Printf("定时任务执行：当前时间 %s\n", time.Now().Format("2006-01-02 15:04:05"))
}
