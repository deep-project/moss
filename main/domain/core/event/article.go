package event

import (
	"moss/domain/core/entity"
)

type ArticleCreateBefore interface {
	ArticleCreateBefore(*entity.Article) error
}

type ArticleCreateAfter interface {
	ArticleCreateAfter(*entity.Article)
}

type ArticleUpdateBefore interface {
	ArticleUpdateBefore(item *entity.Article) error
}

type ArticleUpdateAfter interface {
	ArticleUpdateAfter(item *entity.Article)
}

type ArticleDeleteBefore interface {
	ArticleDeleteBefore(id int) error
}

type ArticleDeleteAfter interface {
	ArticleDeleteAfter(id int)
}

type ArticleGetAfter interface {
	ArticleGetAfter(*entity.Article)
}

type ArticleListAfter interface {
	ArticleListAfter([]entity.ArticleBase)
}
