package main

import (
	"context"
	"icu/config"
	"icu/internal/route"
	"os"
	"os/signal"
	"syscall"
	"time"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置文件
	config.InitConfig()
	// 初始化数据库
	config.InitDB()


	r := gin.Default()


	// 创建一个 http.Server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}


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

    go func() {
        // service connections
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %s\n", err)
        }
    }()
 
    // Wait for interrupt signal to gracefully shutdown the server with
    // a timeout of 5 seconds.
    quit := make(chan os.Signal, 1)
    // kill (no param) default send syscanll.SIGTERM
    // kill -2 is syscall.SIGINT
    // kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Shutdown Server ...")
 
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("Server Shutdown:", err)
    }
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("Closing database connection...")
	config.CloseDB()
    log.Println("Server exiting")
}
