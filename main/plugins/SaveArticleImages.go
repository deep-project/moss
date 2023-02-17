package plugins

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"
	"go.uber.org/zap"
	"image"
	"moss/domain/config"
	"moss/domain/core/entity"
	"moss/domain/core/service"
	pluginEntity "moss/domain/support/entity"
	"moss/infrastructure/persistent/storage"
	"moss/infrastructure/support/upload"
	"moss/infrastructure/utils/imagex"
	"moss/infrastructure/utils/request"
	"strconv"
	"strings"
)

type SaveArticleImages struct {
	EnableOnCreate bool `json:"enable_on_create"` // 创建时执行
	EnableOnUpdate bool `json:"enable_on_update"` // 更新时执行

	MaxWidth          int    `json:"max_width"`           // 最大图片宽度(像素)，大于此宽度将被等比例缩放
	MaxHeight         int    `json:"max_height"`          // 最大图片高度(像素)，大于此高度将被等比例缩放
	ThumbWidth        int    `json:"thumb_width"`         // 缩略图宽度(像素)
	ThumbHeight       int    `json:"thumb_height"`        // 缩略图高度(像素)
	ThumbMinWidth     int    `json:"thumb_min_width"`     // 选取缩略图时，限制最小缩略图宽度(像素)，小于此宽度的图片不会被选取成缩略图
	ThumbMinHeight    int    `json:"thumb_min_height"`    // 选取缩略图时，限制最小缩略图高度(像素)，小于此高度的图片不会被选取成缩略图
	AlwaysResize      bool   `json:"always_resize"`       // 是否始终缩放一下图片，已减少图片体积
	ThumbExtractFocus bool   `json:"thumb_extract_focus"` // 生成缩略图是提取焦点方式生成
	RemoveIfDownFail  bool   `json:"remove_if_down_fail"` // 下载失败是否删除
	DownRetry         int    `json:"down_retry"`          // 重试次数
	DownReferer       string `json:"down_referer"`        // 下载referer
	DownProxy         string `json:"down_proxy"`          // 下载代理

	ctx         *pluginEntity.Plugin
	downReferer []saveArticleImagesDownReferer
}

func NewSaveArticleImages() *SaveArticleImages {
	return &SaveArticleImages{
		EnableOnCreate:    true,
		EnableOnUpdate:    true,
		DownRetry:         3,
		MaxWidth:          1000,
		MaxHeight:         2000,
		ThumbWidth:        230,
		ThumbHeight:       138,
		ThumbMinWidth:     100,
		ThumbMinHeight:    100,
		AlwaysResize:      true,
		ThumbExtractFocus: true,
		RemoveIfDownFail:  true,
		DownReferer:       "bdimg bdstatic http://www.baidu.com/\ntoutiaoimg http://www.toutiao.com/",
	}
}

func (s *SaveArticleImages) Info() *pluginEntity.PluginInfo {
	return &pluginEntity.PluginInfo{
		ID:    "SaveArticleImages",
		About: "save article images",
	}
}

func (s *SaveArticleImages) Run(ctx *pluginEntity.Plugin) error {
	return nil
}

func (s *SaveArticleImages) Load(ctx *pluginEntity.Plugin) error {
	s.ctx = ctx
	service.Article.AddCreateBeforeEvents(s)
	service.Article.AddUpdateBeforeEvents(s)
	return nil
}
func (s *SaveArticleImages) ArticleCreateBefore(item *entity.Article) (err error) {
	if !s.EnableOnCreate {
		return nil
	}
	return s.Save(item)
}
func (s *SaveArticleImages) ArticleUpdateBefore(item *entity.Article) (err error) {
	if !s.EnableOnUpdate {
		return nil
	}
	return s.Save(item)
}

func (s *SaveArticleImages) Save(item *entity.Article) error {
	s.initDownReferer()
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(item.Content))
	if err != nil {
		s.ctx.Log.Error("format html document error", zap.Error(err), zap.String("title", item.Title))
		return err
	}
	doc.Find("img").Each(s.eachSave(item))
	s.saveThumbnail(item)
	html, err := doc.Find("body").Html()
	if err != nil {
		s.ctx.Log.Error("get html code error", zap.Error(err), zap.String("title", item.Title))
		return err
	}
	item.Content = html
	return nil
}

func (s *SaveArticleImages) eachSave(item *entity.Article) func(i int, sn *goquery.Selection) {
	return func(i int, sn *goquery.Selection) {
		src, ok := sn.Attr("src")
		if !ok || src == "" {
			sn.Remove()
			return
		}
		if strings.HasPrefix(src, config.Config.Upload.Domain) { // upload域开头直接跳过
			return
		}
		if !strings.HasPrefix(src, "http") && !strings.HasPrefix(src, "//") { // 非远程图片
			return
		}
		if strings.HasPrefix(src, "data:") { // base64图片
			return
		}
		// 下载图片
		file, err := s.down(item, src)
		if err != nil && s.RemoveIfDownFail {
			sn.Remove()
			return
		}
		// 获取并判断图片类型
		imageType, err := filetype.Image(file)
		if imageType == types.Unknown || err != nil {
			s.ctx.Log.Warn("file is not a image type", s.logInfo(item, src, err)...)
			sn.Remove()
			return
		}
		// 获取图片尺寸
		size, _, err := image.DecodeConfig(bytes.NewReader(file))
		if size.Width == 0 || size.Height == 0 || err != nil {
			s.ctx.Log.Warn("image size error", s.logInfo(item, src, err)...)
			return
		}
		// 计算图片尺寸
		var width, height = imagex.ComputeScale(size.Width, size.Height, s.MaxWidth, s.MaxHeight)
		// 图片缩放，可以减少图片体积
		if s.AlwaysResize || size.Width > width || size.Height > height {
			if file, err = imagex.New().SetWidth(width).SetHeight(height).ResizeByte(file); err != nil {
				s.ctx.Log.Warn("image resize error", s.logInfo(item, src, err)...)
				return
			}
		}
		// 上传图片
		hashSrc := cryptor.Md5String(src)
		uploadResult, err := upload.Upload(hashSrc, imageType.Extension, storage.NewSetValueBytes(file))
		if err != nil {
			s.ctx.Log.Warn("upload image error", s.logInfo(item, src, err)...)
			return
		}
		s.ctx.Log.Info("upload image success", append(s.logInfo(item, src, nil), zap.String("url", uploadResult.URL))...)
		// 设置标签属性
		sn.SetAttr("src", uploadResult.URL)
		sn.SetAttr("width", strconv.Itoa(width))
		sn.SetAttr("height", strconv.Itoa(height))

		// 上传缩略图
		if item.Thumbnail == "" && size.Width >= s.ThumbMinWidth && size.Height >= s.ThumbMinHeight {
			// 直接把内容中的图片保存成缩略图
			if err = s.uploadThumbnail(item, file, hashSrc+"_thumbnail", imageType.Extension); err != nil {
				s.ctx.Log.Warn("upload thumbnail error", s.logInfo(item, src, err)...)
				return
			}
		}
	}
}

func (s *SaveArticleImages) logInfo(item *entity.Article, src string, err error) []zap.Field {
	return []zap.Field{zap.String("url", src), zap.String("title", item.Title), zap.Error(err)}
}

// 上传缩略图
func (s *SaveArticleImages) uploadThumbnail(item *entity.Article, file []byte, name, ext string) (err error) {
	if s.ThumbWidth > 0 || s.ThumbHeight > 0 {
		var imgLib = imagex.New().SetWidth(s.ThumbWidth).SetHeight(s.ThumbHeight)
		if s.ThumbExtractFocus {
			file, err = imgLib.CropByte(file)
		} else {
			file, err = imgLib.ThumbnailByte(file)
		}
		if err != nil {
			return
		}
	}
	thumbUploadResult, err := upload.Upload(name, ext, storage.NewSetValueBytes(file))
	if err != nil {
		return
	}
	s.ctx.Log.Info("upload thumbnail success", zap.String("title", item.Title), zap.String("url", thumbUploadResult.URL))
	item.Thumbnail = thumbUploadResult.URL
	return
}

func (s *SaveArticleImages) down(item *entity.Article, uri string) (file []byte, err error) {
	file, err = request.New().SetRetry(s.DownRetry).SetProxyURLStr(s.DownReferer).SetReferer(s.getDownReferer(uri)).GetBody(uri)
	if err != nil {
		s.ctx.Log.Warn("down file error", s.logInfo(item, uri, err)...)
	}
	return
}

func (s *SaveArticleImages) saveThumbnail(item *entity.Article) {
	if item.Thumbnail == "" {
		return
	}
	if strings.HasPrefix(item.Thumbnail, config.Config.Upload.Domain) { // upload域开头直接跳过
		return
	}
	// 下载图片
	file, err := s.down(item, item.Thumbnail)
	if err != nil && s.RemoveIfDownFail {
		item.Thumbnail = ""
		return
	}
	// 获取并判断图片类型
	imageType, err := filetype.Image(file)
	if imageType == types.Unknown || err != nil {
		s.ctx.Log.Warn("thumbnail is not a image type", s.logInfo(item, item.Thumbnail, err)...)
		item.Thumbnail = ""
		return
	}
	if err = s.uploadThumbnail(item, file, cryptor.Md5String(item.Thumbnail)+"_thumbnail", imageType.Extension); err != nil {
		s.ctx.Log.Warn("upload thumbnail error", s.logInfo(item, item.Thumbnail, err)...)
	}
}

type saveArticleImagesDownReferer struct {
	rule    string
	referer string
}

func (s *SaveArticleImages) initDownReferer() {
	if s.DownReferer == "" {
		return
	}
	for _, line := range strings.Split(s.DownReferer, "\n") {
		arr := strings.Split(line, " ")
		arrLen := len(arr)
		if arrLen < 2 {
			continue
		}
		referer := arr[arrLen-1]
		newArr := arr[:arrLen-1]
		for _, rule := range newArr {
			s.downReferer = append(s.downReferer, saveArticleImagesDownReferer{rule: rule, referer: referer})
		}
	}
}

func (s *SaveArticleImages) getDownReferer(src string) string {
	for _, v := range s.downReferer {
		if strings.Contains(src, v.rule) {
			return v.referer
		}
	}
	return ""
}
