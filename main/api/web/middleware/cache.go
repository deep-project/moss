package middleware

import (
	"github.com/gofiber/fiber/v2"
	"moss/domain/config"
	"moss/infrastructure/support/cache"
	"moss/infrastructure/support/log"
)

func Cache(ctx *fiber.Ctx) error {

	if !config.Config.Cache.Enable || ctx.Method() != "GET" {
		return ctx.Next()
	}

	name := ctx.Route().Name
	key := ctx.Path()
	option := config.Config.Cache.GetOption(name)

	if option == nil || !option.Enable {
		return ctx.Next()
	}

	if key == "" || key == "/" {
		key = "default"
	}

	// 默认不打印错误，否则找不到文件错误会爆满
	if val, err := cache.Get(name, key); err == nil {
		return ctx.Type("html").Send(val)
	}

	next := ctx.Next()

	if ctx.Response().StatusCode() == 200 {
		go func() {
			if err := cache.Set(name, key, ctx.Response().Body(), option.TTL.Duration()); err != nil {
				log.Warn("set cache error", log.Err(err))
			}
		}()
	}

	return next
}
