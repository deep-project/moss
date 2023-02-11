package controller

import (
	"github.com/gofiber/fiber/v2"
	"moss/api/web/mapper"
	"moss/domain/core/entity"
	"moss/domain/core/service"
)

func LinkList(ctx *fiber.Ctx) error {
	repoCtx, err := mapper.BodyToContext(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.Link.List(&repoCtx)))
}

func LinkCount(ctx *fiber.Ctx) error {
	where, err := mapper.BodyToWhere(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.Link.CountByWhere(&where)))
}

func LinkGet(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(service.Link.Get(id)))
}

func LinkCreate(ctx *fiber.Ctx) error {
	obj, err := mapper.BodyToCurdModel[entity.Link](ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(obj, service.Link.Create(obj)))
}

func LinkUpdate(ctx *fiber.Ctx) error {
	obj, err := mapper.BodyToCurdModel[entity.Link](ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(service.Link.Update(obj)))
}

func LinkDelete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(service.Link.Delete(id)))
}

func LinkBatchDelete(ctx *fiber.Ctx) error {
	ids, err := mapper.BodyToIntSet(ctx.Body())
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(service.Link.BatchDelete(ids)))
}

func LinkExistsURL(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.MessageResultData(service.Link.ExistsURL(string(ctx.Body()))))
}

func LinkLikeURL(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.MessageResultData(service.Link.ListLikeURL(nil, string(ctx.Body()))))
}

func LinkStatus(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	var item entity.Link
	err = ctx.BodyParser(&item)
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	if item.Status {
		err = service.Link.EnableLink(id)
	} else {
		err = service.Link.DisableLink(id)
	}
	return ctx.JSON(mapper.MessageResult(err))
}
