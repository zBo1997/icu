package route

import (
	"icu/internal/controller"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 配置所有的路由
func SetupRoutes(r *gin.Engine) {
	userController := controller.NewUserController()
	authController := controller.NewAuthController()
	fileController := controller.NewFileController()
	chatController := controller.NewChatController()
	// 例如：测试 API 是否正常
	r.GET("/api/ping", PingHandler)
	// sse长连接
	r.GET("/api/sse/:conversationId", chatController.ChatAI)
	// sse长连接
	r.POST("/api/sendMessage/", chatController.SendMessage)
	// 上传文件
	r.POST("/api/upload",authController.JwtMiddleware, fileController.UpLoadFile)
	// 获取文件
	r.GET("/api/file/:filename", fileController.GetFile)
	// 例如：获取用户信息[校验中间件]
	r.GET("/api/user/:id", authController.JwtMiddleware, userController.GetUser)
	// 例如：获取用户信息[校验中间件]
	r.POST("/api/updateAvatar", authController.JwtMiddleware, userController.UpdateAvatar)
	// 注册 API
	r.POST("/api/register", authController.RegisterHandler)
	// 登录 API
	r.POST("/api/login", authController.LoginHandler)

}

// PingHandler 用于测试 API 是否正常
func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
