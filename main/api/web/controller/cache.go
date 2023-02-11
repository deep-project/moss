package controller

import (
	"github.com/gofiber/fiber/v2"
	"moss/api/web/mapper"
	"moss/domain/config"
	"moss/infrastructure/support/cache"
)

func CacheInit(ctx *fiber.Ctx) error {
	if err := cache.Init(); err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(nil))
}

func CacheClear(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	// 执行之前，先查询一下是否存在设置的前缀，防止提交其他目录字符串导致安全问题
	if opt := config.Config.Cache.GetOption(name); opt == nil {
		return ctx.JSON(mapper.MessageFail("option not found"))
	}
	err := cache.ClearBucket(name)
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(nil))
}
