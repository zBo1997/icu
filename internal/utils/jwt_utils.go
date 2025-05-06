package utils

import (
	"errors"
	"icu/config"
	"log"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = config.GetKey("jwt", "secret_key") // 用于签发 JWT 的密钥

// ParseUserIDFromToken 从 JWT 中解析 userId
func ParseUserIDFromToken(tokenStr string) (uint, error) {
	if tokenStr == "" {
		return 0, errors.New("token 为空")
	}

	// 解析 token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// 检查签名方法
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("无效的签名方法")
		}
		return []byte(jwtKey), nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("无效的 token")
	}

	// 提取 Claims（确保与生成时类型一致）
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("无效的 claims 结构")
	}

	userIDStr := claims["sub"].(string)
	log.Printf("用户编号: %+v\n", userIDStr)
	if userIDStr == "" {
		return 0, errors.New("userID 不存在")
	}

	// 转换为 uint
	userIDUint, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return 0, errors.New("userID 格式错误")
	}

	return uint(userIDUint), nil
}
