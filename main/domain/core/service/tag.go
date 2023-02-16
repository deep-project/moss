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

var Tag = new(TagService)

type TagService struct {
	CreateBeforeEvents []event.TagCreateBefore
	CreateAfterEvents  []event.TagCreateAfter
	UpdateBeforeEvents []event.TagUpdateBefore
	UpdateAfterEvents  []event.TagUpdateAfter
	DeleteBeforeEvents []event.TagDeleteBefore
	DeleteAfterEvents  []event.TagDeleteAfter
	GetAfterEvents     []event.TagGetAfter
	ListAfterEvents    []event.TagListAfter
}

func (s *TagService) AddCreateBeforeEvents(ev ...event.TagCreateBefore) {
	s.CreateBeforeEvents = append(s.CreateBeforeEvents, ev...)
}
func (s *TagService) AddCreateAfterEvents(ev ...event.TagCreateAfter) {
	s.CreateAfterEvents = append(s.CreateAfterEvents, ev...)
}
func (s *TagService) AddUpdateBeforeEvents(ev ...event.TagUpdateBefore) {
	s.UpdateBeforeEvents = append(s.UpdateBeforeEvents, ev...)
}
func (s *TagService) AddUpdateAfterEvents(ev ...event.TagUpdateAfter) {
	s.UpdateAfterEvents = append(s.UpdateAfterEvents, ev...)
}
func (s *TagService) AddDeleteBeforeEvents(ev ...event.TagDeleteBefore) {
	s.DeleteBeforeEvents = append(s.DeleteBeforeEvents, ev...)
}
func (s *TagService) AddDeleteAfterEvents(ev ...event.TagDeleteAfter) {
	s.DeleteAfterEvents = append(s.DeleteAfterEvents, ev...)
}
func (s *TagService) AddGetAfterEvents(ev ...event.TagGetAfter) {
	s.GetAfterEvents = append(s.GetAfterEvents, ev...)
}
func (s *TagService) AddListAfterEvents(ev ...event.TagListAfter) {
	s.ListAfterEvents = append(s.ListAfterEvents, ev...)
}

// listAfterEvents
func (s *TagService) listAfterEvents(list []entity.Tag) {
	for _, e := range s.ListAfterEvents {
		e.TagListAfter(list)
	}
}

// getAfterEvents
func (s *TagService) getAfterEvents(item *entity.Tag) {
	for _, e := range s.GetAfterEvents {
		e.TagGetAfter(item)
	}
}

func (s *TagService) Save(item *entity.Tag) (err error) {
	if item.ID == 0 {
		return s.Create(item)
	}
	return s.Update(item)
}

func (s *TagService) Create(item *entity.Tag) (err error) {
	for _, e := range s.CreateBeforeEvents {
		if err = e.TagCreateBefore(item); err != nil {
			return
		}
	}
	if err = s.postCheck(item); err != nil {
		return
	}
	if item.CreateTime == 0 {
		item.CreateTime = time.Now().Unix()
	}
	if err = repository.Tag.Create(item); err != nil {
		return
	}
	for _, e := range s.CreateAfterEvents {
		e.TagCreateAfter(item)
	}
	return
}

// CreateInBatches 批量创建
func (s *TagService) CreateInBatches(items []entity.Tag, batchSize int) (err error) {
	for k := range items {
		for _, e := range s.CreateBeforeEvents {
			if err = e.TagCreateBefore(&items[k]); err != nil {
				return
			}
		}
		if err = s.postCheck(&items[k]); err != nil {
			return
		}
		if items[k].CreateTime == 0 {
			items[k].CreateTime = time.Now().Unix()
		}
	}
	if err = repository.Tag.CreateInBatches(items, batchSize); err != nil {
		return
	}
	for _, item := range items {
		for _, e := range s.CreateAfterEvents {
			e.TagCreateAfter(&item)
		}
	}
	return
}

func (s *TagService) Update(item *entity.Tag) (err error) {
	if item.ID == 0 {
		return message.ErrIdRequired
	}
	for _, e := range s.UpdateBeforeEvents {
		if err = e.TagUpdateBefore(item); err != nil {
			return
		}
	}
	if err = s.postCheck(item); err != nil {
		return
	}
	if err = repository.Tag.Update(item); err != nil {
		return
	}
	for _, e := range s.UpdateAfterEvents {
		e.TagUpdateAfter(item)
	}
	return
}

func (s *TagService) postCheck(item *entity.Tag) error {
	if item.Slug == "" {
		return message.ErrSlugRequired
	}
	if item.Name == "" {
		return message.ErrNameRequired
	}
	if strings.HasPrefix(item.Slug, " ") {
		return message.ErrSlugStartSpaceRequired
	}
	if strings.HasSuffix(item.Slug, " ") {
		return message.ErrSlugEndSpaceRequired
	}
	return nil
}

func (s *TagService) Delete(id int) (err error) {
	if id == 0 {
		return message.ErrIdRequired
	}
	for _, e := range s.DeleteBeforeEvents {
		if err = e.TagDeleteBefore(id); err != nil {
			return
		}
	}
	if err = repository.Tag.Delete(id); err != nil {
		return
	}
	for _, e := range s.DeleteAfterEvents {
		e.TagDeleteAfter(id)
	}
	return
}

func (s *TagService) Get(id int) (res *entity.Tag, err error) {
	if id == 0 {
		return nil, message.ErrSlugRequired
	}
	if res, err = repository.Tag.Get(id); err != nil {
		return
	}
	if res.ID == 0 {
		return nil, message.ErrRecordNotFound
	}
	for _, e := range s.GetAfterEvents {
		e.TagGetAfter(res)
	}
	return
}

func (s *TagService) GetBySlug(slug string) (res *entity.Tag, err error) {
	if slug == "" {
		return nil, message.ErrIdRequired
	}
	if res, err = repository.Tag.GetBySlug(slug); err != nil {
		return
	}
	if res.ID == 0 {
		return nil, message.ErrRecordNotFound
	}
	for _, e := range s.GetAfterEvents {
		e.TagGetAfter(res)
	}
	return
}

func (s *TagService) ExistsSlug(slug string) (bool, error) {
	id, err := s.GetIdBySlug(slug)
	return id > 0, err
}

func (s *TagService) ExistsName(name string) (bool, error) {
	id, err := s.GetIdByName(name)
	return id > 0, err
}

// GetIdByName 通过name获取ID
func (s *TagService) GetIdByName(name string) (id int, err error) {
	if name == "" {
		return 0, message.ErrNameRequired
	}
	return repository.Tag.GetIdByName(name)
}

// GetIdBySlug 通过slug获取ID
func (s *TagService) GetIdBySlug(slug string) (id int, err error) {
	if slug == "" {
		return 0, message.ErrSlugRequired
	}
	return repository.Tag.GetIdBySlug(slug)
}

// GetIdByNameOrCreate 通过name获取主键,不存在则创建
func (s *TagService) GetIdByNameOrCreate(name string) (id int, err error) {
	if id, err = s.GetIdByName(name); err != nil {
		return
	}
	if id == 0 {
		id, err = s.CreateByNameReturnID(name)
	}
	return
}

func (s *TagService) CreateByName(name string) (res *entity.Tag, err error) {
	res = &entity.Tag{Name: name}
	err = s.Create(res)
	return
}

func (s *TagService) CreateByNameReturnID(name string) (int, error) {
	item, err := s.CreateByName(name)
	return item.ID, err
}

func (s *TagService) CountByWhere(where *context.Where) (res int64, err error) {
	return repository.Tag.CountByWhere(where)
}

// CountTotal 统计总数
func (s *TagService) CountTotal() (int64, error) {
	return repository.Tag.CountTotal()
}

///////////////////////////////

func (s *TagService) List(ctx *context.Context) (res []entity.Tag, err error) {
	res, err = repository.Tag.List(ctx)
	s.listAfterEvents(res)
	return
}

// ListByIds 通过tagID获取列表
func (s *TagService) ListByIds(ctx *context.Context, ids []int) (res []entity.Tag, err error) {
	if len(ids) == 0 {
		return
	}
	res, err = repository.Tag.ListByIds(ctx, ids)
	s.listAfterEvents(res)
	return
}

// ListByArticleID 通过文章ID获取列表
func (s *TagService) ListByArticleID(ctx *context.Context, id int) (res []entity.Tag, err error) {
	return s.ListByArticleIds(ctx, []int{id})
}

// ListByArticleIds 通过文章ID获取列表
func (s *TagService) ListByArticleIds(ctx *context.Context, articleIds []int) (res []entity.Tag, err error) {
	ids, err := Mapping.ListTagIdByArticleIds(ctx, articleIds)
	if err != nil {
		return
	}
	res, err = s.ListByIds(nil, ids)
	return
}

// ListAfterCreateTime 根据创建时间调用列表
func (s *TagService) ListAfterCreateTime(ctx *context.Context, t int64) (res []entity.Tag, err error) {
	res, err = repository.Tag.ListAfterCreateTime(ctx, t)
	s.listAfterEvents(res)
	return
}

// PseudorandomList 伪随机列表
func (s *TagService) PseudorandomList(ctx *context.Context) (res []entity.Tag, err error) {
	maxID, err := repository.Tag.MaxID()
	if err != nil {
		return
	}
	return s.ListByIds(ctx, pseudorandomIds(maxID, ctx.Limit))
}
