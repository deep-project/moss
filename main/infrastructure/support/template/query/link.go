package query

import (
	"moss/domain/core/entity"
	"moss/domain/core/repository/context"
	"moss/domain/core/service"
	"moss/infrastructure/support/log"
)

type Link struct {
	limit   int
	order   string
	comment string
}

func NewLink() *Link {
	return &Link{}
}

func (l *Link) Limit(val int) *Link {
	l.limit = val
	return l
}

func (l *Link) Order(val string) *Link {
	l.order = val
	return l
}

func (l *Link) Comment(val string) *Link {
	l.comment = val
	return l
}

func (l *Link) context() *context.Context {
	if l.limit == 0 {
		l.limit = 20 // 强制限制数量
	}
	return context.NewContextWithComment(l.limit, l.order, l.comment)
}

// Get by id
func (l *Link) Get(id int) *entity.Link {
	res, err := service.Link.Get(id)
	log.WarnShortcut("template query error", err)
	return res
}

// List 调用文章列表
func (l *Link) List() (res []entity.Link) {
	res, err := service.Link.List(l.context())
	log.WarnShortcut("template query error", err)
	return
}

// ListByID 根据ID调用列表
func (l *Link) ListByID(ids ...int) (res []entity.Link) {
	res, err := service.Link.ListByIds(l.context(), ids)
	log.WarnShortcut("template query error", err)
	return
}

// ListPublic 公开的链接列表
func (l *Link) ListPublic() (res []entity.Link) {
	res, err := service.Link.ListPublic(l.context())
	log.WarnShortcut("template query error", err)
	return
}
