package controller

import (
	"moss/api/web/mapper"
	appService "moss/application/service"
	"moss/domain/core/entity"
	"moss/domain/core/service"

	"github.com/gofiber/fiber/v2"
)

func StoreList(ctx *fiber.Ctx) error {
	repoCtx, err := mapper.BodyToContext(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.Store.List(&repoCtx)))
}

func StoreCount(ctx *fiber.Ctx) error {
	where, err := mapper.BodyToWhere(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.Store.CountByWhere(&where)))
}

func StoreGet(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.Store.Get(id)))
}

func StoreCreate(ctx *fiber.Ctx) error {
	var obj entity.Store
	if err := ctx.BodyParser(&obj); err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(obj, service.Store.Create(&obj)))
}

func StoreUpdate(ctx *fiber.Ctx) error {
	var obj entity.Store
	if err := ctx.BodyParser(&obj); err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(service.Store.Update(&obj)))
}

func StoreDelete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(service.Store.Delete(id)))
}

func StoreBatchDelete(ctx *fiber.Ctx) error {
	ids, err := mapper.BodyToIntSet(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(service.Store.BatchDelete(ids)))
}

func StorePost(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(appService.StorePost(id)))
}
