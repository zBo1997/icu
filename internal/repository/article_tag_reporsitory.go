package repository

import (
	"icu/config"

	"gorm.io/gorm"
)

type ArticleTag struct {
	gorm.Model
	ArticleId uint64 `json:"articleId"`
	TagId     uint64 `json:"tagId"`
}

type ArticleTagRepository struct {
	db *gorm.DB
}

func NewArticleTagRepository() *ArticleTagRepository {
	return &ArticleTagRepository{db: config.GetDB()}
}

// 根据文章ID查询标签
func (r *ArticleTagRepository) FindTagsByArticleId(articleId uint64) ([]Tag, error) {
	var tags []Tag
	err := r.db.
		Model(&ArticleTag{}).
		Select("tags.id, tags.tag").
		Joins("JOIN tags ON tags.id = article_tags.tag_id").
		Where("article_tags.article_id = ?", articleId).
		Order("tags.id").
		Scan(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *ArticleTagRepository) FindTagMapByArticleIds(articleIds []uint64) (map[uint64][]Tag, error) {
	res := make(map[uint64][]Tag, len(articleIds))
	if len(articleIds) == 0 {
		return res, nil
	}
	var rows []struct {
		ArticleId uint64
		Tag
	}
	err := r.db.
		Model(&ArticleTag{}).
		Select("article_tags.article_id, tags.id, tags.tag").
		Joins("JOIN tags ON tags.id = article_tags.tag_id").
		Where("article_tags.article_id IN ?", articleIds).
		Order("article_tags.article_id, tags.id").
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}

	for _, r := range rows {
		res[r.ArticleId] = append(res[r.ArticleId], r.Tag)
	}
	return res, nil
}

// 保存文章和标签的关系
func (t *ArticleTagRepository) CreateArticleTag(articleId uint64, param *ArticleTag) error {
	var tags []ArticleTag
	err := t.db.Table("tags").Where("id = ?", param.TagId).Find(&tags).Error
	if err != nil {
		return err
	}
	//将标签插入到文章标签表中
	for _, tag := range tags {
		err = t.db.Table("article_tags").Create(&map[string]any{
			"article_id": articleId,
			"tag_id":     tag.ID,
		}).Error
		if err != nil {
			return err
		}
	}
	return nil
}
