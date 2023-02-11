package router

import (
	"moss/api/web/middleware"
	"moss/domain/config"
	"moss/infrastructure/general/conf"
	"moss/infrastructure/support/log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Router struct {
	app *fiber.App
}

func New() *Router {
	r := &Router{}
	r.app = r.newFiber()
	return r
}

func (r *Router) newFiber() *fiber.App {
	app := fiber.New(config.Config.Router.GetOptions())

	// http://127.0.0.1:8989/debug/pprof
	// go tool pprof -http=:9090 http://127.0.0.1:8989/debug/pprof/heap?auth=XPAwGQzAwo
	// go tool pprof -http=:9090 http://127.0.0.1:8989/debug/pprof/goroutine
	app.All("/debug/pprof/*", middleware.Pprof, pprof.New())

	// 捕捉堆栈错误
	app.Use(recover.New())
	app.Use(middleware.CatchPanicError)
	// http log
	app.Use(middleware.HttpLog)
	// ETag
	if config.Config.Router.ETag {
		app.Use(etag.New())
	}
	// 压缩
	app.Use(compress.New(compress.Config{Level: compress.Level(config.Config.Router.CompressLevel)}))

	// admin
	app.Route(config.Config.Router.GetAdminPath(), r.RegisterAdmin)

	// home
	app.Route("/", r.RegisterHome)

	return app
}

func (r *Router) Run() error {
	log.Info("app starting...")
	return r.app.Listen(conf.Addr)
}

func (r *Router) Reload() {
	r.app.Server().Handler = r.newFiber().Handler()
}
