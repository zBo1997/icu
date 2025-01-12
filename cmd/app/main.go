package main

import (
	"icu/internal/route"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 注册路由
	route.SetupRoutes(r)

	// 启动服务
	if err := r.Run(":8080"); err != nil {
		panic("Failed to start the server")
	}
}
