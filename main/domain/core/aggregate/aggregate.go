package aggregate

import (
	"moss/domain/core/entity"
)

type EntityInterface interface {
	entity.Article | entity.Category | entity.Tag | entity.Link | entity.Store
}

type CategoryTree struct {
	entity.Category
	Children []CategoryTree `json:"children"`
}

// ArticlePost 文章提交模型
type ArticlePost struct {
	entity.Article
	Tags         []string `json:"tags"`          // 直接创建tags
	CategoryName string   `json:"category_name"` // 直接通过分类名创建，优先级小于category_id
}
