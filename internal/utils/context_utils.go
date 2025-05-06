package utils

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUserIDFromContext 从 *gin.Context 获取当前用户的 userId，并转换为 int64
func GetUserIDFromContext(c *gin.Context) (int64, error) {
	// 从上下文中获取 userId
	userID, exists := c.Get("userId")
	if !exists {
		return 0, errors.New("用户未登录")
	}

	// 尝试将 userId 转换为 int64
	switch v := userID.(type) {
	case int64:
		return v, nil
	case uint:
		return int64(v), nil
	case int:
		return int64(v), nil
	case string:
		// 如果是字符串，尝试解析为 int64
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, errors.New("userId 格式错误")
		}
		return id, nil
	default:
		return 0, errors.New("无法解析 userId")
	}
}
