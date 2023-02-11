package event

import (
	"moss/domain/core/entity"
)

type CategoryCreateBefore interface {
	CategoryCreateBefore(*entity.Category) error
}

type CategoryCreateAfter interface {
	CategoryCreateAfter(*entity.Category)
}

type CategoryUpdateBefore interface {
	CategoryUpdateBefore(item *entity.Category) error
}

type CategoryUpdateAfter interface {
	CategoryUpdateAfter(item *entity.Category)
}

type CategoryDeleteBefore interface {
	CategoryDeleteBefore(id int) error
}

type CategoryDeleteAfter interface {
	CategoryDeleteAfter(id int)
}

type CategoryGetAfter interface {
	CategoryGetAfter(*entity.Category)
}

type CategoryListAfter interface {
	CategoryListAfter([]entity.Category)
}
