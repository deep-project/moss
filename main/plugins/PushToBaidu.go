package plugins

import (
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"moss/domain/config"
	"moss/domain/core/entity"
	"moss/domain/core/service"
	pluginEntity "moss/domain/support/entity"
	"moss/infrastructure/utils/request"
	"strings"
)

type PushToBaidu struct {
	EnableOnCreate bool   `json:"enable_on_create"` // 创建时执行
	EnableOnUpdate bool   `json:"enable_on_update"` // 更新时执行
	ApiURL         string `json:"api_url"`          // api地址

	ctx *pluginEntity.Plugin
}

func NewPushToBaidu() *PushToBaidu {
	return &PushToBaidu{}
}

func (p *PushToBaidu) Info() *pluginEntity.PluginInfo {
	return &pluginEntity.PluginInfo{
		ID:    "PushToBaidu",
		About: "push article to baidu when created or updated",
	}
}
func (p *PushToBaidu) Run(ctx *pluginEntity.Plugin) error { return nil }

func (p *PushToBaidu) Load(ctx *pluginEntity.Plugin) error {
	p.ctx = ctx
	service.Article.AddCreateAfterEvents(p)
	service.Article.AddUpdateAfterEvents(p)
	return nil
}
func (p *PushToBaidu) ArticleCreateAfter(item *entity.Article) {
	if p.EnableOnCreate {
		p.pushArticle(item, "create")
	}
}

func (p *PushToBaidu) ArticleUpdateAfter(item *entity.Article) {
	if p.EnableOnUpdate {
		p.pushArticle(item, "update")
	}
}

func (p *PushToBaidu) pushArticle(item *entity.Article, action string) {
	p.push("article", item.FullURL(), zap.Any("article", map[string]interface{}{"id": item.ID, "title": item.Title}), zap.String("action", action))
}

func (p *PushToBaidu) push(title, uri string, logs ...zap.Field) {
	res, err := p.PushURL(uri)
	if err != nil {
		p.ctx.Log.Error(title+" push error!", zap.Error(err))
		return
	}
	var logAll = append([]zap.Field{zap.String("url", uri), zap.Any("result", res)}, logs...)
	if res.Success == 0 {
		p.ctx.Log.Error(title+" push error!", append(logAll, zap.Error(err))...)
		return
	}
	p.ctx.Log.Info("article push success.", logAll...)
}

// PushURL 推送url
func (p *PushToBaidu) PushURL(uri ...string) (*PushToBaiduResult, error) {
	if config.Config.Site.URL == "" {
		return nil, errors.New("site url undefined")
	}
	if p.ApiURL == "" {
		return nil, errors.New("api url undefined")
	}
	if len(uri) == 0 {
		return nil, errors.New("uri is required")
	}
	var val = strings.Join(uri, "\n")
	body, err := request.New().PostReturnBody(p.ApiURL, strings.NewReader(val))
	if err != nil {
		return nil, err
	}
	var res PushToBaiduResult
	if err = json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// PushToBaiduResult 百度提交结果
type PushToBaiduResult struct {
	Success     int      `json:"success"`       // 成功推送的url条数
	Remain      int      `json:"remain"`        // 当天剩余的可推送url条数
	NotSameSite []string `json:"not_same_site"` // 由于不是本站url而未处理的url列表
	NotValid    []string `json:"not_valid"`     // 不合法的url列表
	Error       int      `json:"error"`         // 错误码，与状态码相同
	Message     string   `json:"message"`       // 错误描述
}
