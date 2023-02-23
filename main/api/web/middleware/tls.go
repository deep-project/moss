package middleware

import (
	"github.com/gofiber/fiber/v2"
	"moss/domain/config"
	"strconv"
	"strings"
)

func TLS(ctx *fiber.Ctx) error {

	if !config.Config.TLS.Enable {
		return ctx.Next()
	}

	// 强制http跳转到https
	if config.Config.TLS.ForceHTTPS && ctx.Protocol() == "http" {
		domain := strings.Split(ctx.Hostname(), ":")[0]
		port := config.Config.TLS.ListenPort()
		var portStr string
		if port != 443 {
			portStr = ":" + strconv.Itoa(port)
		}
		tlsURL := "https://" + domain + portStr + ctx.OriginalURL()
		return ctx.Redirect(tlsURL, config.Config.TLS.GetRedirectStatus())
	}

	return ctx.Next()

}
