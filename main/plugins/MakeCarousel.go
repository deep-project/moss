package plugins

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"go.uber.org/zap"
	"image"
	"moss/domain/config"
	"moss/domain/config/entity"
	configService "moss/domain/config/service"
	"moss/domain/core/repository/context"
	"moss/domain/core/service"
	pluginEntity "moss/domain/support/entity"
	"moss/infrastructure/utils/request"
	"strings"
)

type MakeCarousel struct {
	Limit          int     `json:"limit"`
	QuerySize      int     `json:"query_size"`       // 查询数量
	MinWidth       int     `json:"min_width"`        // 限制选取图片的最小宽度
	MinHeight      int     `json:"min_height"`       // 限制选取图片的最小高度
	MaxHeightRatio float64 `json:"max_height_ratio"` // 最大高度比例 默认0.7 如果高度大于宽度的0.7 则略过
	DownTimeout    int     `json:"down_timeout"`     // 下载超时

	ctx *pluginEntity.Plugin
}

func NewMakeCarousel() *MakeCarousel {
	return &MakeCarousel{
		Limit:          5,
		QuerySize:      100,
		MinWidth:       240,
		MinHeight:      150,
		MaxHeightRatio: 0.7,
		DownTimeout:    20,
	}
}

func (p *MakeCarousel) Info() *pluginEntity.PluginInfo {
	return &pluginEntity.PluginInfo{
		ID:         "MakeCarousel",
		About:      "make carousel using articles",
		RunEnable:  true,
		CronEnable: true,
		PluginInfoPersistent: pluginEntity.PluginInfoPersistent{
			CronStart: false,
			CronExp:   "@every 24h",
		},
	}
}

func (p *MakeCarousel) Load(ctx *pluginEntity.Plugin) error { return nil }

func (p *MakeCarousel) Run(ctx *pluginEntity.Plugin) (err error) {
	p.ctx = ctx
	if p.QuerySize <= 0 {
		p.QuerySize = 100
	}
	if p.MaxHeightRatio == 0 {
		p.MaxHeightRatio = 0.7
	}
	if p.DownTimeout == 0 {
		p.DownTimeout = 30
	}
	if p.Limit <= 0 {
		p.ctx.Log.Warn("limit is undefined")
		return
	}
	itemsBase, err := service.Article.ListExistThumbnail(context.NewContext(p.QuerySize, "id desc"))
	if err != nil {
		p.ctx.Log.Error("query base list error", zap.Error(err))
		return
	}
	if len(itemsBase) == 0 {
		p.ctx.Log.Warn("not found article existed thumbnail")
		return
	}
	var ids []int
	for _, item := range itemsBase {
		ids = append(ids, item.ID)
	}
	itemsDetail, err := service.Article.ListDetailByIds(nil, ids)
	if err != nil {
		p.ctx.Log.Error("query detail list error", zap.Error(err), zap.Ints("ids", ids))
		return
	}
	items := service.Article.MergeBaseListAndDetailList(itemsBase, itemsDetail)
	if len(items) == 0 {
		p.ctx.Log.Warn("not found items", zap.Ints("ids", ids))
		return
	}

	var success int
	var res []entity.TemplateCarousel
	for _, item := range items {
		if success >= p.Limit {
			break
		}
		img := p.findImg(item.Content)
		if img == "" {
			continue
		}
		res = append(res, entity.TemplateCarousel{Image: img, Link: item.URL(), Title: item.Title})
		success++
	}
	if success == 0 {
		p.ctx.Log.Warn("not found article", zap.Ints("ids", ids))
		return
	}
	// 把旧的加到后面
	res = append(res, config.Config.Template.Carousel...)
	// 去重
	res = p.removeRepeat(res)
	// 限制数量
	if len(res) > p.Limit {
		res = res[0:p.Limit]
	}
	// 赋值新的结果
	config.Config.Template.Carousel = res
	// 推送到数据库
	if err = configService.Push(config.Config.Template); err != nil {
		p.ctx.Log.Error("push data error", zap.Error(err))
		return
	}
	p.ctx.Log.Info("make successful", zap.Int("length", len(res)), zap.Any("result", res))
	return
}

func (p *MakeCarousel) findImg(content string) (res string) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		p.ctx.Log.Error("format content error", zap.Error(err))
		return
	}
	doc.Find("img").EachWithBreak(func(i int, s *goquery.Selection) bool {
		src, _ := s.Attr("src")
		if src == "" || strings.HasPrefix(src, "data:") {
			return true
		}
		if strings.HasPrefix(src, "//") {
			src = "http:" + src
		}
		if strings.HasPrefix(src, "/") {
			src = config.Config.Site.GetURL() + src
		}
		// 图片已存在
		if p.isExist(src, config.Config.Template.Carousel) {
			p.ctx.Log.Warn("image is exist", zap.String("src", src))
			return false
		}
		// 下载图片
		file, err := request.New().SetTimeoutSeconds(p.DownTimeout).GetBody(src)
		if err != nil {
			p.ctx.Log.Error("image download failed", zap.Error(err), zap.String("src", src))
			return true
		}
		// 获取图片尺寸
		size, _, _ := image.DecodeConfig(bytes.NewReader(file))
		if size.Width == 0 || size.Height == 0 {
			p.ctx.Log.Warn("image size is zero", zap.String("src", src))
			return true
		}
		// 不符合最小设定
		if size.Width < p.MinWidth || size.Height < p.MinHeight {
			p.ctx.Log.Warn("size too small", zap.String("src", src), zap.Int("width", size.Width), zap.Int("height", size.Height))
			return true
		}
		// 如果高宽比大于设定的比例
		if float64(size.Height) > float64(size.Width)*p.MaxHeightRatio {
			p.ctx.Log.Warn("ratio does not conform", zap.String("src", src), zap.Int("width", size.Width), zap.Int("height", size.Height))
			return true
		}
		res = src
		return false // 退出循环
	})
	return
}

// 去重
func (p *MakeCarousel) removeRepeat(input []entity.TemplateCarousel) (res []entity.TemplateCarousel) {
	for _, item := range input {
		if p.isExist(item.Image, res) {
			continue
		}
		res = append(res, item)
	}
	return
}

func (p *MakeCarousel) isExist(img string, input []entity.TemplateCarousel) bool {
	for _, item := range input {
		if img == item.Image {
			return true
		}
	}
	return false
}
