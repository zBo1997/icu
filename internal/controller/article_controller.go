package controller

import (
	"net/http"
	"strconv"

	"icu/internal/service"

	"github.com/gin-gonic/gin"
)

// ArticleController 文章控制器，用于处理与文章相关的业务逻辑
type ArticleController struct {
	service    *service.ArticleService
	tagService *service.TagService
}

func NewArticleController() *ArticleController {
	return &ArticleController{
		service:    service.NewArticleService(),
		tagService: service.NewTagService(),
	}
}

// GetUserHandler 获取文章相关信息的处理函数
func (a *ArticleController) PageArticle(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	ariticles, total, err := a.service.PageArticle(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get notifications",
		})
		return
	}
	count := page * pageSize
	//count 转换为 int64
	countInt64 := int64(count)
	hasMore := countInt64 < total

	c.JSON(http.StatusOK, gin.H{
		"items":   ariticles,
		"hasMore": hasMore,
		"total":   total,
	})
}
// GetArticleHandler 获取文章相关信息的处理函数
func (a *ArticleController) GetArticle(c *gin.Context) {
	articleId, err := strconv.Atoi(c.Param("articleId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的文章编号",
		})
		return
	}

	article, err := a.service.GetArticle(articleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取文章是吧",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"article": article,
	})
}
