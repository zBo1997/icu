package controller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// UserService 用于处理与用户相关的业务逻辑
type ChatController struct {
}

func NewChatController() *ChatController {
	return &ChatController{}
}

// GetUserHandler 获取用户信息的处理函数
func (a *ChatController) ChatAI(c *gin.Context) {
	// 设置响应头
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	// 创建一个通道，用于向客户端推送消息
	messageChan := make(chan string)

	// 启动一个 Goroutine，定期向客户端推送消息
	arr := []int{1, 2, 3}
	go func() {
		for _, _ := range arr {
			message := fmt.Sprintf("系统消息: %s", time.Now().Format("2006-01-02 15:04:05"))
			messageChan <- message
			time.Sleep(2 * time.Second) // 每隔 2 秒推送一次
		}
	}()

	// 监听客户端断开连接
	clientGone := c.Request.Context().Done()
	for {
		select {
		case <-clientGone:
			fmt.Println("客户端断开连接")
			return
		case message := <-messageChan:
			// 推送消息到客户端
			c.SSEvent("message", message)
			c.Writer.Flush() // 立即刷新响应
		}
	}
}
