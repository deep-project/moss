package service

import (
	"moss/domain/core/entity"
	"moss/domain/core/event"
	"moss/domain/core/repository"
	"moss/domain/core/repository/context"
	"moss/infrastructure/general/message"
	"strings"
	"time"
)

var Link = new(LinkService)

type LinkService struct {
	CreateBeforeEvents []event.LinkCreateBefore
	CreateAfterEvents  []event.LinkCreateAfter
	UpdateBeforeEvents []event.LinkUpdateBefore
	UpdateAfterEvents  []event.LinkUpdateAfter
	DeleteBeforeEvents []event.LinkDeleteBefore
	DeleteAfterEvents  []event.LinkDeleteAfter
	GetAfterEvents     []event.LinkGetAfter
	ListAfterEvents    []event.LinkListAfter
}

func (s *LinkService) AddCreateBeforeEvents(ev ...event.LinkCreateBefore) {
	s.CreateBeforeEvents = append(s.CreateBeforeEvents, ev...)
}
func (s *LinkService) AddCreateAfterEvents(ev ...event.LinkCreateAfter) {
	s.CreateAfterEvents = append(s.CreateAfterEvents, ev...)
}
func (s *LinkService) AddUpdateBeforeEvents(ev ...event.LinkUpdateBefore) {
	s.UpdateBeforeEvents = append(s.UpdateBeforeEvents, ev...)
}
func (s *LinkService) AddUpdateAfterEvents(ev ...event.LinkUpdateAfter) {
	s.UpdateAfterEvents = append(s.UpdateAfterEvents, ev...)
}
func (s *LinkService) AddDeleteBeforeEvents(ev ...event.LinkDeleteBefore) {
	s.DeleteBeforeEvents = append(s.DeleteBeforeEvents, ev...)
}
func (s *LinkService) AddDeleteAfterEvents(ev ...event.LinkDeleteAfter) {
	s.DeleteAfterEvents = append(s.DeleteAfterEvents, ev...)
}
func (s *LinkService) AddGetAfterEvents(ev ...event.LinkGetAfter) {
	s.GetAfterEvents = append(s.GetAfterEvents, ev...)
}
func (s *LinkService) AddListAfterEvents(ev ...event.LinkListAfter) {
	s.ListAfterEvents = append(s.ListAfterEvents, ev...)
}

// listAfterEvents
func (s *LinkService) listAfterEvents(list []entity.Link) {
	for _, e := range s.ListAfterEvents {
		e.LinkListAfter(list)
	}
}

func (s *LinkService) Save(item *entity.Link) (err error) {
	if item.ID == 0 {
		return s.Create(item)
	}
	return s.Update(item)
}

func (s *LinkService) Create(item *entity.Link) (err error) {
	for _, e := range s.CreateBeforeEvents {
		if err = e.LinkCreateBefore(item); err != nil {
			return
		}
	}
	if err = s.postCheck(item); err != nil {
		return
	}
	if item.CreateTime == 0 {
		item.CreateTime = time.Now().Unix()
	}
	if err = repository.Link.Create(item); err != nil {
		return
	}
	for _, e := range s.CreateAfterEvents {
		e.LinkCreateAfter(item)
	}
	return
}

func (s *LinkService) Update(item *entity.Link) (err error) {
	if item.ID == 0 {
		return message.ErrIdRequired
	}
	for _, e := range s.UpdateBeforeEvents {
		if err = e.LinkUpdateBefore(item); err != nil {
			return
		}
	}
	if err = s.postCheck(item); err != nil {
		return
	}
	if err = repository.Link.Update(item); err != nil {
		return
	}
	for _, e := range s.UpdateAfterEvents {
		e.LinkUpdateAfter(item)
	}
	return
}

func (s *LinkService) postCheck(item *entity.Link) error {
	if item.Name == "" {
		return message.ErrNameRequired
	}
	if item.URL == "" {
		return message.ErrUrlRequired
	}
	return nil
}

func (s *LinkService) Delete(id int) (err error) {
	if id == 0 {
		return message.ErrIdRequired
	}
	for _, e := range s.DeleteBeforeEvents {
		if err = e.LinkDeleteBefore(id); err != nil {
			return
		}
	}
	if err = repository.Link.Delete(id); err != nil {
		return
	}
	for _, e := range s.DeleteAfterEvents {
		e.LinkDeleteAfter(id)
	}
	return
}

func (s *LinkService) BatchDelete(ids []int) (err error) {
	for _, id := range ids {
		if err = s.Delete(id); err != nil {
			return
		}
	}
	return
}

func (s *LinkService) Get(id int) (res *entity.Link, err error) {
	if id == 0 {
		return nil, message.ErrIdRequired
	}
	if res, err = repository.Link.Get(id); err != nil {
		return
	}
	if res.ID == 0 {
		return nil, message.ErrRecordNotFound
	}
	for _, e := range s.GetAfterEvents {
		e.LinkGetAfter(res)
	}
	return
}

func (s *LinkService) ExistsURL(url string) (bool, error) {
	if url == "" {
		return false, message.ErrUrlRequired
	}
	id, err := repository.Link.GetIdByURL(url)
	return id > 0, err
}

//////////////////

func (s *LinkService) CountByWhere(where *context.Where) (res int64, err error) {
	return repository.Link.CountByWhere(where)
}

// CountTotal 统计总数
func (s *LinkService) CountTotal() (int64, error) {
	return repository.Link.CountTotal()
}

func (s *LinkService) DisableLink(id int) error {
	return repository.Link.DisableLink(id)
}

func (s *LinkService) EnableLink(id int) error {
	return repository.Link.EnableLink(id)
}

func (s *LinkService) List(ctx *context.Context) (res []entity.Link, err error) {
	res, err = repository.Link.List(ctx)
	s.listAfterEvents(res)
	return
}

// ListByIds 根据id调用列表
func (s *LinkService) ListByIds(ctx *context.Context, ids []int) (res []entity.Link, err error) {
	if len(ids) == 0 {
		return
	}
	res, err = repository.Link.ListByIds(ctx, ids)
	s.listAfterEvents(res)
	return
}

// ListPublic 前台公开列表
func (s *LinkService) ListPublic(ctx *context.Context) (res []entity.Link, err error) {
	if res, err = repository.Link.ListPublic(ctx); err != nil {
		return
	}
	s.listAfterEvents(res)
	return
}

// ListLikeURL 相似链接列表
func (s *LinkService) ListLikeURL(ctx *context.Context, url string) (res []entity.Link, err error) {
	if url == "" {
		return nil, message.ErrUrlRequired
	}
	url = strings.TrimPrefix(url, "http://")
	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimPrefix(url, "//")
	return repository.Link.ListLikeURL(ctx, url)
}

// ListDetectLink 开启检查的链接列表
func (s *LinkService) ListDetectLink(ctx *context.Context) (res []entity.Link, err error) {
	list, err := repository.Link.ListDetectLink(ctx)
	if err != nil {
		return
	}
	// 排除延迟检测
	now := time.Now().Unix()
	for _, item := range list {
		if item.DetectDelay == 0 || now >= item.CreateTime+item.DetectDelay*60 {
			res = append(res, item)
		}
	}
	s.listAfterEvents(res)
	return
}
