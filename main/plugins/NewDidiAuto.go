package plugins

import (
	"encoding/json"
	"fmt"
	"moss/domain/core/entity"
	"moss/domain/core/service"
	pluginEntity "moss/domain/support/entity"
	"moss/infrastructure/support/cache"
	"moss/infrastructure/utils/request"
	"time"

	"github.com/bitly/go-simplejson"
	"go.uber.org/zap"
)

const (
	didiWaitPageURL  = "https://didi.seowhy.com/management-3.html?siteId=%d"                           // 等待上链页面
	didiCheckPageURL = "https://didi.seowhy.com/www/management/management_pass?check_id=%d&pagetype=3" // 验证上链页面
)

// DidiAuto 滴滴友链自动任务
type DidiAuto struct {
	SiteID          int    `json:"site_id"`           // 站点ID
	Cookie          string `json:"cookie"`            // cookie
	DetectDelay     int64  `json:"detect_delay"`      // 检测延迟(分钟)
	Retry           int    `json:"retry"`             // 重试次数
	RetrySleep      int    `json:"retry_sleep"`       // 重试等待时间（秒）
	ClearCacheURL   string `json:"clear_cache_url"`   // 清除缓存要触发的URL
	ClearCacheSleep int    `json:"clear_cache_sleep"` // 清除缓存url后等待时间（秒）

	ctx *pluginEntity.Plugin
}

func NewDidiAuto() *DidiAuto {
	return &DidiAuto{DetectDelay: 720, Retry: 3, RetrySleep: 5, ClearCacheSleep: 5}
}

func (d *DidiAuto) Info() *pluginEntity.PluginInfo {
	return &pluginEntity.PluginInfo{
		ID:         "DidiAuto",
		About:      "滴滴友链自动任务",
		RunEnable:  true,
		CronEnable: true,
		PluginInfoPersistent: pluginEntity.PluginInfoPersistent{
			CronStart: false,
			CronExp:   "@every 5m",
		},
	}
}

func (d *DidiAuto) Load(ctx *pluginEntity.Plugin) error {
	return nil
}

func (d *DidiAuto) Run(ctx *pluginEntity.Plugin) (err error) {
	if d.SiteID == 0 {
		ctx.Log.Error("未设置站点ID")
		return
	}
	if d.Cookie == "" {
		ctx.Log.Error("未设置cookie")
		return
	}
	if d.DetectDelay < 0 {
		d.DetectDelay = 0
	}
	d.ctx = ctx
	waitLink, err := d.getWaitLink()
	if err != nil {
		return
	}
	if len(waitLink) == 0 {
		d.ctx.Log.Info("没有待上链的站点")
		return
	}
	for _, item := range waitLink {
		d.run(item)
	}

	// 清除缓存
	if h, _ := cache.ActiveDriver(); h != nil {
		if err := cache.ClearBucket("home"); err != nil {
			d.ctx.Log.Warn("清除首页缓存失败", zap.Error(err))
		}
		time.Sleep(2 * time.Second) // 暂停两秒 给URL触发清除缓存留下时间
	}

	if d.ClearCacheURL != "" {
		body, _ := request.New().GetBody(d.ClearCacheURL)
		d.ctx.Log.Info("触发清除缓存URL结果", zap.String("body", string(body)))
		time.Sleep(time.Duration(d.ClearCacheSleep) * time.Second) // 等待缓存清理成功
	}

	// 验证上链
	for _, item := range waitLink {
		d.checkPass(item, 0)
	}

	d.ctx.Log.Info("全部执行完毕")
	return
}

type didiWaitLinkResp struct {
	ID                  int      `json:"id"`                  // 本次任务id
	SiteID              int      `json:"siteId"`              // 本站的 site id
	ResponseDomain      string   `json:"responseDomain"`      // 待换链接的域名
	ResponseKeywordsArr []string `json:"responseKeywordsArr"` // 待换链接的关键词列表
}

func (d *DidiAuto) getWaitLink() (res []didiWaitLinkResp, err error) {
	uri := fmt.Sprintf(didiWaitPageURL, d.SiteID)
	body, err := request.New().SetCookie(d.Cookie).SetReferer(uri).
		AddHeader("X-Requested-With", "XMLHttpRequest").AddHeader("Host", "didi.seowhy.com").GetBody(uri)
	if err != nil {
		d.ctx.Log.Error("获取站点列表出错", zap.Error(err), zap.String("url", uri))
		return
	}
	logBody := zap.String("body", string(body))
	data, err := simplejson.NewJson(body)
	if err != nil {
		d.ctx.Log.Error("序列化时出错", zap.Error(err), logBody)
		return
	}
	b, err := data.GetPath("data").MarshalJSON()
	if err != nil {
		d.ctx.Log.Error("获取json出错", zap.Error(err), logBody)
		return
	}
	if err = json.Unmarshal(b, &res); err != nil {
		d.ctx.Log.Error("解析json出错", zap.Error(err), logBody)
	}
	return
}

func (d *DidiAuto) run(item didiWaitLinkResp) {
	if item.ResponseDomain == "" {
		d.ctx.Log.Warn("待上链地址为空!")
		return
	}
	logDomain := zap.String("待上链地址", item.ResponseDomain)
	if len(item.ResponseKeywordsArr) == 0 {
		d.ctx.Log.Warn("锚文本列表为空!", logDomain)
		return
	}
	name := item.ResponseKeywordsArr[0]
	if name == "" {
		d.ctx.Log.Warn("锚文本字符串为空!", logDomain)
		return
	}
	if err := service.Link.Create(&entity.Link{Name: name, URL: item.ResponseDomain, Detect: true, DetectDelay: d.DetectDelay, Status: true}); err != nil {
		d.ctx.Log.Error("创建友情链接出错!", zap.Error(err), logDomain)
	} else {
		d.ctx.Log.Info(fmt.Sprintf("%s 上链成功", item.ResponseDomain))
	}
}

type didiPassDataResp struct {
	IsTrue int `json:"isTrue"` // 是否确认
}

func (d *DidiAuto) checkPass(item didiWaitLinkResp, n int) {
	uri := fmt.Sprintf(didiCheckPageURL, item.ID)
	body, err := request.New().SetCookie(d.Cookie).SetReferer(uri).
		AddHeader("X-Requested-With", "XMLHttpRequest").AddHeader("Host", "didi.seowhy.com").GetBody(uri)
	if err != nil {
		d.ctx.Log.Error("访问验证地址出错", zap.Error(err), zap.String("url", uri))
		return
	}
	logBody := zap.String("body", string(body))
	data, err := simplejson.NewJson(body)
	if err != nil {
		d.ctx.Log.Error("序列化时出错", zap.Error(err), logBody)
		return
	}
	b, err := data.GetPath("data").MarshalJSON()
	if err != nil {
		d.ctx.Log.Error("获取json出错", zap.Error(err), logBody)
		return
	}
	var res didiPassDataResp
	if err = json.Unmarshal(b, &res); err != nil {
		d.ctx.Log.Error("解析json出错", zap.Error(err), logBody)
		return
	}

	if res.IsTrue == 0 {
		if n < d.Retry {
			d.ctx.Log.Warn(fmt.Sprintf("验证第 %d 次失败, %d秒后继续验证...", n+1, d.RetrySleep))
			time.Sleep(time.Duration(d.RetrySleep) * time.Second)
			d.checkPass(item, n+1)
			return
		}
		d.ctx.Log.Error("验证失败", logBody, zap.String("link", item.ResponseDomain))
		return
	}
	d.ctx.Log.Info(fmt.Sprintf("%s 验证成功", item.ResponseDomain))
}
