package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var DB *sql.DB

// InitConfig 初始化配置文件
func InitConfig() {
	viper.SetConfigName("config")    // 配置文件名（不带扩展名）
	viper.AddConfigPath(".")         // 配置文件所在的路径
	viper.SetConfigType("yaml")      // 配置文件类型为 YAML

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
}

// InitDB 初始化数据库连接
func InitDB() {
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

	// 初始化数据库连接
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 检查数据库连接
	err = DB.Ping()
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	log.Print("成功连接到数据库")
}

// CloseDB 关闭数据库连接
func CloseDB() {
	err := DB.Close()
	if err != nil {
		log.Fatalf("关闭数据库连接失败: %v", err)
	}
}
