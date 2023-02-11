package controller

import (
	"github.com/gofiber/fiber/v2"
	"moss/api/web/mapper"
	appService "moss/application/service"
	"moss/domain/support/service"
	"strconv"
)

func PluginList(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.MessageResultData(appService.PluginList(), nil))
}

func PluginOptions(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.MessageResultData(service.Plugin.GetOptions(ctx.Params("id"))))
}

func PluginSaveOptions(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.MessageResult(service.Plugin.UpdateOptions(ctx.Params("id"), ctx.Body())))
}

func PluginRun(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.MessageResult(service.Plugin.Run(ctx.Params("id"))))
}

func PluginCronStart(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.MessageResult(service.Plugin.UpdateCronStart(ctx.Params("id"), true)))
}

func PluginCronStop(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.MessageResult(service.Plugin.UpdateCronStart(ctx.Params("id"), false)))
}

func PluginUpdateCronExp(ctx *fiber.Ctx) error {
	return ctx.JSON(mapper.MessageResult(service.Plugin.UpdateCronExp(ctx.Params("id"), string(ctx.Body()))))
}

func PluginLogList(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "100"))
	return ctx.JSON(mapper.MessageResultData(appService.PluginLogList(ctx.Params("id"), page, limit)))
}
