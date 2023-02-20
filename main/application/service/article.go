package service

import (
	"moss/application/dto"
	"moss/domain/core/repository/context"
	"moss/domain/core/service"
)

// CreatePost 提交文章
func CreatePost(action string, item *dto.ArticlePost) (err error) {
	if item.Article.CategoryID == 0 && item.CategoryName != "" {
		cate, err := CategoryGetOrCreate(item.CategoryName)
		if err != nil {
			return err
		}
		item.Article.CategoryID = cate.ID
	}
	var ctx = context.NewArticlePost(item.UniqueTitle, item.UniqueSource)
	if action == "create" {
		err = service.Article.Create(ctx, &item.Article)
	} else {
		err = service.Article.Update(ctx, &item.Article)
	}
	if err != nil {
		return err
	}
	return CreateArticleTagsByNameList(item.ID, item.Tags)
}

// CreateArticleTagsByNameList 创建文件标签列表
func CreateArticleTagsByNameList(articleID int, tagNameList []string) (err error) {
	for _, name := range tagNameList {
		if err = CreateArticleTagByName(articleID, name); err != nil {
			return
		}
	}
	return
}

// CreateArticleTagByName 通过标签name创建文章标签
func CreateArticleTagByName(articleID int, tagName string) error {
	tagID, err := service.Tag.GetIdByNameOrCreate(tagName)
	if err != nil {
		return err
	}
	return service.Mapping.CreateArticleTag(articleID, tagID)
}

// DeleteArticle 删除文章
func DeleteArticle(id int) error {
	if err := service.Article.Delete(id); err != nil {
		return err
	}
	return service.Mapping.DeleteArticle(id)
}

// BatchDeleteArticle 批量删除文章
func BatchDeleteArticle(ids []int) (err error) {
	for _, id := range ids {
		if err = DeleteArticle(id); err != nil {
			return
		}
	}
	return
}

// DeleteArticleTagByName 通过tagName删除文章标签
func DeleteArticleTagByName(articleID int, tagName string) error {
	tagID, err := service.Tag.GetIdByName(tagName)
	if err != nil {
		return err
	}
	return service.Mapping.DeleteArticleTag(articleID, tagID)
}

// DeleteArticleTagByID 通过tagID删除文章标签
func DeleteArticleTagByID(articleID, tagID int) error {
	return service.Mapping.DeleteArticleTag(articleID, tagID)
}

// DeleteArticleTagByIds 通过tagID列表删除文章标签
func DeleteArticleTagByIds(articleID int, tagIds []int) error {
	return service.Mapping.DeleteArticleTagByTagIds(articleID, tagIds)
}
