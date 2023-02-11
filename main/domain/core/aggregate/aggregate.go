package aggregate

import (
	"moss/domain/core/entity"
)

type EntityInterface interface {
	entity.Article | entity.Category | entity.Tag | entity.Link
}

type CategoryTree struct {
	entity.Category
	Children []CategoryTree `json:"children"`
}
