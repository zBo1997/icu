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
	articleController := controller.NewArticleController()
	commentController := controller.NewCommentController()
	captchaController := controller.NewCaptchaController()
	tagController := controller.NewTagController()
	// 例如：测试 API 是否正常
	r.GET("/api/ping", PingHandler)
	// sse长连接
	r.GET("/api/sse", chatController.ChatAI)
	
	
	// 上传文件
	r.POST("/api/upload",authController.JwtMiddleware, fileController.UpLoadFile)
	// 获取文件
	r.GET("/api/file/:filename", fileController.GetFile)

	// 文章分页列表
	r.GET("/api/pageArticle", articleController.PageArticle)
	// 文章详情
	r.GET("/api/article/:articleId", articleController.GetArticle)
	//发布文章
	r.POST("/api/article/pulish", authController.JwtMiddleware, articleController.PublishArticle)

	// 根据文章获取评论
	r.GET("/api/comments/:articleId", commentController.GetCommentsHandler)
	// 添加评论
	r.POST("/api/comments/add/:articleId", authController.JwtMiddleware, commentController.AddCommentHandler)
	// 删除评论
	r.DELETE("/api/comments/delete/:commentId", authController.JwtMiddleware, commentController.DeleteCommentHandler)
	
	// 标签相关
	r.GET("/api/tags", tagController.GetTagList)
	
	// 获取验证码
	r.GET("/api/captcha", captchaController.GetCaptchaHandler)
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
