package service

import (
	"icu/config"
	"icu/internal/model"
	"icu/internal/repository"
	"strconv"
	"sync"
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
	users, articleTags, imageMap, err := a.assembleArticleExtraInfo(articles)
	if err != nil {
		return modelArticles, 0, err
	}
	//把Article转换为model.Article
	for _, article := range articles {
		user := users[article.UserId]
		tags := articleTags[uint64(article.ID)]
		names := make([]string, len(tags))
		images := imageMap[uint64(article.ID)]
		for i, t := range tags {
			names[i] = t.Tag
		}
		keys := make([]model.ArticleWithImage, len(images))
		for i, img := range images {
			keys[i] = model.ArticleWithImage{
				ID:        uint64(img.ID),
				ArticleID: img.ArticleID,
				Caption:   img.Caption,
				ImageKey:  img.ImageKey,
				SortOrder: img.SortOrder,
				CreatedAt: img.CreatedAt,
				UpdatedAt: img.UpdatedAt,
			}
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
			CreatedAt: article.CreatedAt,
			UpdatedAt: article.UpdatedAt,
		}
		modelArticles = append(modelArticles, modelArticle)
	}
	return modelArticles, total, nil
}

// 根据文章编号获取文章信息
func (a *ArticleService) GetArticle(articleId int) (model.Article, error) {
	var modelArticle model.Article
	//查询文章信息
	article, err := a.articleRepo.GetArticle(articleId)
	if err != nil {
		return model.Article{}, err
	}
	user, err := a.userRepository.GetUserByID(strconv.FormatUint(article.UserId, 10))
	if err != nil {
		return model.Article{}, err
	}
	tags, err := a.articleTagRepository.FindTagsByArticleId(uint64(article.ID))
	if err != nil {
		return model.Article{}, err
	}
	names := make([]string, len(tags))
	for i, t := range tags {
		names[i] = t.Tag
	}
	modelArticle = model.Article{
		ID:        uint64(article.ID),
		Title:     article.Title,
		Content:   article.Content,
		Name:      user.Name,
		AvatarUrl: user.Avatar,
		TagNames:  names,
		UserId:    uint64(user.ID),
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}
	return modelArticle, nil
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

// channel并发组装文章的额外信息，比如标签，图片等
func (a *ArticleService) assembleArticleExtraInfo(articles []repository.Article) (users map[uint64]repository.User, articleTags map[uint64][]repository.Tag, imageMap map[uint64][]repository.ArticleImage, err error) {
	users = make(map[uint64]repository.User)
	articleTags = make(map[uint64][]repository.Tag)
	imageMap = make(map[uint64][]repository.ArticleImage)

	var (
		waitGroup sync.WaitGroup
		mu        sync.Mutex
		once      sync.Once
	)

	userIds := make([]uint64, len(articles))
	for i, article := range articles {
		userIds[i] = article.UserId
	}

	articleIds := make([]uint64, len(articles))
	for i, article := range articles {
		articleIds[i] = uint64(article.ID)
	}
	waitGroup.Add(3)

	go func() {
		defer waitGroup.Done()
		result, e := a.userRepository.FindMapByIds(userIds)
		if e != nil {
			once.Do(func() {
				err = e
			})
			return
		}
		mu.Lock()
		users = result
		mu.Unlock()
	}()

	go func() {
		defer waitGroup.Done()
		result, e := a.articleTagRepository.FindTagMapByArticleIds(articleIds)
		if e != nil {
			once.Do(func() {
				err = e
			})
		}
		mu.Lock()
		articleTags = result
		mu.Unlock()
	}()

	go func() {
		defer waitGroup.Done()
		result, e := a.articleImageRepo.FindMapByArticleIds(articleIds)
		if e != nil {
			once.Do(func() {
				err = e
			})
			return
		}
		mu.Lock()
		imageMap = result
		mu.Unlock()
	}()

	waitGroup.Wait()

	return users, articleTags, imageMap, nil
}
