package controller

import (
	"icu/internal/service"

	"github.com/gin-gonic/gin"
)

// UserService 用于处理与用户相关的业务逻辑
type UserController struct {
	service *service.UserService
	// 获取用户定义
	GetUser func(c *gin.Context) 
}

func NewUserController() *UserController{
	return &UserController{
		service: service.NewUserService(),
	}
}

// GetUserHandler 获取用户信息的处理函数
func GetUser(c *gin.Context) {
	id := c.Param("id")
	userService := service.NewUserService()
	user, err := userService.GetUser(id)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "User not found",
		})
		return
	}
	c.JSON(200, user)
}
