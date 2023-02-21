package repository

import (
	"moss/domain/core/entity"
	"moss/domain/core/repository/context"
	"moss/domain/core/repository/gormx"
	"moss/infrastructure/persistent/db"
)

var Mapping = new(MappingRepo)

type MappingRepo struct {
}

func (r *MappingRepo) MigrateTable() error {
	return db.DB.AutoMigrate(&entity.MappingTag{})
}

// CreateArticleTag 创建文章标签映射
func (r *MappingRepo) CreateArticleTag(articleID, tagID int) error {
	return db.DB.Create(&entity.MappingTag{ArticleID: articleID, TagID: tagID}).Error
}

// DeleteArticleTag 删除文章标签映射 通过文章ID和标签ID
func (r *MappingRepo) DeleteArticleTag(articleID, tagID int) error {
	return db.DB.Where("article_id = ? and tag_id = ?", articleID, tagID).Delete(&entity.MappingTag{}).Error
}

// DeleteArticleTagByTagIds 删除文章标签映射 通过文章ID和标签id列表
func (r *MappingRepo) DeleteArticleTagByTagIds(articleID int, tagIds []int) error {
	return db.DB.Where("article_id = ? and tag_id in ?", articleID, tagIds).Delete(&entity.MappingTag{}).Error
}

// DeleteArticle 删除文章
func (r *MappingRepo) DeleteArticle(articleID int) error {
	return db.DB.Where("article_id = ?", articleID).Delete(&entity.MappingTag{}).Error
}

// DeleteTag 删除标签
func (r *MappingRepo) DeleteTag(tagID int) error {
	return db.DB.Where("tag_id = ?", tagID).Delete(&entity.MappingTag{}).Error
}

// ListArticleIdsByTagIds 通过标签ID查询文章ID列表
func (r *MappingRepo) ListArticleIdsByTagIds(ctx *context.Context, tagIds []int) (articleIds []int, err error) {
	err = db.DB.Model(&entity.MappingTag{}).Scopes(gormx.Context(ctx)).Where("tag_id in ?", tagIds).Pluck("article_id", &articleIds).Error
	return
}

// ListTagIdByArticleIds 通过文章ID获取tagID列表
func (r *MappingRepo) ListTagIdByArticleIds(ctx *context.Context, articleIds []int) (tagIds []int, err error) {
	err = db.DB.Model(&entity.MappingTag{}).Scopes(gormx.Context(ctx)).Where("article_id in ?", articleIds).Pluck("tag_id", &tagIds).Error
	return
}

// CountByTagID 根据tagID统计
func (r *MappingRepo) CountByTagID(tagID int) (res int64, err error) {
	err = db.DB.Model(entity.MappingTag{}).Where("tag_id = ?", tagID).Count(&res).Error
	return
}
