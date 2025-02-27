package plugins

import (
	"bytes"
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

type PushToBing struct {
	EnableOnCreate bool   `json:"enable_on_create"` // 创建时执行
	EnableOnUpdate bool   `json:"enable_on_update"` // 更新时执行
	ApiKey         string `json:"api_key"`          // api地址

	ctx *pluginEntity.Plugin
}

func NewPushToBing() *PushToBing {
	return &PushToBing{}
}

func (p *PushToBing) Info() *pluginEntity.PluginInfo {
	return &pluginEntity.PluginInfo{
		ID:    "PushToBing",
		About: "push article to bing when created or updated",
	}
}
func (p *PushToBing) Run(ctx *pluginEntity.Plugin) error {
	return nil
}

func (p *PushToBing) Load(ctx *pluginEntity.Plugin) error {
	p.ctx = ctx
	service.Article.AddCreateAfterEvents(p)
	service.Article.AddUpdateAfterEvents(p)
	return nil
}
func (p *PushToBing) ArticleCreateAfter(item *entity.Article) {
	if p.EnableOnCreate {
		p.pushArticle(item, "create")
	}
}

func (p *PushToBing) ArticleUpdateAfter(item *entity.Article) {
	if p.EnableOnUpdate {
		p.pushArticle(item, "update")
	}
}

func (p *PushToBing) pushArticle(item *entity.Article, action string) {
	p.push("article", item.FullURL(), zap.Any("article",
		map[string]interface{}{"id": item.ID, "title": item.Title}),
		zap.String("action", action))
}

func (p *PushToBing) push(title, uri string, logs ...zap.Field) {
	res, err := p.PushURL(uri)
	if err != nil {
		p.ctx.Log.Error(title+" push error!", zap.Error(err))
		return
	}
	var logAll = append([]zap.Field{zap.String("url", uri), zap.Any("result", res)}, logs...)
	if res != 200 {
		p.ctx.Log.Error(title+" push error! code: ", append(logAll, zap.Error(err))...)
		return
	}
	p.ctx.Log.Info("article push success.", logAll...)
}

// PushURL 推送url
func (p *PushToBing) PushURL(uri ...string) (int, error) {
	if config.Config.Site.URL == "" {
		return -1, errors.New("site url undefined")
	}
	if p.ApiKey == "" {
		return -2, errors.New("api key undefined")
	}
	if len(uri) == 0 {
		return -3, errors.New("uri is required")
	}
	var push pushToBingReq
	push.Host = strings.ReplaceAll(strings.ReplaceAll(
		config.Config.Site.URL, "https://", ""), "http://", "")
	push.Key = p.ApiKey
	push.KeyLocation = config.Config.Site.URL + "/" + p.ApiKey + ".txt"
	push.URLList = uri
	reqBody, err := json.Marshal(push)
	if err != nil {
		return -4, err
	}
	r := request.New()
	r.AddHeader("Content-Type", "application/json; charset=utf-8")
	body, err := r.Post("https://api.indexnow.org/IndexNow", bytes.NewReader(reqBody))
	if err != nil {
		return -5, err
	}
	return body.StatusCode, nil
}

// pushToBingReq 必应提交
type pushToBingReq struct {
	Host        string   `json:"host"`
	Key         string   `json:"key"`
	KeyLocation string   `json:"keyLocation"`
	URLList     []string `json:"urlList"`
}
