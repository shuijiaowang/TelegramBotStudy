package tgbot

// Message 统一消息结构
type Message struct {
	FromUserID int64  // 发送者ID
	ToBotID    int64  // 目标机器人ID
	Content    string // 消息内容
	Type       string // 消息类型：text/image/file
	RawData    any    // 原始数据（Tg原始消息）
}
