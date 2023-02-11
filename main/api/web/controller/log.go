package controller

import (
	"github.com/gofiber/fiber/v2"
	"moss/api/web/mapper"
	"moss/application/service"
	"moss/infrastructure/support/log"
	"strconv"
)

func LogInit(ctx *fiber.Ctx) error {
	log.Init()
	return ctx.JSON(mapper.MessageResult(nil))
}

func LogRead(ctx *fiber.Ctx) error {
	var page, _ = strconv.Atoi(ctx.Query("page"))
	var limit, _ = strconv.Atoi(ctx.Query("limit"))
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}
	return ctx.JSON(mapper.MessageResultData(service.LogRead(ctx.Params("id"), page, limit)))
}
