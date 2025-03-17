package controller

import (
	"encoding/json"
	"icu/internal/model"
	"log"
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
	messageChan := make(chan model.Message)

	// 监听客户端断开连接
	clientGone := c.Request.Context().Done()

	// 启动一个 Goroutine，模拟生成消息
	go func() {
		defer close(messageChan) // 关闭通道

		// 模拟多段对话
		messages := []string{
			"你好，我的名字叫小猪/n/n",
			"请问？/n/n",
			"我有什么可以帮助你的？",
		}

		for i, content := range messages {
			select {
			case <-clientGone:
				log.Println("客户端断开连接，停止推送")
				return
			default:
				// 构造消息
				message := model.Message{
					Id: 	  i + 1,
					ConversationId: "123456",
					Type:      "text",
					Content:   content,
					IsEnd:     i == len(messages)-1, // 最后一条消息标记为结束
					Timestamp: time.Now().Format(time.RFC3339),
				}
				messageChan <- message
				time.Sleep(2 * time.Second) // 模拟延迟
			}
		}
	}()

	// 监听消息和客户端断开事件
	for {
		select {
		case <-clientGone:
			log.Println("客户端断开连接")
			return
		case message, ok := <-messageChan:
			if !ok {
				log.Println("消息通道已关闭")
				return
			}
			// 将消息转换为 JSON
			jsonMessage, err := json.Marshal(message)
			if err != nil {
				log.Println("JSON 编码失败:", err)
				return
			}
			// 推送消息到客户端
			c.SSEvent("message", string(jsonMessage))
		    c.Writer.Flush();

			// 如果消息标记为结束，则关闭连接
			if message.IsEnd {
				log.Println("会话结束，关闭连接")
				return
			}
		}
	}
}