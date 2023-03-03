package controller

import (
	"github.com/gofiber/fiber/v2"
	"moss/api/web/mapper"
	appService "moss/application/service"
	"moss/domain/core/entity"
	"moss/domain/core/service"
)

func CategoryList(ctx *fiber.Ctx) error {
	repoCtx, err := mapper.BodyToContext(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.Category.List(&repoCtx)))
}

func CategoryCount(ctx *fiber.Ctx) error {
	where, err := mapper.BodyToWhere(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.Category.CountByWhere(&where)))
}

func CategoryGet(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.Category.Get(id)))
}

func CategoryCreate(ctx *fiber.Ctx) error {
	obj, err := mapper.BodyToCurdModel[entity.Category](ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(obj, service.Category.Create(obj)))
}

func CategoryUpdate(ctx *fiber.Ctx) error {
	obj, err := mapper.BodyToCurdModel[entity.Category](ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(service.Category.Update(obj)))
}

func CategoryDelete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(service.Category.Delete(id)))
}

func CategoryBatchDelete(ctx *fiber.Ctx) error {
	ids, err := mapper.BodyToIntSet(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(service.Category.BatchDelete(ids)))
}

func CategoryExistsSlug(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.MessageResultData(service.Category.ExistsSlug(string(ctx.Body()))))
}

func CategoryExistsName(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.MessageResultData(service.Category.ExistsName(string(ctx.Body()))))
}

func CategoryTree(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.MessageResultData(appService.CategoryTree()))
}

func CategoryBatchSetParentCategory(ctx *fiber.Ctx) error {
	ids, err := mapper.BodyToIntSet(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	parentID, err := ctx.ParamsInt("parent_id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(service.Category.BatchSetParentCategory(parentID, ids)))
}
