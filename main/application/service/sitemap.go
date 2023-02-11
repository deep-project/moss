package service

import (
	"moss/application/dto"
	"moss/domain/config"
	configEntity "moss/domain/config/entity"
	"moss/domain/core/entity"
	"moss/domain/core/repository/context"
	"moss/domain/core/service"
	"strings"
	"time"
)

// Sitemap 站点地图数据
var Sitemap = new(sitemap)

type sitemap struct{}

// ArticleList 文章站点地图数据列表
func (s *sitemap) ArticleList() (res []entity.ArticleBase, err error) {
	if config.Config.Sitemap.Article.Limit > 0 {
		res, err = service.Article.ListAfterCreateTime(s.listOption(config.Config.Sitemap.Article))
	}
	return
}

// CategoryList 分类站点地图数据列表
func (s *sitemap) CategoryList() (res []entity.Category, err error) {
	if config.Config.Sitemap.Category.Limit > 0 {
		res, err = service.Category.ListAfterCreateTime(s.listOption(config.Config.Sitemap.Category))
	}
	return
}

// TagList 标签站点地图数据列表
func (s *sitemap) TagList() (res []entity.Tag, err error) {
	if config.Config.Sitemap.Tag.Limit > 0 {
		res, err = service.Tag.ListAfterCreateTime(s.listOption(config.Config.Sitemap.Tag))
	}
	return
}

func (s *sitemap) listOption(opt *configEntity.SitemapOption) (ctx *context.Context, t int64) {
	if opt.InHours > 0 {
		t = time.Now().Unix() - int64(opt.InHours)*60*60
	}
	return context.NewContext(opt.Limit, ""), t
}

func (s *sitemap) ArticleText() (res string, err error) {
	var urls []string
	items, err := s.ArticleList()
	if err != nil {
		return
	}
	for _, item := range items {
		urls = append(urls, item.FullURL())
	}
	return strings.Join(urls, "\n"), nil
}

func (s *sitemap) CategoryText() (res string, err error) {
	var urls []string
	items, err := s.CategoryList()
	if err != nil {
		return
	}
	for _, item := range items {
		urls = append(urls, item.FullURL())
	}
	return strings.Join(urls, "\n"), nil
}

func (s *sitemap) TagText() (res string, err error) {
	var urls []string
	items, err := s.TagList()
	if err != nil {
		return
	}
	for _, item := range items {
		urls = append(urls, item.FullURL())
	}
	return strings.Join(urls, "\n"), nil
}

func (s *sitemap) ArticleXML() (res string, err error) {
	var xml = dto.NewSitemapXML()
	items, err := s.ArticleList()
	if err != nil {
		return
	}
	for _, item := range items {
		xml.URL = append(xml.URL, dto.NewSitemapURL(item.FullURL(), item.CreateTime, config.Config.Sitemap.Article.ChangeFreq, config.Config.Sitemap.Article.Priority))
	}
	return xml.String()
}

func (s *sitemap) CategoryXML() (res string, err error) {
	var xml = dto.NewSitemapXML()
	items, err := s.CategoryList()
	if err != nil {
		return
	}
	for _, item := range items {
		xml.URL = append(xml.URL, dto.NewSitemapURL(item.FullURL(), item.CreateTime, config.Config.Sitemap.Category.ChangeFreq, config.Config.Sitemap.Category.Priority))
	}
	return xml.String()
}

func (s *sitemap) TagXML() (res string, err error) {
	var xml = dto.NewSitemapXML()
	items, err := s.TagList()
	if err != nil {
		return
	}
	for _, item := range items {
		xml.URL = append(xml.URL, dto.NewSitemapURL(item.FullURL(), item.CreateTime, config.Config.Sitemap.Tag.ChangeFreq, config.Config.Sitemap.Tag.Priority))
	}
	return xml.String()
}
