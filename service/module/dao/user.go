package dao

import (
	"SService/module/model"
	"SService/pkg/database"
	"fmt"
)

// 根据用户名查询用户
func FindUserByName(name string) (*model.User, error) {
	var user model.User
	// Where 条件查询 + First 获取第一条记录
	result := database.DB.Where("username = ?", name).First(&user)
	return &user, result.Error
}

// 创建用户
func CreateUser(user *model.User) error {
	// Create 插入记录
	fmt.Printf("UUID类型: %T, 长度: %d字节\n", user.UUID, len(user.UUID))
	result := database.DB.Create(user)
	return result.Error
}
