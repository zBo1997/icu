package model

// Article 模型结构
type ArticlePublish struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Tags     uint   `json:"tags"`
	ImageKey string `json:"imageKey"`
}
