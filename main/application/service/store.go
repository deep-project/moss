package service

import (
	"moss/application/mapper"
	"moss/domain/core/aggregate"
	"moss/domain/core/service"
)

// StorePost 发布仓库文章
func StorePost(id int) (post *aggregate.ArticlePost, err error) {
	item, err := service.Store.Get(id)
	if err != nil {
		return
	}
	post = mapper.StoreToArticlePost(item)
	if err = ArticlePost("create", post); err != nil {
		return
	}
	return post, service.Store.Delete(id)
}
