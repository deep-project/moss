package utils

import (
	"moss/domain/core/aggregate"
	"moss/domain/core/entity"
)

func CategoryEntityListToCategoryTreeList(items []entity.Category) (res []aggregate.CategoryTree) {
	for _, item := range items {
		res = append(res, CategoryEntityToCategoryTree(item))
	}
	return
}

func CategoryEntityToCategoryTree(item entity.Category) (res aggregate.CategoryTree) {
	res.Category = item
	return
}

func MakeCategoryTree(list []aggregate.CategoryTree, parentID int) (res []aggregate.CategoryTree) {
	for _, v := range list {
		if v.ParentID == parentID {
			var children = MakeCategoryTree(list, v.ID)
			if len(children) > 0 {
				v.Children = children
			}
			res = append(res, v)
		}
	}
	return res
}

// FindCategoryWithAncestors 查找类目和其祖先
func FindCategoryWithAncestors(id int, ranges []entity.Category) (res []entity.Category) {
	var fn func(int)
	fn = func(cid int) {
		for _, item := range ranges {
			if item.ID == cid {
				res = append(res, item)
				if item.ParentID != 0 {
					fn(item.ParentID)
				}
			}
		}
	}
	fn(id)
	return
}
