package query

import (
	"github.com/duke-git/lancet/v2/slice"
	"moss/domain/config"
	"moss/domain/core/aggregate"
	"moss/domain/core/entity"
	"moss/domain/core/repository/context"
	"moss/domain/core/service"
	"moss/infrastructure/support/log"
)

type Category struct {
	limit   int
	order   string
	comment string
}

func NewCategory() *Category {
	return &Category{}
}

func (c *Category) Limit(val int) *Category {
	c.limit = val
	return c
}

func (c *Category) Order(val string) *Category {
	c.order = val
	return c
}

func (c *Category) Comment(val string) *Category {
	c.comment = val
	return c
}

func (c *Category) context() *context.Context {
	if c.limit == 0 {
		c.limit = 20 // 强制限制数量
	}
	return context.NewContextWithComment(c.limit, c.order, c.comment)
}

// Get by id
func (c *Category) Get(id int) *entity.Category {
	res, err := service.Category.Get(id)
	log.WarnShortcut("template query error", err)
	return res
}

// List 调用列表
func (c *Category) List() (res []entity.Category) {
	res, err := service.Category.List(c.context())
	log.WarnShortcut("template query error", err)
	return
}

// ListByID 根据ID调用文章列表
func (c *Category) ListByID(ids ...int) (res []entity.Category) {
	res, err := service.Category.ListByIds(c.context(), ids)
	log.WarnShortcut("template query error", err)
	return
}

// PseudorandomList 伪随机列表
func (c *Category) PseudorandomList() (res []entity.Category) {
	res, err := service.Category.PseudorandomList(c.context())
	log.WarnShortcut("template query error", err)
	return
}

// GetWithAncestors 获取分类和其祖先
func (c *Category) GetWithAncestors(id int) (res []entity.Category) {
	res, err := service.Category.GetWithAncestors(context.NewContextWithComment(config.Config.More.ViewAllCategoryLimit, c.order, c.comment), id)
	log.ErrorShortcut("template query error", err)
	return
}

// GetWithParent 获取分类和其夫分类
func (c *Category) GetWithParent(id int) (res []entity.Category) {
	res, err := service.Category.GetWithParent(id)
	slice.Reverse(res)
	log.ErrorShortcut("template query error", err)
	return
}

// Descendants 获取分类所有后代
func (c *Category) Descendants(rootID int) (res []aggregate.CategoryTree) {
	res, err := service.Category.ListDescendants(context.NewContextWithComment(config.Config.More.ViewAllCategoryLimit, c.order, c.comment), rootID)
	log.ErrorShortcut("template query error", err)
	return
}

// Children 子分类列表
func (c *Category) Children(parentID int) (res []entity.Category) {
	res, err := service.Category.ListChildren(c.context(), parentID)
	log.ErrorShortcut("template query error", err)
	return
}

// ListWithChildren 获取列表并包含子分类
func (c *Category) ListWithChildren(ids ...int) (res []aggregate.CategoryTree) {
	var err error
	if len(ids) == 0 {
		res, err = service.Category.ListRootWithChildren(c.context())
	} else {
		res, err = service.Category.ListByIdsWithChildren(c.context(), ids)
	}
	log.ErrorShortcut("template query error", err)
	return
}
