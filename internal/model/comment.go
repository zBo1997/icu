package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ArticleID     uint64    `json:"articleId"`                  // 文章ID
	UserID        uint64    `json:"userId"`                     // 用户ID
	AvatarUrl     string    `json:"avatarUrl"`                  // 用户头像
	UserName      string    `json:"name"`                       // 用户名称
	Comment       string    `json:"comment"`                    // 评论内容
	ParentID      *uint64   `json:"parentId,omitempty"`         // 父评论ID（引用回复）
	ParentName    string    `json:"parentName"`                 // 父评论用户名称（便于前端展示）
	ReplyToUserId *uint64   `json:"replyToUserId,omitempty"`    // 父评论ID（引用回复）
	ReplyUserName string    `json:"replyUserName,omitempty"`    // 父评论ID（引用回复）
	LikesCount    int       `json:"likesCount"`                 // 点赞数量
	Replies       []Comment `gorm:"-" json:"replies,omitempty"` // 子评论，不映射到数据库
}
