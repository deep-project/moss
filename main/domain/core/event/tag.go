package event

import (
	"moss/domain/core/entity"
)

type TagCreateBefore interface {
	TagCreateBefore(*entity.Tag) error
}

type TagCreateAfter interface {
	TagCreateAfter(*entity.Tag)
}

type TagUpdateBefore interface {
	TagUpdateBefore(item *entity.Tag) error
}

type TagUpdateAfter interface {
	TagUpdateAfter(item *entity.Tag)
}

type TagDeleteBefore interface {
	TagDeleteBefore(id int) error
}

type TagDeleteAfter interface {
	TagDeleteAfter(id int)
}

type TagGetAfter interface {
	TagGetAfter(*entity.Tag)
}

type TagListAfter interface {
	TagListAfter([]entity.Tag)
}
