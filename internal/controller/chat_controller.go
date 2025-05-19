package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"icu/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// ChatController 用于处理聊天相关的业务逻辑
type ChatController struct {
}

func NewChatController() *ChatController {
	return &ChatController{}
}

func (a *ChatController) ChatAI(c *gin.Context) {
    conversationId := c.Query("conversationId")
    if conversationId == "" {
        conversationId = fmt.Sprintf("%d", time.Now().UnixNano())
    }

    c.Header("Content-Type", "text/event-stream")
    c.Header("Cache-Control", "no-cache")
    c.Header("Connection", "keep-alive")

    // 获取用户问题（从 query 参数）
    question := c.Query("content")
    if question == "" {
        question = "你好，你是ICU的AI助手，请问有什么可以帮助你的吗？"
    }

    // 初始化 openai 客户端
    client := openai.NewClient(
        option.WithAPIKey("sk-04e1168b173d41cfb356157544a6fee7"),
        option.WithBaseURL("https://api.deepseek.com"),
    )

    stream := client.Chat.Completions.NewStreaming(context.TODO(), openai.ChatCompletionNewParams{
        Messages: []openai.ChatCompletionMessageParamUnion{
            openai.UserMessage(question),
        },
        Model: "deepseek-chat",
    })

    acc := openai.ChatCompletionAccumulator{}
    clientGone := c.Request.Context().Done()

    // 流式推送AI回复
    go func() {
        defer stream.Close()
        for stream.Next() {
            select {
            case <-clientGone:
                log.Println("客户端断开连接，停止推送")
                return
            default:
                chunk := stream.Current()
                acc.AddChunk(chunk)
                var content string
                if len(acc.Choices) > 0 {
                    content = acc.Choices[0].Message.Content
                }
                message := model.Message{
                    Id:             time.Now().UnixNano(),
                    ConversationId: conversationId,
                    Type:           "text",
                    Content:        content,
                    IsEnd:          false,
                    Timestamp:      time.Now().Format(time.RFC3339),
                    Sender:         "system",
                }
                jsonMessage, _ := json.Marshal(message)
                c.SSEvent("message", string(jsonMessage))
                c.Writer.Flush()
            }
        }
        // 结束消息
        message := model.Message{
            Id:             time.Now().UnixNano(),
            ConversationId: conversationId,
            Type:           "text",
            Content:        acc.Choices[0].Message.Content,
            IsEnd:          true,
            Timestamp:      time.Now().Format(time.RFC3339),
            Sender:         "system",
        }
        jsonMessage, _ := json.Marshal(message)
        c.SSEvent("message", string(jsonMessage))
        c.Writer.Flush()
    }()

    <-clientGone
    log.Println("客户端断开连接")
}