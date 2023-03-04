package plugins

import (
	"github.com/duke-git/lancet/v2/random"
	"github.com/jaevor/go-nanoid"
	"github.com/rs/xid"
	"github.com/sony/sonyflake"
	"github.com/speps/go-hashids/v2"
	"github.com/yitter/idgenerator-go/idgen"
	"go.uber.org/zap"
	"moss/domain/core/entity"
	"moss/domain/core/service"
	pluginEntity "moss/domain/support/entity"
	"strconv"
)

type GenerateSlug struct {
	Article  *GenerateSlugOption `json:"article"`
	Category *GenerateSlugOption `json:"category"`
	Tag      *GenerateSlugOption `json:"tag"`
}

func init() {
	idgen.SetIdGenerator(idgen.NewIdGeneratorOptions(1))
}

type GenerateSlugOption struct {
	Style string `json:"style"`
}

func NewGenerateSlug() *GenerateSlug {
	return &GenerateSlug{
		Article:  &GenerateSlugOption{Style: "snowflake"},
		Category: &GenerateSlugOption{Style: "snowflake"},
		Tag:      &GenerateSlugOption{Style: "snowflake"},
	}
}

func (a *GenerateSlug) Info() *pluginEntity.PluginInfo {
	return &pluginEntity.PluginInfo{
		ID:    "GenerateSlug",
		About: "generate slug when created",
	}
}

func (a *GenerateSlug) Load(ctx *pluginEntity.Plugin) error {
	service.Article.AddCreateBeforeEvents(&GenerateSlugEvent{ctx: ctx, opt: a.Article})
	service.Category.AddCreateBeforeEvents(&GenerateSlugEvent{ctx: ctx, opt: a.Category})
	service.Tag.AddCreateBeforeEvents(&GenerateSlugEvent{ctx: ctx, opt: a.Tag})
	return nil
}

func (a *GenerateSlug) Run(ctx *pluginEntity.Plugin) error {
	return nil
}

type GenerateSlugEvent struct {
	ctx *pluginEntity.Plugin
	opt *GenerateSlugOption
}

func (e *GenerateSlugEvent) ArticleCreateBefore(item *entity.Article) (err error) {
	if item.Slug == "" {
		item.Slug, err = e.makeSlug()
	}
	return
}

func (e *GenerateSlugEvent) CategoryCreateBefore(item *entity.Category) (err error) {
	if item.Slug == "" {
		item.Slug, err = e.makeSlug()
	}
	return
}

func (e *GenerateSlugEvent) TagCreateBefore(item *entity.Tag) (err error) {
	if item.Slug == "" {
		item.Slug, err = e.makeSlug()
	}
	return
}

func (e *GenerateSlugEvent) makeSlug() (res string, err error) {
	switch e.opt.Style {
	case "uuid":
		res, err = random.UUIdV4()
	case "xid":
		res, err = xid.New().String(), nil
	case "hashSnowflake":
		res, err = e.hashSnowflake()
	case "idgenerator":
		res, err = e.idgeneratorStr()
	case "hashIdgenerator":
		res, err = e.hashIdgenerator()
	case "nanoid":
		res, err = e.nanoid()
	case "nanoid8":
		res, err = e.nanoid8()
	default: // snowflake
		res, err = e.snowflakeStr()
	}
	if err != nil {
		e.ctx.Log.Error("make slug error", zap.Error(err))
	}
	return
}

var generateIdEventHashIdsHandle, _ = hashids.NewWithData(&hashids.HashIDData{Salt: "moss", Alphabet: hashids.DefaultAlphabet})

func (e *GenerateSlugEvent) hashids(id int64) (string, error) {
	return generateIdEventHashIdsHandle.EncodeInt64([]int64{id})
}

var snowflakeHandle = sonyflake.NewSonyflake(sonyflake.Settings{MachineID: func() (uint16, error) {
	return 1, nil
}})

func (e *GenerateSlugEvent) snowflake() (uint64, error) {
	return snowflakeHandle.NextID()
}

func (e *GenerateSlugEvent) snowflakeStr() (string, error) {
	id, err := e.snowflake()
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(int64(id), 10), nil
}

func (e *GenerateSlugEvent) hashSnowflake() (string, error) {
	id, err := e.snowflake()
	if err != nil {
		return "", err
	}
	return e.hashids(int64(id))
}

func (e *GenerateSlugEvent) idgeneratorStr() (string, error) {
	return strconv.FormatInt(idgen.NextId(), 10), nil
}

func (e *GenerateSlugEvent) hashIdgenerator() (string, error) {
	return e.hashids(idgen.NextId())
}

func (e *GenerateSlugEvent) nanoid() (string, error) {
	genFunc, err := nanoid.Standard(21)
	if err != nil {
		return "", err
	}
	return genFunc(), nil
}

func (e *GenerateSlugEvent) nanoid8() (string, error) {
	genFunc, err := nanoid.Standard(8)
	if err != nil {
		return "", err
	}
	return genFunc(), nil
}
