package plugins

import (
	"github.com/microcosm-cc/bluemonday"
	"moss/domain/core/entity"
	"moss/domain/core/service"
	pluginEntity "moss/domain/support/entity"
)

type ArticleSanitizer struct {
	EnableOnCreate         bool `json:"enable_on_create"`           // 创建时执行
	EnableOnUpdate         bool `json:"enable_on_update"`           // 更新时执行
	AllowRelativeURLs      bool `json:"allow_relative_urls"`        // 禁止本地URL
	RequireNoFollowOnLinks bool `json:"require_no_follow_on_links"` // 所有a标签 都添加 rel="nofollow"

	ctx *pluginEntity.Plugin
}

func NewArticleSanitizer() *ArticleSanitizer {
	return &ArticleSanitizer{
		EnableOnCreate:         true,
		EnableOnUpdate:         false,
		AllowRelativeURLs:      true,
		RequireNoFollowOnLinks: true,
	}
}

func (a *ArticleSanitizer) Info() *pluginEntity.PluginInfo {
	return &pluginEntity.PluginInfo{
		ID:    "ArticleSanitizer",
		About: "to scrub content of XSS when created or updated",
	}
}

func (a *ArticleSanitizer) Load(ctx *pluginEntity.Plugin) error {
	a.ctx = ctx
	service.Article.AddCreateBeforeEvents(a)
	service.Article.AddUpdateBeforeEvents(a)
	return nil
}

func (a *ArticleSanitizer) ArticleCreateBefore(item *entity.Article) error {
	if a.EnableOnCreate {
		return a.sanitize(item, "create")
	}
	return nil
}

func (a *ArticleSanitizer) ArticleUpdateBefore(item *entity.Article) error {
	if a.EnableOnUpdate {
		return a.sanitize(item, "update")
	}
	return nil
}

func (a *ArticleSanitizer) sanitize(item *entity.Article, action string) error {
	p := bluemonday.UGCPolicy()
	p.AllowDataURIImages()                             // 验证base64图片的合法性
	p.RequireParseableURLs(true)                       // 过滤非法url
	p.AddTargetBlankToFullyQualifiedLinks(true)        // a标签增加 _blank
	p.RequireNoFollowOnLinks(a.RequireNoFollowOnLinks) // 所有a标签 都添加 rel="nofollow"
	p.AllowRelativeURLs(a.AllowRelativeURLs)           // 禁止本地url
	p.AllowURLSchemes("mailto", "http", "https")       // 指定url协议头
	item.Content = p.Sanitize(item.Content)
	//a.ctx.Log.Info(fmt.Sprintf("%s sanitize success", item.Title))
	return nil
}

func (a *ArticleSanitizer) Run(ctx *pluginEntity.Plugin) (err error) {
	return nil
}
