package template

import (
	"moss/domain/config"
	"moss/domain/config/entity"
	"moss/infrastructure/support/template/query"
	"moss/infrastructure/support/template/utils"
	"moss/infrastructure/support/template/widget"
)

func initEngineFn(e Engine) {
	e.AddFunc("Config", newConfigHandle())
	e.AddFunc("Widget", widget.New())
	e.AddFunc("Utils", utils.New())

	e.AddFunc("Article", query.NewArticle)
	e.AddFunc("Category", query.NewCategory)
	e.AddFunc("Tag", query.NewTag)
	e.AddFunc("Link", query.NewLink)
}

type configHandle struct {
	Site     *entity.Site
	Router   *entity.Router
	Upload   *entity.Upload
	Cache    *entity.Cache
	Theme    *entity.Theme
	Template *entity.Template
	Sitemap  *entity.Sitemap
}

func newConfigHandle() *configHandle {
	return &configHandle{
		Site:     config.Config.Site,
		Router:   config.Config.Router,
		Upload:   config.Config.Upload,
		Cache:    config.Config.Cache,
		Theme:    config.Config.Theme,
		Template: config.Config.Template,
		Sitemap:  config.Config.Sitemap,
	}
}
