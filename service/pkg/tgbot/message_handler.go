package tgbot

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// 默认消息处理器（接收Tg消息）
func (m *BotManager) defaultMessageHandler(botID int64) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message != nil {
			m.handleMessage(botID, update)
			slog.Info("原始消息", update.Message)
		}
	}
}

// 处理消息（可自定义业务逻辑）
func (m *BotManager) handleMessage(botID int64, update *models.Update) {
	// 示例：自动回复文本消息
	if update.Message.Text != "" {
		m.ReplyToUser(botID, update.Message.From.ID, update.Message.Text)
	}
}

// 回复用户消息
func (m *BotManager) ReplyToUser(botID, userID int64, content string) error {
	botInst, ok := m.GetBot(botID)
	if !ok {
		return fmt.Errorf("机器人 %d 未启动", botID)
	}

	_, err := botInst.Client.SendMessage(botInst.Ctx, &bot.SendMessageParams{
		ChatID: userID,
		Text:   content,
	})
	return err
}
