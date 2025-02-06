package config

import (
	"bytes"
	"embed"
	_ "embed"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//go:embed config.yml
var configFile embed.FS

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

// InitConfig 初始化配置文件
func InitConfig() {
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
}

// InitDB 初始化数据库连接
func InitDB(logFile *os.File) {
	// 从 viper 获取配置
	dbConfig := viper.Sub("database")
	if dbConfig == nil {
		log.Fatal("数据库配置缺失")
	}

	username := dbConfig.GetString("username")
	password := dbConfig.GetString("password")
	host := dbConfig.GetString("host")
	port := dbConfig.GetInt("port")
	dbname := dbConfig.GetString("dbname")

	// 构建 DSN（数据源名称）
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

	// 创建 GORM 的日志记录器
	gormLogger := logger.New(
		log.New(logFile, "\r\n", log.LstdFlags), // 将日志输出到文件
		logger.Config{
			LogLevel:      logger.Info,  // 设置日志级别
			SlowThreshold: time.Second,  // 慢查询阈值
			Colorful:      false,        // 关闭彩色输出
		},
	)

	// 初始化数据库连接
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 检查数据库连接
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("获取数据库实例失败: %v", err)
	}

	// 设置最大连接池数
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(3)

	log.Print("成功连接到数据库")
}

// CloseDB 关闭数据库连接
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("获取数据库实例失败: %v", err)
	}
	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("关闭数据库连接失败: %v", err)
	}
	log.Println("数据库连接已关闭")
}