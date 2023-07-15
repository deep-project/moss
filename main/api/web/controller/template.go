package controller

import (
	"moss/api/web/mapper"
	"moss/application/service"
	"moss/infrastructure/support/log"
	"moss/infrastructure/support/template"
	"path"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ThemeInit(ctx *fiber.Ctx) error {
	if err := template.InitTemplate(); err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResult(nil))
}

func ThemeList(ctx *fiber.Ctx) error {
	list, err := template.ThemeList()
	if err != nil {
		return ctx.JSON(mapper.MessageResult(err))
	}
	return ctx.JSON(mapper.MessageResultData(list, nil))
}

func ThemeScreenshot(ctx *fiber.Ctx) error {
	var id = ctx.Params("id")
	if id == "" {
		return ctx.JSON(mapper.MessageFail("id is required"))
	}
	dir := filepath.Base(id)
	return ctx.JSON(mapper.MessageResultData(template.ReadThemeScreenshot(dir), nil))
}

// TemplatePage 模板自定义页面匹配
func TemplatePage(ctx *fiber.Ctx) error {
	var _path = ctx.Path()
	if _path == "" || _path == "/" {
		return ctx.Next()
	}
	if path.Ext(_path) == "" {
		_path = filepath.Join(_path, "index.html")
	}
	b, err := service.Render.TemplatePage(_path)
	if err != nil {
		if strings.HasPrefix(err.Error(), "template not found:") {
			return ctx.Next()
		}
		log.Error("page controller failed", log.Err(err))
		return ctx.SendStatus(500)
	}
	return ctx.Type("html", "utf-8").SendString(string(b))
}
