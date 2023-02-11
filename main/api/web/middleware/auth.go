package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"moss/api/web/dto"
)

func Auth(attrName string, predicate func(token string) (roleName string, ok bool)) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {

		if attrName == "" {
			return errors.New("attrName undefined")
		}

		token := ctx.Get(attrName) // header

		if token == "" {
			token = ctx.Get("Sec-WebSocket-Protocol") // 兼容 websocket
		} else if token == "" {
			token = ctx.Query(attrName)
		}

		if token == "" {
			return ctx.Status(401).JSON(&dto.MessageResult{Message: "authorization failed"})
		}

		roleName, ok := predicate(token)
		if !ok {
			return ctx.Status(401).JSON(&dto.MessageResult{Message: "authorization failed"})
		}

		ctx.Locals("roleName", roleName)

		return ctx.Next()
	}
}
