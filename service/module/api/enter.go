package api

import (
	"SService/module/service"
)

// HandlerGroup 包含所有处理器的结构
type ApiGroup struct {
	ExampleApi
	UserApi
}

var (
	exampleService = service.ExampleService{}
	userService    = service.UserService{}
)
