package controller

import (
	"github.com/gofiber/fiber/v2"
	"moss/api/web/mapper"
	appService "moss/application/service"
	"moss/domain/core/entity"
	"moss/domain/core/service"
)

func TagList(ctx *fiber.Ctx) error {
	repoCtx, err := mapper.BodyToContext(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.Tag.List(&repoCtx)))
}

func TagCount(ctx *fiber.Ctx) error {
	where, err := mapper.BodyToWhere(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.Tag.CountByWhere(&where)))
}

func TagGet(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.Tag.Get(id)))
}

func TagCreate(ctx *fiber.Ctx) error {
	obj, err := mapper.BodyToCurdModel[entity.Tag](ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(obj, service.Tag.Create(obj)))
}

func TagUpdate(ctx *fiber.Ctx) error {
	obj, err := mapper.BodyToCurdModel[entity.Tag](ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(service.Tag.Update(obj)))
}

func TagDelete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(appService.DeleteTag(id)))
}

func TagBatchDelete(ctx *fiber.Ctx) error {
	ids, err := mapper.BodyToIntSet(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(appService.BatchDeleteTag(ids)))
}

func TagExistsSlug(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.MessageResultData(service.Tag.ExistsSlug(string(ctx.Body()))))
}

func TagExistsName(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.MessageResultData(service.Tag.ExistsName(string(ctx.Body()))))
}

func TagListByArticleID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.Tag.ListByArticleID(nil, id)))
}

func TagGetByIds(ctx *fiber.Ctx) error {
	ids, err := mapper.BodyToIntSet(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.Tag.ListByIds(nil, ids)))
}
