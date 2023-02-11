package gormx

import (
	"gorm.io/gorm"
	"moss/domain/core/repository/context"
)

func Context(ctx *context.Context) ScopeType {
	if ctx == nil {
		return NothingScope
	}
	if ctx.Page == 0 {
		ctx.Page = 1
	}
	if ctx.FastOffset {
		return fastOffset(ctx)
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Scopes(Select(ctx.Select), base(ctx))
	}
}

func base(ctx *context.Context) ScopeType {
	return func(db *gorm.DB) *gorm.DB {
		return db.Scopes(
			Limit(ctx.Limit),
			Page(ctx.Page, ctx.Limit),
			Comment("select", ctx.Comment),
			Use(ctx.Where != nil, Where(ctx.Where)),
		).Order(ctx.Order)
	}
}

func fastOffset(ctx *context.Context) ScopeType {
	return func(db *gorm.DB) *gorm.DB {
		var ids []int
		// 复制一个会话，只查询ID列表
		// 注意1：外部必须带上Model,否则无法查询ID列表
		// 注意2：必须要把外部的其他查询条件放到Context之前，否则无法复制到之前的查询条件.
		// 举例：db.DB.Model(&entity.ArticleBase{}).Scopes(gormx.WhereCategoryIds(categoryIds), gormx.Context(ctx)).Find(&res).Error
		tx := db.Session(&gorm.Session{})
		if err := tx.Scopes(base(ctx)).Pluck("id", &ids).Error; err != nil {
			_ = db.AddError(err)
			return db
		}
		if len(ids) == 0 {
			_ = db.AddError(gorm.ErrRecordNotFound)
			return db
		}
		return db.Where("id in ?", ids).Scopes(Select(ctx.Select)).Order(ctx.Order)
	}
}
