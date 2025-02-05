package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

// FileController 上传照片到指定的文件夹
type FileController struct {

}

func NewFileController() *FileController{
	return &FileController{}
}

// UpLoadFile 上传文件并修改为唯一的文件名
func (a *FileController) UpLoadFile(c *gin.Context) {
	// 创建保存文件的目录（如果不存在）
	uploadDir := "../uploads"
	// 单文件上传
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "无法获取文件",
		})
		return
	}

	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	// 使用当前时间戳和随机数生成唯一的文件名，保持原文件的扩展名
	rand.Seed(uint64(time.Now().UnixNano())) // 初始化随机数种子
	ext := filepath.Ext(file.Filename) // 获取文件的扩展名
	newFileName := generateUniqueFileName(ext) // 生成唯一文件名

	// 生成文件保存路径
	filePath := filepath.Join(uploadDir, newFileName)

	// 保存文件到服务器
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "文件保存失败",
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message":  "文件上传成功",
		"filePath": filePath,
	})
}

// generateUniqueFileName 生成全数字的唯一文件名
func generateUniqueFileName(ext string) string {
	// 获取当前时间戳
	timestamp := time.Now().UnixNano()

	// 生成随机数
	randomNumber := rand.Intn(10000) // 随机生成一个 4 位数的数字

	// 将时间戳和随机数转换为数字字符串
	return fmt.Sprintf("%d%d%s", timestamp, randomNumber, ext)
}

func (a *FileController) GetFile(c *gin.Context) {
	// 设置图片存放的目录
	uploadDir := "../uploads"
	// 获取请求的图片文件名
	fileName := c.Param("filename")
	if fileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "文件名不能为空",
		})
		return
	}
	filePath := filepath.Join(uploadDir, fileName)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "文件不存在",
		})
		return
	}

	// 设置响应头，返回文件
	c.File(filePath)
}
