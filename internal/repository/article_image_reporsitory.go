package repository

import (
	"icu/config"

	"gorm.io/gorm"
)

// ArticleImage 模型结构
type ArticleImage struct {
	gorm.Model
	ArticleID uint64 `json:"articleId"`
	ImageKey  string `json:"imageKey"`
	Caption   string `json:"caption"`
	SortOrder int    `json:"sortOrder"`
}

type ArticleImageRepository struct {
	db *gorm.DB
}

func NewArticleImageRepository() *ArticleImageRepository {
	return &ArticleImageRepository{db: config.GetDB()}
}

// 1. 批量新增（一次性写入，性能好）
func (r *ArticleImageRepository) BatchCreate(images []ArticleImage) error {
	if len(images) == 0 {
		return nil
	}
	return r.db.Create(&images).Error
}

func (r *ArticleImageRepository) FindMapByArticleIds(articleIds []uint64) (map[uint64][]ArticleImage, error) {
	res := make(map[uint64][]ArticleImage, len(articleIds))
	if len(articleIds) == 0 {
		return res, nil
	}

	var images []ArticleImage
	err := r.db.
		Where("article_id IN ?", articleIds).
		Order("article_id, sort_order"). // 先按文章，再按顺序
		Find(&images).Error
	if err != nil {
		return nil, err
	}

	// 内存分组
	for _, img := range images {
		res[img.ArticleID] = append(res[img.ArticleID], img)
	}
	return res, nil
}
