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
	// 例如：测试 API 是否正常
	r.GET("/ping", PingHandler)
	// 上传文件
	r.POST("/upload",authController.JwtMiddleware, fileController.UpLoadFile)
	// 获取文件
	r.GET("/file/:filename",authController.JwtMiddleware, fileController.GetFile)
	// 例如：获取用户信息[校验中间件]
	r.GET("/user/:id", authController.JwtMiddleware, userController.GetUser)
	// 注册 API
	r.POST("/register", authController.RegisterHandler)
	// 登录 API
	r.POST("/login", authController.LoginHandler)

}

// PingHandler 用于测试 API 是否正常
func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
