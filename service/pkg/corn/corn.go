package cron

import (
	"github.com/robfig/cron/v3"
)

// 全局cron实例（如果需要在外部控制停止，可导出）
var Cron *cron.Cron

// Init 初始化定时任务（创建实例+注册任务+启动）
// 返回cron实例，方便外部defer Stop()
func Init() (*cron.Cron, error) {
	// 1. 创建带秒级解析的实例
	Cron = cron.New(cron.WithSeconds())

	// 2. 注册所有任务
	//if err := Setup(Cron); err != nil {
	//	return nil, err
	//}

	// 3. 启动定时任务
	Cron.Start()

	return Cron, nil
}

// Setup 原注册函数（不变）
func Setup(c *cron.Cron) error {
	if err := registerTestTasks(c); err != nil {
		return err
	}
	return nil
}
