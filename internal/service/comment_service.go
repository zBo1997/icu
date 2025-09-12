package service

import (
	"icu/internal/model"
	"icu/internal/repository"
)

type CommentService struct {
	commentRepo *repository.CommentRepository
}

func (c *CommentService) GetCommentByID(commentID uint64) (model.Comment, error) {
	comment, err := c.commentRepo.GetCommentByID(commentID)
	if err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}

// NewCommentService 创建 CommentService 实例
func NewCommentService() *CommentService {
	return &CommentService{
		commentRepo: repository.NewCommentRepository(),
	}
}

// 添加评论
func (c *CommentService) AddComment(comment *model.Comment) (uint64, error) {
	// model 转换为评论结构体
	commentModel := &repository.Comment{
		ArticleID:     comment.ArticleID,
		UserID:        comment.UserID,
		Comment:       comment.Comment,
		ParentID:      comment.ParentID,
		ReplyToUserId: comment.ReplyToUserId, // 回复的用户ID
		LikesCount:    comment.LikesCount,
	}
	// 调用评论仓库的添加方法
	commentID, err := c.commentRepo.AddComment(commentModel)
	if err != nil {
		return 0, err
	}
	// 返回添加后的评论ID
	return uint64(commentID), nil

}

// 根据文章ID获取评论列表
func (c *CommentService) GetCommentsByArticleID(articleID int64) ([]*model.Comment, error) {
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
