package plugins

import (
	"moss/domain/core/entity"
	"moss/domain/core/service"
	pluginEntity "moss/domain/support/entity"
	"moss/infrastructure/utils/htmlx"
)

type GenerateDescription struct {
	Enable bool `json:"enable"` // 启用
	Length int  `json:"length"` // 长度
}

func NewGenerateDescription() *GenerateDescription {
	return &GenerateDescription{Enable: true, Length: 150}
}

func (d *GenerateDescription) Info() *pluginEntity.PluginInfo {
	return &pluginEntity.PluginInfo{
		ID:    "GenerateDescription",
		About: "generate description when created",
	}
}

func (d *GenerateDescription) Load(ctx *pluginEntity.Plugin) error {
	service.Article.AddCreateBeforeEvents(d)
	return nil
}

func (d *GenerateDescription) ArticleCreateBefore(item *entity.Article) (err error) {
	if !d.Enable || d.Length <= 0 || item.Description != "" || item.Content == "" {
		return
	}

	text := htmlx.GetTextFromHTML(item.Content)
	if text == "" {
		return
	}
	textRune := []rune(text)
	if len(textRune) > d.Length {
		item.Description = string(textRune[0:d.Length])
	} else {
		item.Description = string(textRune[0:])
	}
	return
}

func (d *GenerateDescription) Run(ctx *pluginEntity.Plugin) (err error) {
	return nil
}
