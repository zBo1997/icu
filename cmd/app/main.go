package main

import (
	"icu/config"
	"icu/internal/route"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置文件
	config.InitConfig()
	// 初始化数据库
	config.InitDB()
	// 程序退出时关闭数据库连接
	defer config.CloseDB()


	r := gin.Default()

	// 自定义中间件处理错误并记录日志
	r.Use(func(c *gin.Context) {
		// 捕获所有发生的错误和 panic
		defer func() {
			if err := recover(); err != nil {
				// 如果发生 panic，记录 ERROR 级别的日志
				log.Printf("[ERROR] %s %s %s", c.Request.Method, c.Request.URL.Path,c.Errors)
				// 发送 500 错误响应
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "未知错误,请联系我们",
				})
			}
		}()
		// 执行请求
		c.Next()
	})

	route.SetupRoutes(r)

	// 启动服务
	r.Run(":8080")
}
