package controller

import (
	"github.com/gofiber/fiber/v2"
	"moss/api/web/mapper"
	appDTO "moss/application/dto"
	appService "moss/application/service"
	"moss/domain/core/entity"
	"moss/domain/core/service"
)

func ArticleList(ctx *fiber.Ctx) error {
	repoCtx, err := mapper.BodyToContext(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.Article.List(&repoCtx)))
}

func ArticleCount(ctx *fiber.Ctx) error {
	where, err := mapper.BodyToWhere(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.Article.CountByWhere(&where)))
}

func ArticleGet(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.Article.Get(id)))
}

func ArticleCreate(ctx *fiber.Ctx) error {
	obj, err := mapper.BodyToCurdModel[appDTO.ArticleCreate](ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(obj, appService.CreateArticle(obj)))
}

func ArticleUpdate(ctx *fiber.Ctx) error {
	obj, err := mapper.BodyToCurdModel[entity.Article](ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(service.Article.Update(obj)))
}

func ArticleDelete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(appService.DeleteArticle(id)))
}

func ArticleBatchDelete(ctx *fiber.Ctx) error {
	ids, err := mapper.BodyToIntSet(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(appService.BatchDeleteArticle(ids)))
}

func ArticleExistsSlug(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.MessageResultData(service.Article.ExistsSlug(string(ctx.Body()))))
}

func ArticleExistsTitle(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.MessageResultData(service.Article.ExistsTitle(string(ctx.Body()))))
}

func ArticleGetTags(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.Tag.ListByArticleID(nil, id)))
}

// ArticleCreateTag 创建文章标签
func ArticleCreateTag(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(appService.CreateArticleTagByName(id, string(ctx.Body()))))
}

// ArticleCreateTagByNameList 创建文章标签通过name列表
func ArticleCreateTagByNameList(ctx *fiber.Ctx) error {
	tagNameList, err := mapper.BodyToStrSet(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(appService.CreateArticleTagsByNameList(id, tagNameList)))
}

// ArticleDeleteTagByName 删除文章标签
func ArticleDeleteTagByName(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(appService.DeleteArticleTagByName(id, string(ctx.Body()))))
}

// ArticleDeleteTagByIds 删除文章标签
func ArticleDeleteTagByIds(ctx *fiber.Ctx) error {
	tagIds, err := mapper.BodyToIntSet(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(appService.DeleteArticleTagByIds(id, tagIds)))
}
