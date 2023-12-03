package context

import "gorm.io/gorm"

type Context struct {
	Select     string `json:"select"`
	Limit      int    `json:"limit"`
	Order      any    `json:"order"`
	Page       int    `json:"page"`
	FastOffset bool   `json:"fast_offset"` // 是否快速分页（根据id集合再回表）
	Where      *Where `json:"where"`       // 查询条件where
	Comment    string `json:"comment"`
	Scope      []func(*gorm.DB) *gorm.DB
}

func NewContext(limit int, order string) *Context {
	return &Context{Limit: limit, Order: order}
}

func NewContextWithComment(limit int, order, comment string) *Context {
	return &Context{Limit: limit, Order: order, Comment: comment}
}

func NewContextByComment(comment string) *Context {
	return &Context{Comment: comment}
}

type Where struct {
	Field    string        `json:"field"`
	Operator WhereOperator `json:"operator"`
	Value    string        `json:"value"`
}

type WhereOperator = string

const (
	WhereOperatorEqual        = "equal"
	WhereOperatorEqualTrue    = "equalTrue"
	WhereOperatorEqualFalse   = "equalFalse"
	WhereOperatorEqualNull    = "equalNull"
	WhereOperatorGreater      = "greater"
	WhereOperatorGreaterEqual = "greaterEqual"
	WhereOperatorLess         = "less"
	WhereOperatorLessEqual    = "lessEqual"
	WhereOperatorIn           = "in"
	WhereOperatorInInt        = "inInt"
	WhereOperatorLike         = "like"
	WhereOperatorLikeBefore   = "likeBefore"
	WhereOperatorLikeAfter    = "likeAfter"
)
