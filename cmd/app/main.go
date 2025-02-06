package main

import (
	"context"
	"fmt"
	"icu/config"
	"icu/internal/route"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 检查并处理日志文件大小
	logFile, err  := checkLogFileSize()
	// 创建日志文件
	if err != nil {
		log.Fatalf("无法创建日志文件: %v", err)
    }
	//关闭文件
    defer logFile.Close()
	// 初始化配置文件
	config.InitConfig()
	// 初始化数据库
	config.InitDB(logFile)
	// 设置日志输出到文件
	log.SetOutput(logFile)
	
	// 将 Gin 的日志输出到文件
	gin.DefaultWriter = logFile
	r := gin.New()
	r.Use(gin.LoggerWithWriter(logFile))
	r.Use(gin.Recovery())

	//限制文件上传的大小
	r.MaxMultipartMemory = 8 << 20 // 8 MB


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

const (
	logDir      = "./logs"           // 日志文件目录
	logFileName = "app.log"          // 默认日志文件名
	maxSize     = 10 * 1024 * 1024  // 10 MB
)

// 获取当前日期，格式化为 "YYYY-MM-DD" 形式
func getCurrentDate() string {
	return time.Now().Format("2006-01-02")
}

// 检查并处理日志文件的大小，超出限制则切割日志
func checkLogFileSize() (*os.File, error) {
	// 获取日志文件的路径
	logFilePath := filepath.Join(logDir, logFileName)

	// 获取日志文件信息
	info, err := os.Stat(logFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			// 如果日志文件不存在，创建日志目录并创建新文件
			os.MkdirAll(logDir, os.ModePerm)
			return createLogFile(logFilePath)
		} else {
			return nil, fmt.Errorf("无法获取日志文件信息: %v", err)
		}
	}

	// 如果文件大小超过最大限制
	if info.Size() > maxSize {
		// 获取当前日期，用于日志文件重命名
		currentDate := getCurrentDate()

		// 重命名现有日志文件
		newLogFilePath := filepath.Join(logDir, fmt.Sprintf("app-%s.log", currentDate))
		err := os.Rename(logFilePath, newLogFilePath)
		if err != nil {
			return nil, fmt.Errorf("重命名日志文件失败: %v", err)
		}
		log.Printf("日志文件已切割，旧日志已重命名为: %s", newLogFilePath)

		// 创建新的日志文件
		return createLogFile(logFilePath)
	}

	// 返回现有的日志文件
	return os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
}

// 创建日志文件并将其设置为日志输出目标
func createLogFile(filePath string) (*os.File, error) {
	// 打开文件，创建新文件
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("无法创建日志文件: %v", err)
	}
	return file, nil
}