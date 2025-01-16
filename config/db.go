package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

// InitConfig 初始化配置文件
func InitConfig() {
	// 获取当前工作目录
	dir, pathErr := os.Getwd()
	if pathErr != nil {
		fmt.Println("Error:", pathErr)
		return
	}
	rootPath, _ := findProjectRoot(dir)
	// 获取当前工作目录
	viper.SetConfigName("config")    // 配置文件名（不带扩展名）
	viper.AddConfigPath(rootPath)         // 配置文件所在的路径
	viper.SetConfigType("yml")      // 配置文件类型为 YAML

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
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 检查数据库连接
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("获取数据库实例失败: %v", err)
	}

	// 设置最大连接池数
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)

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


// 查找项目根目录
func findProjectRoot(startDir string) (string, error) {
	for {
		// 查找 go.mod 文件
		if _, err := os.Stat(filepath.Join(startDir, "go.mod")); err == nil {
			return startDir, nil
		}

		// 如果已经到达根目录，就退出
		parentDir := filepath.Dir(startDir)
		if parentDir == startDir {
			return "", fmt.Errorf("project root not found")
		}

		// 继续向上查找
		startDir = parentDir
	}
}