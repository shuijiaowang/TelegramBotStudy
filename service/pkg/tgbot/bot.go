package tgbot

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// BotInstance 机器人实例包装器
type BotInstance struct {
	ID     int64              // 机器人ID
	Token  string             // 机器人Token
	Client *bot.Bot           // 底层TgBot客户端
	Ctx    context.Context    // 上下文（用于停止）
	Cancel context.CancelFunc // 取消函数
}

// BotManager 机器人管理器
type BotManager struct {
	bots sync.Map   // 存储机器人实例: key=botID(int64), value=*BotInstance
	mu   sync.Mutex // 串行执行，互斥锁，用于保证"检查机器人是否存在→创建→存储"等复合操作的原子性，防止并发冲突
}

// NewBotManager 创建机器人管理器
func NewBotManager() *BotManager {
	return &BotManager{}
}

// 启动机器人
func (m *BotManager) StartBot(botID int64, token string) (*BotInstance, error) {
	// 防止多个协程同时操作同一机器人或不同机器人时出现逻辑混乱（如重复创建、实例覆盖）
	m.mu.Lock()
	defer m.mu.Unlock()

	// 检查是否已启动
	if existing, ok := m.bots.Load(botID); ok {
		return existing.(*BotInstance), nil
	}
	// 创建上下文和取消函数：用于后续优雅停止机器人（调用Cancel可终止机器人运行）
	ctx, cancel := context.WithCancel(context.Background())

	// 初始化TgBot客户端
	opts := []bot.Option{
		bot.WithDefaultHandler(m.defaultMessageHandler(botID)), //绑定消息处理逻辑
		bot.WithSkipGetMe(), //跳过初始化时的网络请求：减少初始化阶段的网络IO阻塞，加快启动流程（后续可异步验证token）
	}
	b, err := bot.New(token, opts...)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("创建机器人失败: %v", err)
	}

	// 存储实例
	botInstance := &BotInstance{
		ID:     botID,
		Token:  token,
		Client: b,
		Ctx:    ctx,
		Cancel: cancel,
	}
	m.bots.Store(botID, botInstance)

	// 异步启动机器人，防阻塞
	go func() {
		m.registerCommands(b, ctx) //注册命令
		log.Printf("机器人 %d 启动成功", botID)
		b.Start(ctx) // 阻塞直到ctx被取消
	}()

	return botInstance, nil
}

// StopBot 停止机器人
func (m *BotManager) StopBot(botID int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if instance, ok := m.bots.Load(botID); ok {
		botInst := instance.(*BotInstance)
		botInst.Cancel() // 触发上下文取消，停止bot
		m.bots.Delete(botID)
		log.Printf("机器人 %d 已停止", botID)
		return nil
	}
	return fmt.Errorf("机器人 %d 不存在", botID)
}

// 获取机器人实例
func (m *BotManager) GetBot(botID int64) (*BotInstance, bool) {
	inst, ok := m.bots.Load(botID)
	if !ok {
		return nil, false
	}
	return inst.(*BotInstance), true
}

// 注册命令列表到Telegram服务器
func (m *BotManager) registerCommands(b *bot.Bot, parentCtx context.Context) error {
	// 给注册命令设置5秒超时（避免网络问题导致无限等待）
	ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
	defer cancel()
	// 注册命令
	_, err := b.SetMyCommands(ctx, &bot.SetMyCommandsParams{
		Commands: []models.BotCommand{
			{
				Command:     "/help",
				Description: "帮助",
			},
		},
	})

	if err != nil {
		log.Printf("机器人 %d 命令注册失败：%v", b.ID(), err)
		return err
	} else {
		log.Printf("机器人 %d 命令注册成功", b.ID())
		return nil
	}

}
