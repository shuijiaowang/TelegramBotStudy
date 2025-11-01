package service

import (
	"SService/module/dao"
	"SService/module/model"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

// Register 处理用户注册逻辑
func (s *UserService) Register(username, password string) error {
	// 1. 检查用户名是否已存在
	existingUser, err := dao.FindUserByName(username)
	if err == nil && existingUser != nil {
		return errors.New("用户名已存在")
	}

	// 2. 生成UUID（用户唯一标识）
	userUUID := uuid.New()

	// 3. 密码加密（重要：禁止明文存储密码）
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("密码加密失败")
	}

	// 4. 构造用户模型
	user := &model.User{
		Username: username,
		Password: string(hashedPassword), // 存储加密后的密码
		UUID:     userUUID.String(),
	}

	// 5. 调用仓库层保存用户
	if err := dao.CreateUser(user); err != nil {
		return errors.New("注册失败，请重试")
	}

	return nil
}

func (s *UserService) Login(name, password string) (*model.User, bool) {
	user, err := dao.FindUserByName(name)
	if err != nil {
		return nil, false
	}

	// 使用bcrypt验证密码（对比明文密码和加密后的密码）
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// 密码不匹配
		return nil, false
	}
	return user, true
}
