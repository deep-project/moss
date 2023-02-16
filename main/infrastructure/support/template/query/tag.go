package query

import (
	"moss/domain/core/entity"
	"moss/domain/core/repository/context"
	"moss/domain/core/service"
	"moss/infrastructure/support/log"
)

type Tag struct {
	limit   int
	order   string
	comment string
}

func NewTag() *Tag {
	return &Tag{}
}

func (t *Tag) Limit(val int) *Tag {
	t.limit = val
	return t
}

func (t *Tag) Order(val string) *Tag {
	t.order = val
	return t
}

func (t *Tag) Comment(val string) *Tag {
	t.comment = val
	return t
}

func (t *Tag) context() *context.Context {
	if t.limit == 0 {
		t.limit = 20 // 强制限制数量
	}
	return context.NewContextWithComment(t.limit, t.order, t.comment)
}

// Get by id
func (t *Tag) Get(id int) *entity.Tag {
	res, err := service.Tag.Get(id)
	log.ErrorShortcut("template query error", err)
	return res
}

func (t *Tag) List() (res []entity.Tag) {
	res, err := service.Tag.List(t.context())
	log.ErrorShortcut("template query error", err)
	return
}

func (t *Tag) ListByArticleID(ids ...int) (res []entity.Tag) {
	res, err := service.Tag.ListByArticleIds(t.context(), ids)
	log.ErrorShortcut("template query error", err)
	return
}

func (t *Tag) ListByID(ids ...int) (res []entity.Tag) {
	res, err := service.Tag.ListByIds(t.context(), ids)
	log.ErrorShortcut("template query error", err)
	return
}

// PseudorandomList 伪随机列表
func (t *Tag) PseudorandomList() (res []entity.Tag) {
	if t.limit == 0 {
		t.limit = 20
	}
	res, err := service.Tag.PseudorandomList(t.context())
	log.ErrorShortcut("template query error", err)
	return
}
