package config

import (
	"bytes"
	"log"

	"github.com/spf13/viper"
)

func GetKey(key string,value string) string {
	// 从嵌入的文件系统中读取 config.yml 文件
    fileContent, err := configFile.ReadFile("config.yml")
    if err != nil {
        log.Fatalf("无法读取嵌入的配置文件: %v", err)
    }
    // 使用 viper 加载配置
    viper.SetConfigType("yml")
    err = viper.ReadConfig(bytes.NewReader(fileContent))
    if err != nil {
        log.Fatalf("读取配置文件失败: %v", err)
    }
	config := viper.Sub(key)
	// 从 viper 获取 JWT 密钥
	return config.GetString(value)
}