package plugins

import (
	"fmt"
	"go.uber.org/zap"
	appService "moss/application/service"
	"moss/domain/core/aggregate"
	"moss/domain/core/entity"
	"moss/domain/core/repository/context"
	"moss/domain/core/service"
	pluginEntity "moss/domain/support/entity"
	"moss/infrastructure/persistent/db"
)

type PostStore struct {
	Limit           int   `json:"limit"`
	Order           int   `json:"order"`             // 0:最新 1:最早 2:随机
	CategoryIds     []int `json:"category_ids"`      // 指定栏目ID
	DeleteOnFailure bool  `json:"delete_on_failure"` // 失败时删除
	ctx             *pluginEntity.Plugin
}

func NewPostStore() *PostStore {
	return &PostStore{
		Limit:       1,
		Order:       0,
		CategoryIds: []int{},
	}
}

func (p *PostStore) Info() *pluginEntity.PluginInfo {
	return &pluginEntity.PluginInfo{
		ID:         "PostStore",
		About:      "publish articles form storehouse",
		RunEnable:  true,
		CronEnable: true,
		PluginInfoPersistent: pluginEntity.PluginInfoPersistent{
			CronStart: false,
			CronExp:   "@every 1h",
		},
	}
}

func (p *PostStore) Load(ctx *pluginEntity.Plugin) error {
	return nil
}

func (p *PostStore) Run(ctx *pluginEntity.Plugin) (err error) {
	p.ctx = ctx

	if p.Limit <= 0 {
		p.ctx.Log.Warn("limit <= 0")
		return
	}

	//p.ctx.Log.Info("Begin...", zap.Int("limit", p.Limit), zap.Ints("category", p.CategoryIds))

	ids, err := p.ids()
	if err != nil {
		return
	}
	if len(ids) == 0 {
		if len(p.CategoryIds) > 0 {
			p.ctx.Log.Warn("NotFound Articles")
		}
		return
	}
	var success int
	for _, id := range ids {
		if item, err := p.post(id); err == nil {
			p.ctx.Log.Info(fmt.Sprintf("Post success, ID: %d Title: %s", item.ID, item.Title))
			success++
		}
	}
	p.ctx.Log.Info(fmt.Sprintf("End. success total: %d failures: %d", success, len(ids)-success))
	return nil
}

func (p *PostStore) post(id int) (item *aggregate.ArticlePost, err error) {
	if item, err = appService.StorePost(id); err != nil {
		p.ctx.Log.Error("post error", zap.Error(err), zap.Int("id", id))
		if p.DeleteOnFailure {
			if e := service.Store.Delete(id); e != nil {
				p.ctx.Log.Error("delete error", zap.Error(err), zap.Int("id", id))
			}
		}
	}
	return
}

func (p *PostStore) ids() (res []int, err error) {
	var items []entity.Store
	var ctx = context.Context{Select: "id", Limit: p.Limit, Order: p.order()}
	if len(p.CategoryIds) > 0 {
		items, err = service.Store.ListByCategoryIds(&ctx, p.CategoryIds)
	} else {
		items, err = service.Store.List(&ctx)
	}
	if err != nil {
		p.ctx.Log.Error("query list error", zap.Error(err))
		return
	}
	for _, item := range items {
		res = append(res, item.ID)
	}
	return
}

func (p *PostStore) order() any {
	switch p.Order {
	case 2:
		return db.RandomOrder()
	case 1:
		return "id asc"
	default:
		return "id desc"
	}
}
