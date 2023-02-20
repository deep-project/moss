package dto

import (
	"moss/domain/core/entity"
)

// ArticlePost 文章提交模型，可直接创建tags
type ArticlePost struct {
	entity.Article
	Tags         []string `json:"tags"`
	CategoryName string   `json:"category_name"`
	UniqueTitle  bool     `json:"unique_title"`
	UniqueSource bool     `json:"unique_source"`
}
