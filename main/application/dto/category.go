package dto

import (
	"moss/domain/core/entity"
)

type CategoryTree struct {
	entity.Category
	Children []CategoryTree `json:"children"`
}
