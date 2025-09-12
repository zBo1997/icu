package service

import (
	"icu/config"
	"icu/internal/model"
	"icu/internal/repository"
)

// UserService 用于处理与用户相关的业务逻辑
type ArticleService struct {
	userRepository       *repository.UserRepository
	articleRepo          *repository.ArticleRepository
	articleImageRepo     *repository.ArticleImageRepository
	articleTagRepository *repository.ArticleTagRepository
}

// NewUserService 创建 UserService 实例
func NewArticleService() *ArticleService {
	return &ArticleService{
		userRepository:       repository.NewUserRepository(),
		articleRepo:          repository.NewArticleRepository(),
		articleImageRepo:     repository.NewArticleImageRepository(),
		articleTagRepository: repository.NewArticleTagRepository(),
	}
}

// 分页获取文章列表
func (a *ArticleService) PageArticle(page, pageSize int) ([]model.Article, int64, error) {
	var modelArticles []model.Article
	//计算偏移量 当前页数 * 每页数量
	offset := (page - 1) * pageSize
	articles, total, err := a.articleRepo.FindArticle(offset, pageSize)
	if err != nil {
		return modelArticles, 0, err
	}
	if len(articles) == 0 {
		return []model.Article{}, total, nil
	}
	userIdMap := make(map[uint64]struct{})
	articleIds := make([]uint64, 0, len(articles))
	for _, v := range articles {
		if v.UserId != 0 {
			userIdMap[v.UserId] = struct{}{}
		}
		articleIds = append(articleIds, uint64(v.ID))
	}
	userIds := make([]uint64, 0)
	for id := range userIdMap {
		userIds = append(userIds, id)
	}
	users, err := a.userRepository.FindMapByIds(userIds)
	articleTags, err := a.articleTagRepository.FindTagMapByArticleIds(articleIds)
	imageMap, err := a.articleImageRepo.FindMapByArticleIds(articleIds)
	//把Article转换为model.Article
	for _, article := range articles {
		user := users[article.UserId]
		tags := articleTags[uint64(article.ID)]
		names := make([]string, len(tags))
		images := imageMap[uint64(article.ID)]
		for i, t := range tags {
			names[i] = t.Tag
		}
		keys := make([]string, len(images))
		for i, img := range images {
			keys[i] = img.ImageKey
		}
		modelArticle := model.Article{
			ID:        uint64(article.ID),
			Title:     article.Title,
			Content:   article.Content,
			Name:      user.Name,
			AvatarUrl: user.Avatar,
			TagNames:  names,
			UserId:    uint64(user.ID),
			ImageKeys: keys,
		}
		modelArticles = append(modelArticles, modelArticle)
	}
	return modelArticles, total, nil
}

// 根据文章编号获取文章信息
func (a *ArticleService) GetArticle(articleId int) (model.ArticleWithImage, error) {
	//查询文章信息
	return a.articleRepo.GetArticle(articleId)
}

// 发布文章
func (a *ArticleService) PublishArticle(article *model.ArticlePublish, userId uint64) (uint64, error) {

	// 开启事务
	tx := config.GetDB().Begin()
	//创建发布文章对象
	articleModel := repository.Article{
		Title:   article.Title,
		Content: article.Content,
		UserId:  userId,
	}
	//保存文章
	articleId, err := a.articleRepo.CreateArticle(&articleModel)
	if err != nil {
		tx.Rollback() // 回滚事务
		return 0, err
	}

	tag := &repository.ArticleTag{
		ArticleId: articleId,
		TagId:     article.Tags,
	}

	//保存文章标签关系
	err = a.articleTagRepository.CreateArticleTag(articleId, tag)
	if err != nil {
		tx.Rollback() // 回滚事务
		return 0, err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}
	return articleId, nil
}
