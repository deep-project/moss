package repository

import (
	"gorm.io/gorm"
	"moss/domain/config"
	"moss/domain/core/entity"
	"moss/domain/core/repository/context"
	"moss/domain/core/repository/gormx"
	"moss/domain/core/utils"
	"moss/infrastructure/general/message"
	"moss/infrastructure/persistent/db"
)

var Store = new(StoreRepo)

type StoreRepo struct {
}

func (r *StoreRepo) MigrateTable() error {
	return db.DB.AutoMigrate(&entity.Store{})
}

func (r *StoreRepo) Create(item *entity.Store) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := r.checkPost(tx, item); err != nil {
			return err
		}
		return tx.Create(item).Error
	})
}

func (r *StoreRepo) Update(item *entity.Store) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := r.checkPost(tx, item); err != nil {
			return err
		}
		return tx.Select("*").Omit("id").Where("id = ?", item.ID).Updates(item).Error
	})
}

func (r *StoreRepo) checkPost(tx *gorm.DB, item *entity.Store) error {
	var id int
	// 判断 slug 是否存在
	if item.Slug != "" {
		if err := tx.Model(&entity.Store{}).Where("slug = ? and id != ?", item.Slug, item.ID).Limit(1).Pluck("id", &id).Error; err != nil {
			return err
		}
		if id > 0 {
			return message.ErrSlugAlreadyExists
		}
		// 检查 article base 表是否存在标题
		if err := tx.Model(&entity.ArticleBase{}).Where("slug = ?", item.Slug).Limit(1).Pluck("id", &id).Error; err != nil {
			return err
		}
		if id > 0 {
			return message.ErrSlugAlreadyExists
		}
	}
	// 判断 title 是否存在
	if config.Config.More.UniqueTitle {
		if err := tx.Model(&entity.Store{}).Where("title = ? and id != ?", item.Title, item.ID).Limit(1).Pluck("id", &id).Error; err != nil {
			return err
		}
		if id > 0 {
			return message.ErrTitleAlreadyExists
		}
		// 检查 article 表是否存在标题
		if err := tx.Model(&entity.Article{}).Where("title = ?", item.Title).Limit(1).Pluck("id", &id).Error; err != nil {
			return err
		}
		if id > 0 {
			return message.ErrTitleAlreadyExists
		}
	}
	return nil
}

func (r *StoreRepo) Delete(id int) error {
	return db.DB.Delete(&entity.Store{ID: id}).Error
}

func (r *StoreRepo) Get(id int) (*entity.Store, error) {
	var res entity.Store
	err := db.DB.Where("id = ?", id).Find(&res).Error
	return &res, err
}

// CountByWhere 通过where获取统计结果
func (r *StoreRepo) CountByWhere(where *context.Where) (res int64, err error) {
	err = db.DB.Model(&entity.Store{}).Scopes(gormx.Where(where)).Count(&res).Error
	return
}

// CountTotal 统计文章总数
func (r *StoreRepo) CountTotal() (res int64, err error) {
	err = db.DB.Model(entity.Store{}).Count(&res).Error
	return
}

// CountToday 统计今日添加数量
func (r *StoreRepo) CountToday() (res int64, err error) {
	err = db.DB.Model(entity.Store{}).Where("store_create_time > ?", utils.TodayBeginTime().Unix()).Count(&res).Error
	return
}

// CountYesterday 统计昨日添加数量
func (r *StoreRepo) CountYesterday() (res int64, err error) {
	today := utils.TodayBeginTime()
	yesterday := today.AddDate(0, 0, -1)
	err = db.DB.Model(entity.Store{}).Where("store_create_time > ? and store_create_time < ?", yesterday.Unix(), today.Unix()).Count(&res).Error
	return
}

// List 调用列表
func (r *StoreRepo) List(ctx *context.Context) (res []entity.Store, err error) {
	err = db.DB.Model(entity.Store{}).Scopes(gormx.Context(ctx)).Find(&res).Error
	return
}

// ListByCategoryIds 根据分类ID调用文章列表
func (r *StoreRepo) ListByCategoryIds(ctx *context.Context, categoryIds []int) (res []entity.Store, err error) {
	err = db.DB.Model(&entity.Store{}).Scopes(gormx.WhereCategoryIds(categoryIds), gormx.Context(ctx)).Find(&res).Error
	return
}
