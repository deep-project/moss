package gormx

import (
	"gorm.io/gorm"
	"gorm.io/hints"
)

type ScopeType = func(*gorm.DB) *gorm.DB

var NothingScope = func(db *gorm.DB) *gorm.DB {
	return db
}

func Use(val bool, items ...ScopeType) ScopeType {
	return func(db *gorm.DB) *gorm.DB {
		if val {
			return db.Scopes(items...)
		}
		return db
	}
}

func UseOr(val bool, fn1 ScopeType, fn2 ScopeType) ScopeType {
	return func(db *gorm.DB) *gorm.DB {
		if val {
			return db.Scopes(fn1)
		}
		return db.Scopes(fn2)
	}
}

func Select(val string) ScopeType {
	return func(db *gorm.DB) *gorm.DB {
		if val != "" {
			db.Select(val)
		}
		return db
	}
}

func Limit(val int) ScopeType {
	return func(db *gorm.DB) *gorm.DB {
		if val > 0 {
			db.Limit(val)
		}
		return db
	}
}

func Page(page, limit int) ScopeType {
	return func(db *gorm.DB) *gorm.DB {
		if page > 1 && limit > 0 {
			db.Offset((page - 1) * limit)
		}
		return db
	}
}

func Comment(clause, val string) ScopeType {
	return func(db *gorm.DB) *gorm.DB {
		if val != "" {
			db.Clauses(hints.CommentBefore(clause, val))
		}
		return db
	}
}

func WhereIds(ids []int) ScopeType {
	return whereFieldContainInt("id", ids)
}

func WhereCategoryIds(ids []int) ScopeType {
	return whereFieldContainInt("category_id", ids)
}

func WhereParentIds(ids []int) ScopeType {
	return whereFieldContainInt("parent_id", ids)
}

func WhereCategoryID(id int) ScopeType {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("category_id = ?", id)
	}
}

func WhereCreateTimeAfter(t int64) ScopeType {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("create_time > ?", t)
	}
}

func WhereCreateTimeBefore(t int64) ScopeType {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("create_time < ?", t)
	}
}

// 根据val长度生成不同的查询条件
func whereFieldContainInt(field string, val []int) ScopeType {
	return func(db *gorm.DB) *gorm.DB {
		if len(val) == 1 {
			return db.Where(field+" = ?", val[0])
		}
		return db.Where(field+" in ?", val)
	}
}
