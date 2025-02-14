package controller

import (
	"net/http"

	"icu/internal/service"

	"github.com/gin-gonic/gin"
)

// UserService 用于处理与用户相关的业务逻辑
type UserController struct {
	service *service.UserService
}

func NewUserController() *UserController{
	return &UserController{
		service: service.NewUserService(),
	}
}

// GetUserHandler 获取用户信息的处理函数
func (a *UserController) GetUser(c *gin.Context) {
	id := c.Param("userId")
	userService := service.NewUserService()
	user, err := userService.GetUser(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "User not found",
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

//传递用户头像文件名称，更新用户头像
func (a *UserController) UpdateAvatar(c *gin.Context) {
	//更新用户头像
	imgKey,err := a.service.UpdateAvatar(c);
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": map[string]string{"imgKey": imgKey}})
}
	

