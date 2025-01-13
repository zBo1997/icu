package main

import (
	"icu/internal/route"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 使用 gin.Default()，启用内置的 Logger 和 Recovery 中间件
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
