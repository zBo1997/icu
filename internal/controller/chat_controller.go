package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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

//生成conversationId的接口
func (a *ChatController) GenerateConversationID(c *gin.Context) {
    conversationId := fmt.Sprintf("%d", time.Now().UnixNano())
    // 返回 JSON
       c.JSON(http.StatusOK, gin.H{"data": map[string]string{
        "conversationId":    conversationId,
    }})
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
        option.WithAPIKey("sk-0090294ca0a3465d9d854385447f455a"),
        option.WithBaseURL("https://api.deepseek.com"),
    )

    stream := client.Chat.Completions.NewStreaming(context.TODO(), openai.ChatCompletionNewParams{
    Messages: []openai.ChatCompletionMessageParamUnion{
            openai.UserMessage(question),
        },
        Model: "deepseek-chat",
    })

    //acc := openai.ChatCompletionAccumulator{}
    clientGone := c.Request.Context().Done()
    // 增加id
    id := 1
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
            // 封装会话ID和内容
            data := map[string]interface{}{
                "conversationId": conversationId,
                "content":        chunk.Choices[0].Delta.Content,
                "sender":         "system",
            }
            log.Printf("Received chunk: %+v", data)
            jsonBytes, _ := json.Marshal(data)
            fmt.Fprintf(c.Writer, "id: %d\n", id)
            fmt.Fprintf(c.Writer, "event: message\n")
            fmt.Fprintf(c.Writer, "data: %s\n\n", jsonBytes)
            //增加一个全局的Id
            c.Writer.Flush()
            id++
        }
    }
    c.Writer.Flush()
    }()

    <-clientGone
    log.Println("客户端断开连接")
}