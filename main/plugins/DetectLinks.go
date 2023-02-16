package plugins

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"go.uber.org/zap"
	"moss/domain/config"
	"moss/domain/core/entity"
	"moss/domain/core/service"
	pluginEntity "moss/domain/support/entity"
	"moss/infrastructure/utils/request"
	"strings"
)

type DetectLinks struct {
	Referer   string `json:"referer"`
	Proxy     string `json:"proxy"`
	Retry     int    `json:"retry"`
	UserAgent string `json:"user_agent"` // ua头
	Timeout   int    `json:"timeout"`    // 超时，秒

	ctx       *pluginEntity.Plugin
	mySiteURL string
}

func NewDetectLinks() *DetectLinks {
	return &DetectLinks{
		Retry:   2,
		Timeout: 30,
	}
}

func (d *DetectLinks) Info() *pluginEntity.PluginInfo {
	return &pluginEntity.PluginInfo{
		ID:         "DetectLinks",
		About:      "detect links contains my link",
		RunEnable:  true,
		CronEnable: true,
		PluginInfoPersistent: pluginEntity.PluginInfoPersistent{
			CronStart: true,
			CronExp:   "@every 5h",
		},
	}
}

func (d *DetectLinks) Load(ctx *pluginEntity.Plugin) error {
	return nil
}

func (d *DetectLinks) Run(ctx *pluginEntity.Plugin) (err error) {

	d.ctx = ctx
	d.mySiteURL = config.Config.Site.URL
	d.mySiteURL = strings.TrimPrefix(d.mySiteURL, "http://")
	d.mySiteURL = strings.TrimPrefix(d.mySiteURL, "https://")
	d.mySiteURL = strings.TrimPrefix(d.mySiteURL, "//")
	d.mySiteURL = strings.TrimSuffix(d.mySiteURL, "/")

	if d.mySiteURL == "" {
		ctx.Log.Error("my site url is empty")
		return errors.New("my site url is empty")
	}

	list, err := service.Link.ListDetectLink(nil)
	if err != nil {
		ctx.Log.Error("query list error", zap.Error(err))
		return errors.New("query list error")
	}

	ctx.Log.Info("begin...")
	for _, item := range list {
		d.run(&item)
	}
	ctx.Log.Info("end.")
	return nil
}

func (d *DetectLinks) log(item *entity.Link, err error) []zap.Field {
	return []zap.Field{zap.String("name", item.Name), zap.String("url", item.URL), zap.Error(err)}
}

func (d *DetectLinks) run(item *entity.Link) {

	if item.URL == "" || !strings.HasPrefix(item.URL, "http") {
		d.ctx.Log.Warn("url is wrong", d.log(item, nil)...)
		return
	}
	// 抓取html
	body, err := request.New().SetRetry(d.Retry).SetProxyURLStr(d.Proxy).SetTimeoutSeconds(d.Timeout).SetReferer(d.Referer).SetUserAgentMust(d.UserAgent).GetBody(item.URL)
	if err != nil {
		d.disableLink(item) // 访问出错，直接下链
		d.ctx.Log.Warn("get url error", d.log(item, err)...)
		return
	}
	// 格式化
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		d.disableLink(item) // 格式化出错，直接下链
		d.ctx.Log.Error("format html document error", d.log(item, err)...)
		return
	}
	// 查找link
	var isFound bool
	doc.Find("a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		href, _ := s.Attr("href")
		if href == "" {
			return true
		}
		if !strings.HasPrefix(href, "http://"+d.mySiteURL) && !strings.HasPrefix(href, "https://"+d.mySiteURL) {
			return true // 未找到
		}
		// 已找到
		rel, _ := s.Attr("rel") // 检查 rel
		if rel != "" {
			d.ctx.Log.Warn("Link used ref attribute!", append(d.log(item, nil), zap.String("rel", rel))...)
			return false // 退出循环
		}
		isFound = true // 确定找到
		return false   // 退出循环
	})
	if isFound {
		d.enableLink(item)
	} else {
		d.disableLink(item)
	}
}

func (d *DetectLinks) enableLink(item *entity.Link) {
	if item.Status {
		d.ctx.Log.Info("no change required", d.log(item, nil)...)
		return
	}
	if err := service.Link.EnableLink(item.ID); err != nil {
		d.ctx.Log.Error("update error", d.log(item, err)...)
		return
	}
	d.ctx.Log.Info("enable link success", d.log(item, nil)...)
}

func (d *DetectLinks) disableLink(item *entity.Link) {
	if !item.Status {
		d.ctx.Log.Info("no change required", d.log(item, nil)...)
		return
	}
	if err := service.Link.DisableLink(item.ID); err != nil {
		d.ctx.Log.Error("update error", d.log(item, err)...)
		return
	}
	d.ctx.Log.Info("disable link success", d.log(item, nil)...)
}
