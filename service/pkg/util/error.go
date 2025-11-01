package util

// AppError 自定义应用错误，用于封装业务错误信息
type AppError struct {
	Code    int    // HTTP状态码
	Message string // 错误描述（对外展示）//返回给前端
	Err     error  // 原始错误（内部调试用） //控制台日志
}

// Error 实现error接口，便于作为错误类型传递
func (e *AppError) Error() string {
	return e.Message
}

// 可选：添加构造函数，简化创建AppError的过程
func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
