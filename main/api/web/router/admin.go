package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"moss/api/web/controller"
	"moss/api/web/dto"
	"moss/api/web/middleware"
	"moss/domain/config"
	"moss/resources"
	"net/http"
	"strings"
)

func (r *Router) RegisterAdmin(route fiber.Router) {
	route.Route("/api", r.api, "api")
	route.Route("/", web, "web")
}

func (r *Router) api(route fiber.Router) {

	route.Get("/admin/exists", controller.AdminExists)
	route.Post("/admin/create", controller.AdminCreate)
	route.Get("/admin/captcha", controller.AdminCaptcha)
	route.Post("/admin/login", controller.AdminLogin)

	route.Use(auth())

	// router
	route.Post("/router/reload", r.ReloadRouter)

	// config
	route.Get("/config", controller.ConfigList)
	route.Get("/config/:id", controller.ConfigGet)
	route.Post("/config/:id", controller.ConfigUpdate)

	// article
	route.Post("/article/list", controller.ArticleList)
	route.Post("/article/count", controller.ArticleCount)
	route.Get("/article/get/:id", controller.ArticleGet)
	route.Post("/article/create", controller.ArticleCreate)
	route.Post("/article/update", controller.ArticleUpdate)
	route.Post("/article/delete/:id", controller.ArticleDelete)
	route.Post("/article/batchDelete", controller.ArticleBatchDelete)
	route.Post("/article/existsSlug", controller.ArticleExistsSlug)
	route.Post("/article/existsTitle", controller.ArticleExistsTitle)
	route.Get("/article/getTags/:id", controller.ArticleGetTags)
	route.Post("/article/createTag/:id", controller.ArticleCreateTag)
	route.Post("/article/createTagByNameList/:id", controller.ArticleCreateTagByNameList)
	route.Post("/article/deleteTagByName/:id", controller.ArticleDeleteTagByName)
	route.Post("/article/deleteTagByIds/:id", controller.ArticleDeleteTagByIds)
	route.Post("/article/batchSetCategory/:category_id", controller.ArticleBatchSetCategory)

	// category
	route.Post("/category/list", controller.CategoryList)
	route.Post("/category/count", controller.CategoryCount)
	route.Get("/category/get/:id", controller.CategoryGet)
	route.Post("/category/create", controller.CategoryCreate)
	route.Post("/category/update", controller.CategoryUpdate)
	route.Post("/category/delete/:id", controller.CategoryDelete)
	route.Post("/category/batchDelete", controller.CategoryBatchDelete)
	route.Post("/category/existsSlug", controller.CategoryExistsSlug)
	route.Post("/category/existsName", controller.CategoryExistsName)
	route.Get("/category/tree", controller.CategoryTree)
	route.Post("/category/batchSetParentCategory/:parent_id", controller.CategoryBatchSetParentCategory)

	// tag
	route.Post("/tag/list", controller.TagList)
	route.Post("/tag/count", controller.TagCount)
	route.Get("/tag/get/:id", controller.TagGet)
	route.Post("/tag/create", controller.TagCreate)
	route.Post("/tag/update", controller.TagUpdate)
	route.Post("/tag/delete/:id", controller.TagDelete)
	route.Post("/tag/batchDelete", controller.TagBatchDelete)
	route.Post("/tag/existsSlug", controller.TagExistsSlug)
	route.Post("/tag/existsName", controller.TagExistsName)
	route.Get("/tag/list/article/:id", controller.TagListByArticleID)
	route.Post("/tag/getByIds", controller.TagGetByIds)

	// link
	route.Post("/link/list", controller.LinkList)
	route.Post("/link/count", controller.LinkCount)
	route.Get("/link/get/:id", controller.LinkGet)
	route.Post("/link/create", controller.LinkCreate)
	route.Post("/link/update", controller.LinkUpdate)
	route.Post("/link/delete/:id", controller.LinkDelete)
	route.Post("/link/batchDelete", controller.LinkBatchDelete)
	route.Post("/link/existsURL", controller.LinkExistsURL)
	route.Post("/link/likeURL", controller.LinkLikeURL)
	route.Post("/link/status/:id", controller.LinkStatus)

	// store
	route.Post("/store/list", controller.StoreList)
	route.Post("/store/count", controller.StoreCount)
	route.Get("/store/get/:id", controller.StoreGet)
	route.Post("/store/create", controller.StoreCreate)
	route.Post("/store/update", controller.StoreUpdate)
	route.Post("/store/delete/:id", controller.StoreDelete)
	route.Post("/store/batchDelete", controller.StoreBatchDelete)
	route.Post("/store/post/:id", controller.StorePost)

	// log
	route.Post("/log/init", controller.LogInit)
	route.Get("/log/read/:id", controller.LogRead)

	// theme
	route.Post("/theme/init", controller.ThemeInit)
	route.Get("/theme/list", controller.ThemeList)
	route.Get("/theme/screenshot/:id", controller.ThemeScreenshot)

	// cache
	route.Post("/cache/init", controller.CacheInit)
	route.Post("/cache/clear/:name", controller.CacheClear)

	// upload
	route.Post("/upload/init", controller.UploadInit)
	route.Post("/upload", controller.Upload)

	// plugin
	route.Get("/plugin/list", controller.PluginList)
	route.Get("/plugin/options/:id", controller.PluginOptions)
	route.Post("/plugin/saveOptions/:id", controller.PluginSaveOptions)
	route.Post("/plugin/run/:id", controller.PluginRun)
	route.Post("/plugin/cron/start/:id", controller.PluginCronStart)
	route.Post("/plugin/cron/stop/:id", controller.PluginCronStop)
	route.Post("/plugin/cron/exp/:id", controller.PluginUpdateCronExp)
	route.Get("/plugin/log/list/:id", controller.PluginLogList)

	// dashboard
	route.Get("/dashboard/:id", controller.Dashboard.Controller)

}

func auth() any {
	return middleware.Auth("token", func(token string) (string, bool) {
		if config.Config.Admin.VerifyJwtToken(token) {
			return "administrator", true
		}
		if config.Config.API.Enable && token == config.Config.API.SecretKey {
			return "api", true
		}
		return "", false
	})
}

func web(route fiber.Router) {
	route.Use(middleware.ReplaceBodyContent(
		map[string]string{
			// 注意：静态页面打包时dir带前后斜杠，所以这里去掉斜杠
			"{{__DIR__}}": strings.Trim(config.Config.Router.GetAdminPath(), "/"),
		},
		// pwa文件 manifest.webmanifest 的类型是 application/octet-stream
		[]string{"text/html", "javascript", "text/css", "application/octet-stream"}))

	route.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(resources.Admin()),
		Index:        "index.html",
		NotFoundFile: "index.html",
		// Next 如果访问地址非管理路径，则执行Next
		Next: func(ctx *fiber.Ctx) bool {
			path := ctx.Path()
			dir := config.Config.Router.GetAdminPath()
			if dir == path {
				return false
			}
			if !strings.HasSuffix(dir, "/") {
				dir = dir + "/"
			}
			if strings.HasPrefix(path, dir) {
				return false
			}
			return true
		},
	}))
}

func (r *Router) ReloadRouter(ctx *fiber.Ctx) error {
	r.Reload()
	return ctx.JSON(dto.MessageResult{Success: true})
}
