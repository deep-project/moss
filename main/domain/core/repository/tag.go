package repository

import (
	"moss/domain/core/entity"
	"moss/domain/core/repository/context"
	"moss/domain/core/repository/gormx"
	"moss/infrastructure/persistent/db"
)

func init() {

}

var Tag = new(TagRepo)

type TagRepo struct {
}

func (r *TagRepo) MigrateTable() error {
	return db.DB.AutoMigrate(&entity.Tag{})
}

func (r *TagRepo) Create(item *entity.Tag) error {
	return db.DB.Create(item).Error
}

func (r *TagRepo) CreateInBatches(items []entity.Tag, batchSize int) error {
	return db.DB.CreateInBatches(items, batchSize).Error
}

func (r *TagRepo) Update(item *entity.Tag) error {
	return db.DB.Select("*").Omit("id").Where("id = ?", item.ID).Updates(item).Error
}

func (r *TagRepo) Delete(id int) error {
	return db.DB.Delete(&entity.Tag{ID: id}).Error
}

func (r *TagRepo) Get(id int) (*entity.Tag, error) {
	var res entity.Tag
	err := db.DB.Where("id = ?", id).Find(&res).Error
	return &res, err
}

func (r *TagRepo) GetBySlug(slug string) (*entity.Tag, error) {
	var res entity.Tag
	err := db.DB.Where("slug = ?", slug).Find(&res).Error
	return &res, err
}

// GetIdByName 通过name获取ID
func (r *TagRepo) GetIdByName(name string) (id int, err error) {
	err = db.DB.Model(entity.Tag{}).Where("name = ?", name).Limit(1).Pluck("id", &id).Error
	return
}

// GetIdBySlug 通过slug获取ID
func (r *TagRepo) GetIdBySlug(name string) (id int, err error) {
	err = db.DB.Model(entity.Tag{}).Where("slug = ?", name).Limit(1).Pluck("id", &id).Error
	return
}

// CountByWhere 通过where获取统计结果
func (r *TagRepo) CountByWhere(where *context.Where) (res int64, err error) {
	err = db.DB.Model(entity.Tag{}).Scopes(gormx.Where(where)).Count(&res).Error
	return
}

// CountTotal 统计总数
func (r *TagRepo) CountTotal() (res int64, err error) {
	err = db.DB.Model(entity.Tag{}).Count(&res).Error
	return
}

// MaxID 获取最大ID
func (r *TagRepo) MaxID() (res int, err error) {
	err = db.DB.Model(entity.Tag{}).Model(entity.Tag{}).Limit(1).Order("id desc").Limit(1).Pluck("id", &res).Error
	return
}

///////////////////////////////////////////////

// List 调用列表
func (r *TagRepo) List(ctx *context.Context) (res []entity.Tag, err error) {
	err = db.DB.Model(entity.Tag{}).Scopes(gormx.Context(ctx)).Find(&res).Error
	return
}

// ListByIds 通过主键获取列表
func (r *TagRepo) ListByIds(ctx *context.Context, ids []int) (res []entity.Tag, err error) {
	if ctx == nil {
		ctx = &context.Context{}
	}
	err = db.DB.Model(entity.Tag{}).Scopes(gormx.Context(ctx, gormx.WhereIds(ids))).Find(&res).Error
	return
}

// ListAfterCreateTime 根据创建时间调用列表
func (r *TagRepo) ListAfterCreateTime(ctx *context.Context, t int64) (res []entity.Tag, err error) {
	err = db.DB.Model(entity.Tag{}).Scopes(gormx.Context(ctx, gormx.WhereCreateTimeAfter(t))).Find(&res).Error
	return
}

// ListBeforeCreateTime 根据创建时间调用列表
func (r *TagRepo) ListBeforeCreateTime(ctx *context.Context, t int64) (res []entity.Tag, err error) {
	err = db.DB.Model(entity.Tag{}).Scopes(gormx.Context(ctx, gormx.WhereCreateTimeBefore(t))).Find(&res).Error
	return
}
