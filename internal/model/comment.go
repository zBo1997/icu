package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ArticleID  int64  `json:"article_id"`
	UserID     int64  `json:"user_id"`
	Comment    string `json:"comment"`
	ParentID   *int64 `json:"parent_id,omitempty"` // 父评论ID
	LikesCount int    `json:"likes_count"`
}
