package service

import (
	"errors"
	"moss/domain/config"
	"moss/domain/core/entity"
	"moss/domain/core/service"
	"moss/infrastructure/support/template"
	"path/filepath"
	"strconv"
)

var Render = new(RenderService)

type RenderService struct {
}

func (r *RenderService) Index() ([]byte, error) {
	return template.Render("template/index.html", template.Binds{
		Page: template.Page{
			Name:        "index",
			Title:       config.Config.Site.Title,
			Keywords:    config.Config.Site.Keywords,
			Description: config.Config.Site.Description,
		},
	})
}

func (r *RenderService) TemplatePage(path string) ([]byte, error) {
	return template.Render(filepath.Join("page", path), template.Binds{
		Page: template.Page{
			Name: "page",
			Path: path,
		},
		Data: map[string]any{},
	})
}

func (r *RenderService) ArticleBySlug(slug string) (_ []byte, err error) {
	item, err := service.Article.GetBySlug(slug)
	if err != nil {
		return
	}
	return r.Article(item)
}

func (r *RenderService) Article(item *entity.Article) (_ []byte, err error) {
	if item == nil {
		err = errors.New("item is nil")
		return
	}
	return template.Render("template/article.html", template.Binds{
		Page: template.Page{
			Name:        "article",
			Title:       item.Title + " - " + config.Config.Site.Name,
			Keywords:    item.Keywords,
			Description: item.Description,
		},
		Data: item,
	})
}

func (r *RenderService) CategoryBySlug(slug string, page int) (_ []byte, err error) {
	item, err := service.Category.GetBySlug(slug)
	if err != nil {
		return
	}
	return r.Category(item, page)
}

func (r *RenderService) Category(item *entity.Category, page int) (_ []byte, err error) {
	if item == nil {
		err = errors.New("item is nil")
		return
	}
	var pageTitle string
	if page > 1 {
		pageTitle = " - " + strconv.Itoa(page)
	}
	var title = item.Name
	if item.Title != "" {
		title = item.Title
	}
	return template.Render("template/category.html", template.Binds{
		Page: template.Page{
			Name:        "category",
			Title:       title + pageTitle + " - " + config.Config.Site.Name,
			Keywords:    item.Keywords,
			Description: item.Description,
			PageNumber:  page,
		},
		Data: item,
	})
}

func (r *RenderService) TagBySlug(slug string, page int) (_ []byte, err error) {
	item, err := service.Tag.GetBySlug(slug)
	if err != nil {
		return
	}
	return r.Tag(item, page)
}

func (r *RenderService) Tag(item *entity.Tag, page int) (_ []byte, err error) {
	if item == nil {
		err = errors.New("item is nil")
		return
	}
	var pageTitle string
	if page > 1 {
		pageTitle = " - " + strconv.Itoa(page)
	}
	var title = item.Name
	if item.Title != "" {
		title = item.Title
	}
	return template.Render("template/tag.html", template.Binds{
		Page: template.Page{
			Name:        "tag",
			Title:       title + pageTitle + " - " + config.Config.Site.Name,
			Keywords:    item.Keywords,
			Description: item.Description,
			PageNumber:  page,
		},
		Data: item,
	})
}
