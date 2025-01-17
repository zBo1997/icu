package config

import (
	"github.com/spf13/viper"
)

func GetKey(key string) string {
	// 从 viper 获取 JWT 密钥
	return viper.GetString(key)
}