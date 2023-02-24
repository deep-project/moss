package controller

import (
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"moss/domain/config"
	"strings"
)

func AssetsRobotsTxt(ctx *fiber.Ctx) error {
	if config.Config.Template.RobotsTxt == "" {
		return ctx.Next()
	}
	return ctx.SendString(config.Config.Template.RobotsTxt)
}

func AssetsAdsTxt(ctx *fiber.Ctx) error {
	if config.Config.Template.AdsTxt == "" {
		return ctx.Next()
	}
	return ctx.SendString(config.Config.Template.AdsTxt)
}

func FaviconIco(ctx *fiber.Ctx) error {
	if config.Config.Template.FaviconIco == "" {
		return ctx.Next()
	}
	var bs64 = config.Config.Template.FaviconIco
	i := strings.Index(config.Config.Template.FaviconIco, ",")
	if i > 0 {
		bs64 = config.Config.Template.FaviconIco[i+1:]
	}
	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(bs64))
	return ctx.Type("ico").SendStream(dec)
}

func Logo(ctx *fiber.Ctx) error {
	if config.Config.Template.Logo == "" {
		return ctx.Next()
	}
	var bs64 = config.Config.Template.Logo
	i := strings.Index(config.Config.Template.Logo, ",")
	if i > 0 {
		bs64 = config.Config.Template.Logo[i+1:]
	}
	dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(bs64))
	return ctx.Type("png").SendStream(dec)
}
