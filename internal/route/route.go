package route

import (
	"icu/internal/controller"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 配置所有的路由
func SetupRoutes(r *gin.Engine) {
	userController := controller.NewUserController()
	// 例如：测试 API 是否正常
	r.GET("/ping", PingHandler)
	// 例如：获取用户信息
	r.GET("/user/:id", userController.GetUser)
}

// PingHandler 用于测试 API 是否正常
func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
