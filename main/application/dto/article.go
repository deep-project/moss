package dto

import (
	"moss/domain/core/entity"
)

// ArticleCreate 文章创建模型，可直接创建tags
type ArticleCreate struct {
	entity.Article
	Tags         []string `json:"tags"`
	CategoryName string   `json:"category_name"`
}
