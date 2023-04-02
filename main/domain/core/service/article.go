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

var Article = new(ArticleService)

type ArticleService struct {
	CreateBeforeEvents []event.ArticleCreateBefore
	CreateAfterEvents  []event.ArticleCreateAfter
	UpdateBeforeEvents []event.ArticleUpdateBefore
	UpdateAfterEvents  []event.ArticleUpdateAfter
	DeleteBeforeEvents []event.ArticleDeleteBefore
	DeleteAfterEvents  []event.ArticleDeleteAfter
	GetAfterEvents     []event.ArticleGetAfter
	ListAfterEvents    []event.ArticleListAfter
}

func (s *ArticleService) AddCreateBeforeEvents(ev ...event.ArticleCreateBefore) {
	s.CreateBeforeEvents = append(s.CreateBeforeEvents, ev...)
}
func (s *ArticleService) AddCreateAfterEvents(ev ...event.ArticleCreateAfter) {
	s.CreateAfterEvents = append(s.CreateAfterEvents, ev...)
}
func (s *ArticleService) AddUpdateBeforeEvents(ev ...event.ArticleUpdateBefore) {
	s.UpdateBeforeEvents = append(s.UpdateBeforeEvents, ev...)
}
func (s *ArticleService) AddUpdateAfterEvents(ev ...event.ArticleUpdateAfter) {
	s.UpdateAfterEvents = append(s.UpdateAfterEvents, ev...)
}
func (s *ArticleService) AddDeleteBeforeEvents(ev ...event.ArticleDeleteBefore) {
	s.DeleteBeforeEvents = append(s.DeleteBeforeEvents, ev...)
}
func (s *ArticleService) AddDeleteAfterEvents(ev ...event.ArticleDeleteAfter) {
	s.DeleteAfterEvents = append(s.DeleteAfterEvents, ev...)
}
func (s *ArticleService) AddGetAfterEvents(ev ...event.ArticleGetAfter) {
	s.GetAfterEvents = append(s.GetAfterEvents, ev...)
}
func (s *ArticleService) AddListAfterEvents(ev ...event.ArticleListAfter) {
	s.ListAfterEvents = append(s.ListAfterEvents, ev...)
}

// listAfterEvents
func (s *ArticleService) listAfterEvents(list []entity.ArticleBase) {
	for _, e := range s.ListAfterEvents {
		e.ArticleListAfter(list)
	}
}

// getAfterEvents
func (s *ArticleService) getAfterEvents(item *entity.Article) {
	for _, e := range s.GetAfterEvents {
		e.ArticleGetAfter(item)
	}
}

func (s *ArticleService) Get(id int) (res *entity.Article, err error) {
	if id == 0 {
		return nil, message.ErrIdRequired
	}
	if res, err = repository.Article.Get(id); err != nil {
		return
	}
	if res.ID == 0 {
		return nil, message.ErrRecordNotFound
	}
	s.getAfterEvents(res)
	return
}

func (s *ArticleService) GetBySlug(slug string) (res *entity.Article, err error) {
	if slug == "" {
		return nil, message.ErrIdRequired
	}
	if res, err = repository.Article.GetBySlug(slug); err != nil {
		return
	}
	if res.ID == 0 {
		return nil, message.ErrRecordNotFound
	}
	s.getAfterEvents(res)
	return
}

func (s *ArticleService) Save(item *entity.Article) (err error) {
	if item.ID == 0 {
		return s.Create(item)
	}
	return s.Update(item)
}

func (s *ArticleService) Create(item *entity.Article) (err error) {
	for _, e := range s.CreateBeforeEvents {
		if err = e.ArticleCreateBefore(item); err != nil {
			return
		}
	}
	if err = s.postCheck(item); err != nil {
		return
	}
	if item.CreateTime == 0 {
		item.CreateTime = time.Now().Unix()
	}
	if err = repository.Article.Create(item); err != nil {
		return
	}
	for _, e := range s.CreateAfterEvents {
		e.ArticleCreateAfter(item)
	}
	return
}

// CreateInBatches 批量创建
func (s *ArticleService) CreateInBatches(items []entity.Article) (err error) {
	for k := range items {
		for _, e := range s.CreateBeforeEvents {
			if err = e.ArticleCreateBefore(&items[k]); err != nil {
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
	if err = repository.Article.CreateInBatches(items); err != nil {
		return
	}
	for _, item := range items {
		for _, e := range s.CreateAfterEvents {
			e.ArticleCreateAfter(&item)
		}
	}
	return
}

func (s *ArticleService) Update(item *entity.Article) (err error) {
	if item.ID == 0 {
		return message.ErrIdRequired
	}
	for _, e := range s.UpdateBeforeEvents {
		if err = e.ArticleUpdateBefore(item); err != nil {
			return
		}
	}
	if err = s.postCheck(item); err != nil {
		return
	}
	if err = repository.Article.Update(item); err != nil {
		return
	}
	for _, e := range s.UpdateAfterEvents {
		e.ArticleUpdateAfter(item)
	}
	return
}

func (s *ArticleService) postCheck(item *entity.Article) error {
	if item.Slug == "" {
		return message.ErrSlugRequired
	}
	if item.Title == "" {
		return message.ErrTitleRequired
	}
	if item.Content == "" {
		return message.ErrContentRequired
	}
	if strings.HasPrefix(item.Slug, " ") {
		return message.ErrSlugStartSpaceRequired
	}
	if strings.HasSuffix(item.Slug, " ") {
		return message.ErrSlugEndSpaceRequired
	}
	return nil
}

func (s *ArticleService) Delete(id int) (err error) {
	if id == 0 {
		return message.ErrIdRequired
	}
	for _, e := range s.DeleteBeforeEvents {
		if err = e.ArticleDeleteBefore(id); err != nil {
			return
		}
	}
	if err = repository.Article.Delete(id); err != nil {
		return
	}
	for _, e := range s.DeleteAfterEvents {
		e.ArticleDeleteAfter(id)
	}
	return
}

func (s *ArticleService) ExistsSlug(slug string) (bool, error) {
	if slug == "" {
		return false, message.ErrSlugRequired
	}
	id, err := repository.Article.GetIdBySlug(slug)
	return id > 0, err
}

func (s *ArticleService) ExistsTitle(title string) (bool, error) {
	if title == "" {
		return false, message.ErrTitleRequired
	}
	id, err := repository.Article.GetIdByTitle(title)
	return id > 0, err
}

func (s *ArticleService) UpdateViewsBySlug(id string, n int) error {
	return repository.Article.UpdateViewsBySlug(id, n)
}

//////-------list ------

// List 调用文章列表
func (s *ArticleService) List(ctx *context.Context) (res []entity.ArticleBase, err error) {
	res, err = repository.Article.List(ctx)
	s.listAfterEvents(res)
	return
}

// ListExistThumbnail 调用有缩略图文章列表
func (s *ArticleService) ListExistThumbnail(ctx *context.Context) (res []entity.ArticleBase, err error) {
	res, err = repository.Article.ListExistThumbnail(ctx)
	s.listAfterEvents(res)
	return
}

// ListByIds 根据id调用文章列表
func (s *ArticleService) ListByIds(ctx *context.Context, ids []int) (res []entity.ArticleBase, err error) {
	if len(ids) == 0 {
		return
	}
	res, err = repository.Article.ListByIds(ctx, ids)
	s.listAfterEvents(res)
	return
}

func (s *ArticleService) ListByCategoryID(ctx *context.Context, id int) (res []entity.ArticleBase, err error) {
	return s.ListByCategoryIds(ctx, []int{id})
}

// ListByCategoryIds 通过分类ID调用文章列表
func (s *ArticleService) ListByCategoryIds(ctx *context.Context, ids []int) (res []entity.ArticleBase, err error) {
	if len(ids) == 0 {
		return
	}
	res, err = repository.Article.ListByCategoryIds(ctx, ids)
	s.listAfterEvents(res)
	return
}

// ListAfterCreateTime 根据创建时间调用列表
func (s *ArticleService) ListAfterCreateTime(ctx *context.Context, t int64) (res []entity.ArticleBase, err error) {
	res, err = repository.Article.ListAfterCreateTime(ctx, t)
	s.listAfterEvents(res)
	return
}

func (s *ArticleService) ListByTagID(ctx *context.Context, tagID int) (res []entity.ArticleBase, err error) {
	return s.ListByTagIds(ctx, []int{tagID})
}

// ListByTagIds 根据标签ID调用文章列表
func (s *ArticleService) ListByTagIds(ctx *context.Context, tagIds []int) (res []entity.ArticleBase, err error) {
	ids, err := Mapping.ListArticleIdsByTagIds(ctx, tagIds)
	if err != nil {
		return
	}
	res, err = s.ListByIds(context.NewContext(ctx.Limit, "id desc"), ids)
	return
}

// PseudorandomList 伪随机列表
func (s *ArticleService) PseudorandomList(ctx *context.Context) (res []entity.ArticleBase, err error) {
	maxID, err := repository.Article.MaxID()
	if err != nil {
		return
	}
	return s.ListByIds(ctx, pseudorandomIds(maxID, ctx.Limit))
}

// ListDetail 调用详情表文章列表
func (s *ArticleService) ListDetail(ctx *context.Context) (res []entity.ArticleDetail, err error) {
	res, err = repository.Article.ListDetail(ctx)
	return
}

func (s *ArticleService) ListDetailByIds(ctx *context.Context, ids []int) (res []entity.ArticleDetail, err error) {
	if len(ids) == 0 {
		return
	}
	res, err = repository.Article.ListDetailByIds(ctx, ids)
	return
}

// CountByCategoryID 根据分类ID统计文章数
func (s *ArticleService) CountByCategoryID(categoryID int) (int64, error) {
	return repository.Article.CountByCategoryID(categoryID)
}

func (s *ArticleService) CountByWhere(where *context.Where) (res int64, err error) {
	return repository.Article.CountByWhere(where)
}

// CountTotal 统计文章总数
func (s *ArticleService) CountTotal() (int64, error) {
	return repository.Article.CountTotal()
}

// CountToday 统计今日添加数量
func (s *ArticleService) CountToday() (int64, error) {
	return repository.Article.CountToday()
}

// CountYesterday 统计昨日添加数量
func (s *ArticleService) CountYesterday() (int64, error) {
	return repository.Article.CountYesterday()
}

// CountLastFewDays 统计最近几日的
func (s *ArticleService) CountLastFewDays(n int) (int64, error) {
	return repository.Article.CountLastFewDays(n)
}

func (s *ArticleService) MergeBaseListAndDetailList(v1 []entity.ArticleBase, v2 []entity.ArticleDetail) (res []entity.Article) {
	for _, v := range v1 {
		detail, found := s.FindDetailListByID(v2, v.ID)
		if !found {
			continue
		}
		res = append(res, entity.Article{ArticleBase: v, ArticleDetail: detail})
	}
	return
}

func (s *ArticleService) FindDetailListByID(list []entity.ArticleDetail, id int) (res entity.ArticleDetail, found bool) {
	for _, v := range list {
		if v.ArticleID == id {
			return v, true
		}
	}
	return res, false
}

// BatchSetCategory 批量设置分类
func (s *ArticleService) BatchSetCategory(categoryID int, ids []int) error {
	return repository.Article.BatchSetCategory(categoryID, ids)
}
