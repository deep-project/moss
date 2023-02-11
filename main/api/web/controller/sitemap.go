package controller

import (
	"github.com/gofiber/fiber/v2"
	appService "moss/application/service"
)

var Sitemap = new(SitemapController)

type SitemapController struct{}

func (s SitemapController) ArticleTXT(ctx *fiber.Ctx) error {
	res, err := appService.Sitemap.ArticleText()
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.SendString(res)
}

func (s SitemapController) CategoryTXT(ctx *fiber.Ctx) error {
	res, err := appService.Sitemap.CategoryText()
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.SendString(res)
}

func (s SitemapController) TagTXT(ctx *fiber.Ctx) error {
	res, err := appService.Sitemap.TagText()
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.SendString(res)
}

func (s SitemapController) ArticleXML(ctx *fiber.Ctx) error {
	res, err := appService.Sitemap.ArticleXML()
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.Type("xml").SendString(res)
}

func (s SitemapController) CategoryXML(ctx *fiber.Ctx) error {
	res, err := appService.Sitemap.CategoryXML()
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.Type("xml").SendString(res)
}

func (s SitemapController) TagXML(ctx *fiber.Ctx) error {
	res, err := appService.Sitemap.TagXML()
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}
	return ctx.Type("xml").SendString(res)
}
