package tgbot

import "sync"

// 简单存储会话信息（替代原项目DAO）
type SimpleStorage struct {
	chatSessions sync.Map // key=telegramID(int64), value=会话信息
}

func NewSimpleStorage() *SimpleStorage {
	return &SimpleStorage{}
}

// 存储用户会话
func (s *SimpleStorage) SaveSession(telegramID int64, data any) {
	s.chatSessions.Store(telegramID, data)
}

// 获取用户会话
func (s *SimpleStorage) GetSession(telegramID int64) (any, bool) {
	return s.chatSessions.Load(telegramID)
}
