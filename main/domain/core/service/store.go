package service

import (
	"moss/domain/core/entity"
	"moss/domain/core/repository"
	"moss/domain/core/repository/context"
	"moss/infrastructure/general/message"
	"strings"
	"time"
)

var Store = new(StoreService)

type StoreService struct {
}

func (s *StoreService) Create(item *entity.Store) (err error) {
	if err = s.postCheck(item); err != nil {
		return
	}
	if item.StoreCreateTime == 0 {
		item.StoreCreateTime = time.Now().Unix()
	}
	return repository.Store.Create(item)
}

func (s *StoreService) Update(item *entity.Store) (err error) {
	if item.ID == 0 {
		return message.ErrIdRequired
	}
	if err = s.postCheck(item); err != nil {
		return
	}
	return repository.Store.Update(item)
}

func (s *StoreService) postCheck(item *entity.Store) error {
	if item.Title == "" {
		return message.ErrTitleRequired
	}
	if strings.HasPrefix(item.Slug, " ") {
		return message.ErrSlugStartSpaceRequired
	}
	if strings.HasSuffix(item.Slug, " ") {
		return message.ErrSlugEndSpaceRequired
	}
	return nil
}

func (s *StoreService) Delete(id int) (err error) {
	if id == 0 {
		return message.ErrIdRequired
	}
	return repository.Store.Delete(id)
}

func (s *StoreService) BatchDelete(ids []int) (err error) {
	for _, id := range ids {
		if err = s.Delete(id); err != nil {
			return
		}
	}
	return
}

func (s *StoreService) Get(id int) (res *entity.Store, err error) {
	if id == 0 {
		return nil, message.ErrIdRequired
	}
	if res, err = repository.Store.Get(id); err != nil {
		return
	}
	if res.ID == 0 {
		return nil, message.ErrRecordNotFound
	}
	return
}

func (s *StoreService) CountByWhere(where *context.Where) (res int64, err error) {
	return repository.Store.CountByWhere(where)
}

// CountTotal 统计文章总数
func (s *StoreService) CountTotal() (int64, error) {
	return repository.Store.CountTotal()
}

// CountToday 统计今日添加数量
func (s *StoreService) CountToday() (int64, error) {
	return repository.Store.CountToday()
}

// CountYesterday 统计昨日添加数量
func (s *StoreService) CountYesterday() (int64, error) {
	return repository.Store.CountYesterday()
}

func (s *StoreService) List(ctx *context.Context) (res []entity.Store, err error) {
	res, err = repository.Store.List(ctx)
	return
}

func (s *StoreService) ListByCategoryID(ctx *context.Context, id int) (res []entity.Store, err error) {
	return s.ListByCategoryIds(ctx, []int{id})
}

// ListByCategoryIds 通过分类ID调用文章列表
func (s *StoreService) ListByCategoryIds(ctx *context.Context, ids []int) (res []entity.Store, err error) {
	if len(ids) == 0 {
		return
	}
	res, err = repository.Store.ListByCategoryIds(ctx, ids)
	return
}
