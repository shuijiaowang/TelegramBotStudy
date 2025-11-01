package service

import (
	"fmt"
)

// ExampleService 示例服务结构体
type ExampleService struct{}

// Hello 返回问候语
func (s *ExampleService) AddExample() string {
	return fmt.Sprintf("Hello")
}
