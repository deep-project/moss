package widget

import (
	"go.uber.org/zap"
	"math"
	"moss/domain/config"
	"moss/domain/config/entity"
	"moss/domain/core/aggregate"
	coreEntity "moss/domain/core/entity"
	"moss/domain/core/repository/context"
	"moss/domain/core/service"
	coreUtils "moss/domain/core/utils"
	"moss/infrastructure/support/log"
)

type Widget struct {
}

func New() *Widget {
	return &Widget{}
}

func (*Widget) Head() string {
	return config.Config.Template.Head
}

func (*Widget) Footer() string {
	return config.Config.Template.Footer
}

func (w *Widget) Carousel() (res []entity.TemplateCarousel) {
	if !config.Config.Template.EnableCarousel {
		return
	}
	return config.Config.Template.Carousel
}

// Menu 模板导航
func (w *Widget) Menu() []aggregate.CategoryTree {
	var items []coreEntity.Category
	var err error
	if len(config.Config.Template.Menu.Select) > 0 {
		// 根据选择调用的导航数据
		items, err = service.Category.ListByIds(context.NewContextWithComment(config.Config.Template.Menu.Limit, "", "Menu"), config.Config.Template.Menu.Select)
		items = coreUtils.SortByIds[coreEntity.Category](items, config.Config.Template.Menu.Select) // 根据选择的ids排序
	} else {
		// 默认调用全部导航数据
		items, err = service.Category.List(context.NewContextWithComment(config.Config.Template.Menu.Limit, "", "Menu"))
	}
	if err != nil {
		log.Error("template widget error", zap.Error(err))
		return nil
	}
	return coreUtils.MakeCategoryTree(coreUtils.CategoryEntityListToCategoryTreeList(items), 0)
}

// Link 链接列表
func (w *Widget) Link() (res []coreEntity.Link) {
	res, err := service.Link.ListPublic(nil)
	log.ErrorShortcut("template widget error", err)
	return
}

// IndexList 首页列表
func (w *Widget) IndexList() (res []coreEntity.ArticleBase) {
	return w.simpleList(config.Config.Template.IndexList)
}

// GlobalList 全局列表
func (w *Widget) GlobalList() (res []coreEntity.ArticleBase) {
	return w.simpleList(config.Config.Template.GlobalList)
}

// 调用简单的列表
func (w *Widget) simpleList(opt *entity.TemplateList) (res []coreEntity.ArticleBase) {
	if opt.Limit <= 0 {
		return
	}
	var err error
	var ctx = context.NewContext(opt.Limit, opt.Order)
	if len(opt.CategoryIds) > 0 {
		res, err = service.Article.ListByCategoryIds(ctx, opt.CategoryIds)
	} else {
		res, err = service.Article.List(ctx)
	}
	log.ErrorShortcut("template widget error", err)
	return
}

// Breadcrumb 面包屑 通过分类ID调用
func (w *Widget) Breadcrumb(categoryID int) (res []coreEntity.Category) {
	res, err := service.Category.GetWithAncestorsReverse(context.NewContextWithComment(config.Config.More.ViewAllCategoryLimit, "", "Breadcrumb"), categoryID)
	log.ErrorShortcut("template widget error", err)
	return
}

// TagCloud 标签云
func (w *Widget) TagCloud() (res []coreEntity.Tag) {
	if config.Config.Template.TagCloud.Limit <= 0 {
		return
	}
	var err error
	var ctx = context.NewContextWithComment(config.Config.Template.TagCloud.Limit, config.Config.Template.TagCloud.Order, "TagCloud")
	if len(config.Config.Template.TagCloud.Select) > 0 {
		res, err = service.Tag.ListByIds(ctx, config.Config.Template.TagCloud.Select)
	} else {
		res, err = service.Tag.List(ctx)
	}
	log.ErrorShortcut("template widget error", err)
	return
}

// PageListResult 分页查询列表结果
type PageListResult struct {
	List          any   // 数据列表
	ListLength    int   // 列表长度
	Count         int64 // 数据总数
	PageNumber    int   // 当前页码
	PageTotal     int   // 总页数
	ExistNextPage bool  // 是否存在下一页
	DisableCount  bool  // 是否禁用count
}

// CategoryPageList 分类页列表
func (w *Widget) CategoryPageList(categoryID, pageNumber int) (res PageListResult) {
	if pageNumber == 0 {
		pageNumber = 1
	}
	var (
		opt               = config.Config.Template.CategoryPageList
		fastOffsetMinPage = config.Config.More.FastOffsetMinPage // 加速分页查询时，最小分页数
		fastOffset        = fastOffsetMinPage > 0 && pageNumber > fastOffsetMinPage
	)
	// 查询数据函数
	var listFun = func() (any, int) {
		list, err := service.Article.ListByCategoryID(&context.Context{Limit: opt.Limit, Order: opt.Order, Page: pageNumber, FastOffset: fastOffset}, categoryID)
		log.ErrorShortcut("template widget error", err)
		return list, len(list)
	}
	// 统计总数函数
	var countFun = func() (res int64) {
		res, err := service.Article.CountByCategoryID(categoryID)
		log.ErrorShortcut("template widget error", err)
		return
	}
	return w.pageList(opt, pageNumber, listFun, countFun)
}

// TagPageList 标签页列表
func (w *Widget) TagPageList(tagID, pageNumber int) (res PageListResult) {
	// 查询数据函数
	var listFun = func() (any, int) {
		list, err := service.Article.ListByTagID(&context.Context{
			Limit: config.Config.Template.TagPageList.Limit,
			Order: "id desc",
			Page:  pageNumber,
		}, tagID)
		log.ErrorShortcut("template widget error", err)
		return list, len(list)
	}
	// 统计总数函数
	var countFun = func() (res int64) {
		res, err := service.Mapping.CountByTagID(tagID)
		log.ErrorShortcut("template widget error", err)
		return
	}
	return w.pageList(config.Config.Template.TagPageList, pageNumber, listFun, countFun)
}

func (w *Widget) pageList(opt *entity.TemplateList, pageNumber int, listFun func() (any, int), countFun func() int64) (res PageListResult) {
	if pageNumber == 0 {
		pageNumber = 1
	}
	if opt.Limit <= 0 || (opt.MaxPage > 0 && pageNumber > opt.MaxPage) {
		return
	}
	res.PageNumber = pageNumber
	res.List, res.ListLength = listFun()
	res.DisableCount = opt.DisableCount
	if opt.DisableCount { // 如果禁用count 是否存在下一页取决于本页数量是否等于设定
		res.ExistNextPage = res.ListLength >= opt.Limit
		return
	}
	res.Count = countFun()
	res.PageTotal = w.computePageTotal(res.Count, opt.Limit, opt.MaxPage)
	res.ExistNextPage = res.PageTotal > res.PageNumber
	return
}

// 计算总页数
func (w *Widget) computePageTotal(count int64, limit, maxPage int) (res int) {
	if count == 0 || limit == 0 {
		return
	}
	res = int(math.Ceil(float64(count) / float64(limit)))
	if maxPage > 0 && res > maxPage {
		res = maxPage // 限制最大页
	}
	return
}
