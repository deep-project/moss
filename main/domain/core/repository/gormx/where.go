package gormx

import (
	"github.com/duke-git/lancet/v2/slice"
	"gorm.io/gorm"
	"moss/domain/core/repository/context"
	"regexp"
	"strings"
)

func Where(w *context.Where) ScopeType {
	return func(db *gorm.DB) *gorm.DB {
		reg, _ := regexp.Compile("[^a-zA-Z0-9_]+")
		w.Field = reg.ReplaceAllString(w.Field, "")
		if w.Field == "" {
			return db
		}
		switch w.Operator {
		case context.WhereOperatorEqual:
			return db.Where(w.Field+"=?", w.Value)
		case context.WhereOperatorEqualTrue:
			return db.Where(w.Field + "=true")
		case context.WhereOperatorEqualFalse:
			return db.Where(w.Field + "=false")
		case context.WhereOperatorEqualNull:
			return db.Where(w.Field + "=''")
		case context.WhereOperatorGreater:
			return db.Where(w.Field+"> ?", w.Value)
		case context.WhereOperatorGreaterEqual:
			return db.Where(w.Field+">= ?", w.Value)
		case context.WhereOperatorLess:
			return db.Where(w.Field+"< ?", w.Value)
		case context.WhereOperatorLessEqual:
			return db.Where(w.Field+"<= ?", w.Value)
		case context.WhereOperatorIn:
			return db.Where(w.Field+" in ?", strings.Split(w.Value, ","))
		case context.WhereOperatorInInt:
			return db.Where(w.Field+" in ?", slice.IntSlice(strings.Split(w.Value, ",")))
		case context.WhereOperatorLike:
			return db.Where(w.Field+" like ?", "%"+w.Value+"%")
		case context.WhereOperatorLikeBefore:
			return db.Where(w.Field+" like ?", w.Value+"%")
		case context.WhereOperatorLikeAfter:
			return db.Where(w.Field+" like ?", "%"+w.Value)
		default:
			return db.Where(w.Field+"=?", w.Value)
		}
	}
}
