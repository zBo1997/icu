package service

import (
	"icu/internal/model"
	"icu/internal/repository"
)

type CommentService struct {
	commentRepo *repository.CommentRepository
}

// NewCommentService 创建 CommentService 实例
func NewCommentService() *CommentService {
	return &CommentService{
		commentRepo: repository.NewCommentRepository(),
	}
}

// 添加评论
func (c *CommentService) AddComment(comment	*model.Comment) error {
	// model 转换为评论结构体
	 commentModel := &repository.Comment{
	 	ArticleID:  comment.ArticleID,
	 	UserID:     comment.UserID,
	 	Comment:    comment.Comment,
	 	ParentID:   comment.ParentID,
	 	LikesCount: comment.LikesCount,
	 }
	
	return c.commentRepo.AddComment(commentModel)
}

// 根据文章ID获取评论列表
func (c *CommentService) GetCommentsByArticleID(articleID int) ([]model.Comment, error) {
	return c.commentRepo.GetCommentsByArticleID(articleID)
}

// 删除评论
func (c *CommentService) DeleteComment(commentID int64) error {
	return c.commentRepo.DeleteComment(commentID)
}

// 更新评论点赞数量
func (c *CommentService) UpdateLikesCount(commentID int64, increment int) error {
	return c.commentRepo.UpdateLikesCount(commentID, increment)
}
