package repository

import (
	"moss/domain/core/entity"
	"moss/domain/core/repository/context"
	"moss/domain/core/repository/gormx"
	"moss/infrastructure/persistent/db"
)

var Category = new(CategoryRepo)

type CategoryRepo struct {
}

func (r *CategoryRepo) MigrateTable() error {
	return db.DB.AutoMigrate(&entity.Category{})
}

func (r *CategoryRepo) Create(item *entity.Category) error {
	return db.DB.Create(item).Error
}

func (r *CategoryRepo) CreateInBatches(items []entity.Category, batchSize int) error {
	return db.DB.CreateInBatches(items, batchSize).Error
}

func (r *CategoryRepo) Update(item *entity.Category) error {
	return db.DB.Select("*").Omit("id").Where("id = ?", item.ID).Updates(item).Error
}

func (r *CategoryRepo) Delete(id int) error {
	return db.DB.Delete(&entity.Category{ID: id}).Error
}

func (r *CategoryRepo) Get(id int) (*entity.Category, error) {
	var res entity.Category
	err := db.DB.Where("id = ?", id).Find(&res).Error
	return &res, err
}

func (r *CategoryRepo) GetByName(name string) (*entity.Category, error) {
	var res entity.Category
	err := db.DB.Where("name = ?", name).Find(&res).Error
	return &res, err
}

func (r *CategoryRepo) GetBySlug(slug string) (*entity.Category, error) {
	var res entity.Category
	err := db.DB.Where("slug = ?", slug).Find(&res).Error
	return &res, err
}

func (r *CategoryRepo) GetIdByName(name string) (id int, err error) {
	err = db.DB.Model(entity.Category{}).Where("name = ?", name).Limit(1).Pluck("id", &id).Error
	return
}

func (r *CategoryRepo) GetIdBySlug(slug string) (id int, err error) {
	err = db.DB.Model(entity.Category{}).Where("slug = ?", slug).Limit(1).Pluck("id", &id).Error
	return
}

// MaxID 获取最大ID
func (r *CategoryRepo) MaxID() (res int, err error) {
	err = db.DB.Model(entity.Category{}).Limit(1).Order("id desc").Limit(1).Pluck("id", &res).Error
	return
}

// CountByWhere 通过where获取统计结果
func (r *CategoryRepo) CountByWhere(where *context.Where) (res int64, err error) {
	err = db.DB.Model(entity.Category{}).Scopes(gormx.Where(where)).Count(&res).Error
	return
}

// CountTotal 统计总数
func (r *CategoryRepo) CountTotal() (res int64, err error) {
	err = db.DB.Model(entity.Category{}).Count(&res).Error
	return
}

////////////////////////

// List 调用列表
func (r *CategoryRepo) List(ctx *context.Context) (res []entity.Category, err error) {
	err = db.DB.Model(entity.Category{}).Scopes(gormx.Context(ctx)).Find(&res).Error
	return
}

// ListByIds 通过ids获取列表
func (r *CategoryRepo) ListByIds(ctx *context.Context, ids []int) (res []entity.Category, err error) {
	err = db.DB.Model(entity.Category{}).Scopes(gormx.WhereIds(ids), gormx.Context(ctx)).Find(&res).Error
	return
}

// ListByParentIds 通过parentID获取列表
func (r *CategoryRepo) ListByParentIds(ctx *context.Context, ids []int) (res []entity.Category, err error) {
	err = db.DB.Model(entity.Category{}).Scopes(gormx.WhereParentIds(ids), gormx.Context(ctx)).Find(&res).Error
	return
}

// ListAfterCreateTime 根据创建时间调用列表
func (r *CategoryRepo) ListAfterCreateTime(ctx *context.Context, t int64) (res []entity.Category, err error) {
	err = db.DB.Model(entity.Category{}).Scopes(gormx.WhereCreateTimeAfter(t), gormx.Context(ctx)).Find(&res).Error
	return
}

// ListBeforeCreateTime 根据创建时间调用列表
func (r *CategoryRepo) ListBeforeCreateTime(ctx *context.Context, t int64) (res []entity.Category, err error) {
	err = db.DB.Model(entity.Category{}).Scopes(gormx.WhereCreateTimeBefore(t), gormx.Context(ctx)).Find(&res).Error
	return
}

// GetWithParent 获取分类和其夫分类
func (r *CategoryRepo) GetWithParent(id int) (res []entity.Category, err error) {
	var current entity.Category
	if err = db.DB.Where("id = ?", id).Find(&current).Error; err != nil || current.ID == 0 {
		return
	}
	res = append(res, current)
	if current.ParentID == 0 {
		return
	}
	var parent entity.Category
	if err = db.DB.Where("id = ?", current.ParentID).Find(&parent).Error; err != nil || parent.ID == 0 {
		return
	}
	res = append(res, parent)
	return
}
