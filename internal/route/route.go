package route

import (
	"icu/internal/service"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 配置所有的路由
func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", PingHandler)

	// 例如：获取用户信息
	r.GET("/user/:id", GetUserHandler)
}

// PingHandler 用于测试 API 是否正常
func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// GetUserHandler 获取用户信息的处理函数
func GetUserHandler(c *gin.Context) {
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
