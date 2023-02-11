package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/panjf2000/ants/v2"
	appService "moss/application/service"
	"moss/domain/config"
	"moss/domain/core/service"
	"moss/infrastructure/general/message"
	"moss/infrastructure/support/log"
	"moss/infrastructure/support/template"
)

func HomeIndex(ctx *fiber.Ctx) error {
	b, err := appService.Render.Index()
	if err != nil {
		log.Error("index controller failed", log.Err(err))
		return ctx.SendStatus(500)
	}
	return ctx.Type("html", "utf-8").SendString(string(b))
}

func HomeCategory(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	page, _ := ctx.ParamsInt("page", 1)
	if page == 0 {
		page = 1
	}
	// 超出最大页数限制
	if config.Config.Template.CategoryPageList.MaxPage > 0 && page > config.Config.Template.CategoryPageList.MaxPage {
		return ctx.SendStatus(404)
	}
	b, err := appService.Render.CategoryBySlug(slug, page)
	if err != nil {
		if errors.Is(err, message.ErrRecordNotFound) {
			return ctx.Next()
		}
		log.Error("category controller failed", log.Err(err))
		return ctx.SendStatus(500)
	}
	return ctx.Type("html", "utf-8").SendString(string(b))
}

func HomeTag(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	page, _ := ctx.ParamsInt("page", 1)
	if page == 0 {
		page = 1
	}
	// 限制最大页数
	if config.Config.Template.TagPageList.MaxPage > 0 && page > config.Config.Template.TagPageList.MaxPage {
		return ctx.SendStatus(404)
	}
	b, err := appService.Render.TagBySlug(slug, page)
	if err != nil {
		if errors.Is(err, message.ErrRecordNotFound) {
			return ctx.Next()
		}
		log.Error("tag controller failed", log.Err(err))
		return ctx.SendStatus(500)
	}

	return ctx.Type("html", "utf-8").SendString(string(b))
}

func HomeArticle(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")
	b, err := appService.Render.ArticleBySlug(slug)
	if err != nil {
		if errors.Is(err, message.ErrRecordNotFound) {
			return ctx.Next()
		}
		log.Error("article controller failed", log.Err(err))
		return ctx.SendStatus(500)
	}
	return ctx.Type("html", "utf-8").SendString(string(b))
}

func HomeArticleViews(ctx *fiber.Ctx) error {
	if !ctx.XHR() {
		return ctx.SendStatus(200)
	}
	if config.Config.More.ArticleViewsPool == 0 || viewsPool.Free() == 0 {
		return ctx.SendStatus(200)
	}
	if err := viewsPool.Invoke(ctx.Params("slug")); err != nil {
		log.Warn("article views put in pool failed", log.Err(err))
	}
	return ctx.SendStatus(200)
}

var viewsPool, _ = ants.NewPoolWithFunc(config.Config.More.ArticleViewsPool, articleViewUpdate)

func articleViewUpdate(val any) {
	slug, ok := val.(string)
	if !ok {
		log.Warn("article slug transform error in views update")
		return
	}
	if err := service.Article.UpdateViewsBySlug(slug, 1); err != nil {
		log.Warn("article views update error", log.Err(err))
	}
}

func HomeNotFound(ctx *fiber.Ctx) error {
	b, err := template.Render("template/notFound.html", template.Binds{
		Page: template.Page{
			Name: "notFound",
			Path: ctx.Path(),
		},
	})
	if err != nil {
		return ctx.SendStatus(404)
	}
	return ctx.Type("html", "utf-8").Status(404).SendString(string(b))
}
