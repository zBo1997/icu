package controller

import (
	"icu/internal/model"
	"icu/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService *service.CommentService
}

// NewCommentController 创建 CommentController 实例
func NewCommentController() *CommentController {
	return &CommentController{
		commentService: service.NewCommentService(),
	}
}

// AddCommentHandler 处理添加评论的请求
func (cc *CommentController) AddCommentHandler(c *gin.Context) {
	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "非法请求参数"})
		return
	}

	if err := cc.commentService.AddComment(comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "评论失败，请稍后重试"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "评论成功"})
}

// GetCommentsHandler 处理获取文章评论的请求
func (cc *CommentController) GetCommentsHandler(c *gin.Context) {
	articleIDStr := c.Query("article_id")
	articleID, err := strconv.ParseInt(articleIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章编号"})
		return
	}

	comments, err := cc.commentService.GetCommentsByArticleID(articleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取评论失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

// DeleteCommentHandler 处理删除评论的请求
func (cc *CommentController) DeleteCommentHandler(c *gin.Context) {
	commentIDStr := c.Param("id")
	commentID, err := strconv.ParseInt(commentIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	if err := cc.commentService.DeleteComment(commentID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除评论失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除评论成功"})
}

// UpdateLikesHandler 处理更新评论点赞数量的请求
func (cc *CommentController) UpdateLikesHandler(c *gin.Context) {
	commentIDStr := c.Param("id")
	commentID, err := strconv.ParseInt(commentIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的评论编号"})
		return
	}

	var payload struct {
		Increment int `json:"increment"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := cc.commentService.UpdateLikesCount(commentID, payload.Increment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update likes count"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "点赞成功"})
}
