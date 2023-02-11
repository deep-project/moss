package event

import (
	"moss/domain/core/entity"
)

type LinkCreateBefore interface {
	LinkCreateBefore(*entity.Link) error
}

type LinkCreateAfter interface {
	LinkCreateAfter(*entity.Link)
}

type LinkUpdateBefore interface {
	LinkUpdateBefore(item *entity.Link) error
}

type LinkUpdateAfter interface {
	LinkUpdateAfter(item *entity.Link)
}

type LinkDeleteBefore interface {
	LinkDeleteBefore(id int) error
}

type LinkDeleteAfter interface {
	LinkDeleteAfter(id int)
}

type LinkGetAfter interface {
	LinkGetAfter(*entity.Link)
}

type LinkListAfter interface {
	LinkListAfter([]entity.Link)
}
