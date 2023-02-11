package service

import (
	"moss/domain/core/repository"
	"moss/domain/core/repository/context"
	"moss/infrastructure/general/message"
)

var Mapping = new(MappingService)

type MappingService struct {
}

// CreateArticleTag 创建文章标签
func (s *MappingService) CreateArticleTag(articleID, tagID int) error {
	if articleID == 0 || tagID == 0 {
		return message.ErrIdRequired
	}
	return repository.Mapping.CreateArticleTag(articleID, tagID)
}

// DeleteArticleTag 删除文章标签
func (s *MappingService) DeleteArticleTag(articleID, tagID int) error {
	if articleID == 0 || tagID == 0 {
		return message.ErrIdRequired
	}
	return repository.Mapping.DeleteArticleTag(articleID, tagID)
}

// DeleteArticleTagByTagIds 删除文章标签
func (s *MappingService) DeleteArticleTagByTagIds(articleID int, tagIds []int) error {
	if articleID == 0 || len(tagIds) == 0 {
		return message.ErrIdRequired
	}
	return repository.Mapping.DeleteArticleTagByTagIds(articleID, tagIds)
}

// DeleteArticle 删除文章
func (s *MappingService) DeleteArticle(articleID int) error {
	if articleID == 0 {
		return message.ErrIdRequired
	}
	return repository.Mapping.DeleteArticle(articleID)
}

// DeleteTag 删除标签
func (s *MappingService) DeleteTag(tagID int) error {
	if tagID == 0 {
		return message.ErrIdRequired
	}
	return repository.Mapping.DeleteTag(tagID)
}

// ListArticleIdsByTagIds 通过标签ID查询文章ID列表
func (s *MappingService) ListArticleIdsByTagIds(ctx *context.Context, tagIds []int) (res []int, err error) {
	if len(tagIds) == 0 {
		return
	}
	return repository.Mapping.ListArticleIdsByTagIds(ctx, tagIds)
}

// ListTagIdByArticleID 通过标签ID查询标签ID列表
func (s *MappingService) ListTagIdByArticleID(ctx *context.Context, articleID int) (res []int, err error) {
	return s.ListTagIdByArticleIds(ctx, []int{articleID})
}

// ListTagIdByArticleIds 通过文章ID获取标签ID列表
func (s *MappingService) ListTagIdByArticleIds(ctx *context.Context, ids []int) (res []int, err error) {
	if len(ids) == 0 {
		return
	}
	return repository.Mapping.ListTagIdByArticleIds(ctx, ids)
}

// CountByTagID 根据标签ID统计文章数
func (s *MappingService) CountByTagID(tagID int) (int64, error) {
	return repository.Mapping.CountByTagID(tagID)
}
