package api

import (
	"SService/pkg/util"
	"SService/pkg/util/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserApi struct{}

func (h *UserApi) Register(c *gin.Context) {
	// 定义注册请求参数结构体
	var req struct {
		Username string `json:"username" binding:"required,min=1,max=20"` // 用户名长度限制
		Password string `json:"password" binding:"required,min=1"`        // 密码长度限制
	}

	// 绑定并验证请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("无效的请求格式：用户名至少3位，密码至少6位", c)
		return
	}

	// 调用服务层注册方法
	err := userService.Register(req.Username, req.Password)
	if err != nil {
		response.FailWithMessage("注册失败", c)
		return
	}
	// 注册成功
	response.Ok(c)
}
func (h *UserApi) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("无效的请求格式", c)
		return
	}

	user, ok := userService.Login(req.Username, req.Password)
	if !ok {
		response.FailWithMessage("用户名或密码错误", c)
		return
	}
	userUUID, err := uuid.Parse(user.UUID)
	if err != nil {
		response.FailWithMessage("UUID格式错误", c)
		return
	}
	// 生成JWT令牌
	token, err := util.GenerateToken(int(user.ID), user.Username, userUUID)
	if err != nil {
		response.FailWithMessage("生成令牌失败", c)
		return
	}

	response.OkWithData(gin.H{
		"id":       user.ID,
		"username": user.Username,
		"uuid":     user.UUID, // 可返回给前端用于展示或后续操作
		"token":    token,
	}, c)
}

func (h *UserApi) Test(c *gin.Context) {
	response.OkWithDetailed("data-ok", "msg-ok", c)
}
