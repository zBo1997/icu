package repository

import (
	"icu/config"
	"icu/internal/model"

	"gorm.io/gorm"
)

// 添加评论结构体
type Comment struct {
	gorm.Model
	ArticleID  int64  `json:"articleId"`          // 文章ID
	UserID     int64  `json:"userId"`             // 用户ID
	Comment    string `json:"comment"`            // 评论内容
	ParentID   *int64 `json:"parentId,omitempty"` // 父评论ID（引用回复）
	LikesCount int    `json:"likesCount"`         // 点赞数量
}

type CommentRepository struct {
	db *gorm.DB
}


// NewCommentRepository 创建一个新的 CommentRepository 实例
func NewCommentRepository() *CommentRepository {
	return &CommentRepository{db: config.GetDB()}
}


//查询单挑评论
func (r *CommentRepository) GetCommentByID(commentID uint) (model.Comment, error) {
	var comment model.Comment

    // 查询单条评论并关联用户信息和父评论用户名称
    err := r.db.Table("comments").
        Select("comments.*, users.name as user_name, users.avatar as avatar_url, parent_user.name as parent_name").
        Joins("LEFT JOIN users ON comments.user_id = users.id").
        Joins("LEFT JOIN comments as parent_comment ON comments.parent_id = parent_comment.id").
        Joins("LEFT JOIN users as parent_user ON parent_comment.user_id = parent_user.id").
        Where("comments.id = ?", commentID).
        First(&comment).Error

    if err != nil {
        return model.Comment{}, err
    }
    return comment, nil
}


// AddComment 添加一条新的评论
func (r *CommentRepository) AddComment(comment *Comment) (uint, error) {
	//返回添加后的评论ID
	err := r.db.Create(comment).Error
	if err != nil {
		return 0, err
	}
	return comment.ID, nil
}

// GetCommentsByArticleID 根据文章ID获取所有评论（支持引用回复）
func (r *CommentRepository) GetCommentsByArticleID(articleID int) ([]model.Comment, error) {
	var comments []model.Comment

	// 查询评论并关联用户信息和父评论用户名称
	err := r.db.Table("comments").
		Select("comments.*, users.name as user_name, users.avatar as avatar_url, parent_user.name as parent_name").
		Joins("LEFT JOIN users ON comments.user_id = users.id").
		Joins("LEFT JOIN comments as parent_comment ON comments.parent_id = parent_comment.id").
		Joins("LEFT JOIN users as parent_user ON parent_comment.user_id = parent_user.id").
		Where("comments.article_id = ?", articleID).
		Order("comments.parent_id ASC, comments.created_at ASC"). // 按父评论和时间排序
		Find(&comments).Error

	if err != nil {
		return nil, err
	}
	return comments, nil
}

// DeleteComment 删除一条评论
func (r *CommentRepository) DeleteComment(commentID int64) error {
	return r.db.Delete(&model.Comment{}, commentID).Error
}

// UpdateLikesCount 更新评论的点赞数量
func (r *CommentRepository) UpdateLikesCount(commentID int64, increment int) error {
	return r.db.Model(&model.Comment{}).
		Where("id = ?", commentID).
		Update("likes_count", gorm.Expr("likes_count + ?", increment)).Error
}
