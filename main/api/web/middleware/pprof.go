package middleware

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
	"moss/domain/config"
	"strings"
)

func Pprof(ctx *fiber.Ctx) error {

	ctx.Set("Expires", "0")
	ctx.Set("Pragma", "No-Cache")
	ctx.Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")

	if ctx.Query("auth") != config.Config.Router.PprofSecret {
		return ctx.Status(404).SendString("auth failed")
	}

	next := ctx.Next()

	if config.Config.Router.PprofSecret == "" {
		return next
	}

	if strings.Contains(string(ctx.Response().Header.ContentType()), "text/html") {
		body := string(ctx.Response().Body())
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(body))
		if err != nil {
			return ctx.Next()
		}
		dom.Find("a").Each(func(i int, s *goquery.Selection) {
			if href, ok := s.Attr("href"); ok {
				if strings.Contains(href, "?") {
					s.SetAttr("href", href+"&auth="+config.Config.Router.PprofSecret)
				} else {
					s.SetAttr("href", href+"?auth="+config.Config.Router.PprofSecret)
				}
			}
		})
		if html, err := dom.Html(); err == nil {
			ctx.Context().SetBody([]byte(html))
		}
	}
	return next
}
