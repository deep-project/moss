package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"moss/infrastructure/support/log"
	"runtime"
)

func CatchPanicError(ctx *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			var (
				ok  bool
				err error
				buf = make([]byte, 10240)
				url = ctx.BaseURL() + ctx.Path()
			)
			if err, ok = r.(error); !ok {
				err = fmt.Errorf("%v", r)
			}
			buf = buf[:runtime.Stack(buf, false)]
			log.DPanic("Panic Error!",
				log.String("url", url),
				log.String("error", err.Error()),
				log.String("buf", string(buf)),
			)
			_ = ctx.Status(500).SendString("Internal Server Error")
		}
	}()
	return ctx.Next()
}
