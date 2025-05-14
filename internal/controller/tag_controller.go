package controller

import (
	"icu/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	service *service.TagService
}

func NewTagController() *TagController {
	return &TagController{
		service: service.NewTagService(),
	}
}

// GetTagListHandler 获取标签列表的处理函数
func (a *TagController) GetTagList(c *gin.Context) {
	tagList, err := a.service.GetTagList()
	if err != nil {
		c.JSON(200, gin.H{
			"message": "标签列表获取失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tagList})
}
