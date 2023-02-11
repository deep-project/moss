package captcha

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"github.com/afocus/captcha"
	"github.com/allegro/bigcache/v3"
	"github.com/duke-git/lancet/v2/random"
	"image"
	"image/color"
	"image/png"
	"moss/resources"
	"strings"
	"time"
)

var Client = New()

type Captcha struct {
	store *bigcache.BigCache
	draw  *captcha.Captcha
}

type Options struct {
	Expire          time.Duration       // 过期时间
	Width           int                 // 图片宽
	Height          int                 // 图片高
	NoiseLevel      captcha.DisturLevel // 干扰强度
	FrontColor      []color.Color       // 设置文字色,多个随机使用
	BackgroundColor []color.Color       // 设置背景色,多个随机使用
}

func DefaultOptions() Options {
	return Options{
		Expire:     1 * time.Minute,
		Width:      198,
		Height:     63,
		NoiseLevel: captcha.MEDIUM,
	}
}

func New(val ...Options) *Captcha {
	var opt Options
	if len(val) > 0 {
		opt = val[0]
	} else {
		opt = DefaultOptions()
	}

	store, _ := bigcache.New(context.Background(), bigcache.Config{
		Shards:             16,
		LifeWindow:         opt.Expire,
		CleanWindow:        5 * time.Second,
		MaxEntriesInWindow: 1000,
		MaxEntrySize:       64,
		HardMaxCacheSize:   8192,
	})
	draw := captcha.New()
	draw.SetSize(opt.Width, opt.Height) // 尺寸
	draw.SetDisturbance(opt.NoiseLevel) // 干扰强度
	draw.SetFrontColor(opt.FrontColor...)
	draw.SetBkgColor(opt.BackgroundColor...)

	font, _ := resources.App.ReadFile("app/comic.ttf")
	_ = draw.AddFontFromBytes(font)
	return &Captcha{store: store, draw: draw}
}

func (c *Captcha) Delete(id string) {
	_ = c.store.Delete(id)
}

func (c *Captcha) Verify(id, answer string) error {
	if answer == "" {
		return errors.New("captcha is required")
	}
	if id == "" {
		return errors.New("captcha id is required")
	}
	b, err := c.store.Get(id)
	if err != nil {
		if err.Error() == "Entry not found" {
			return errors.New("captcha has expired")
		}
		return err
	}
	if strings.ToLower(string(b)) != strings.ToLower(answer) {
		return errors.New("captcha is wrong")
	}
	return nil
}

func (c *Captcha) Number(n int) *Result {
	return c.Generate(n, captcha.NUM)
}

// LetterLower 小写字母
func (c *Captcha) LetterLower(n int) *Result {
	return c.Generate(n, captcha.LOWER)
}

// LetterUpper 大写字母
func (c *Captcha) LetterUpper(n int) *Result {
	return c.Generate(n, captcha.UPPER)
}

// String 全部
func (c *Captcha) String(n int) *Result {
	return c.Generate(n, captcha.ALL)
}

// StringSimple 去掉复杂易混淆的字符
func (c *Captcha) StringSimple(n int) *Result {
	return c.Generate(n, captcha.CLEAR)
}

func (c *Captcha) Generate(len int, t captcha.StrType) *Result {
	return c.newResult(c.draw.Create(len, t))
}

func (c *Captcha) newResult(img image.Image, answer string) *Result {
	uuid, _ := random.UUIdV4()
	_ = c.store.Set(uuid, []byte(answer))
	return &Result{img: img, answer: answer, id: uuid}
}

type Result struct {
	img    image.Image
	answer string
	id     string
}

func (r *Result) Image() (img image.Image, id string) {
	return r.img, r.id
}

func (r *Result) Base64() (bs64 string, id string) {
	buff := bytes.NewBuffer(nil)
	_ = png.Encode(buff, r.img)
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(buff.Bytes()), r.id
}
