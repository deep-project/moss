package middleware

import (
	"strings"

	"github.com/duke-git/lancet/v2/slice"
	"github.com/gofiber/fiber/v2"
)

// ReplaceBodyContent replace body content
func ReplaceBodyContent(list map[string]string, mustContentTypes []string) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		next := ctx.Next()

		if list == nil {
			return next
		}

		if len(mustContentTypes) > 0 {
			bodyType := string(ctx.Response().Header.ContentType())
			_, ok := slice.Find(mustContentTypes, func(_ int, val string) bool {
				return strings.Contains(bodyType, val)
			})
			if !ok {
				return next
			}
		}

		body := string(ctx.Response().Body())
		for src, dst := range list {
			body = strings.ReplaceAll(body, src, dst)
		}
		ctx.Context().SetBody([]byte(body))

		return next
	}
}
