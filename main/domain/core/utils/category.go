package utils

import (
	"go.uber.org/zap"
	"moss/domain/core/aggregate"
	"moss/domain/core/entity"
	"moss/infrastructure/support/log"
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
	var count = 0
	fn = func(cid int) {
		count++
		for _, item := range ranges {
			if item.ID == cid {
				// 如果查询了100次还在继续，防止程序进入死循环，直接break
				if count > 100 {
					log.Warn("循环超过100次", zap.Any("结果集", res))
					break
				}
				// 如果条目已经存在结果集中，说明类目即将陷入死循环，防止崩溃，直接break
				if idInCategories(item.ID, res) {
					log.Warn("已存在结果在结果集中，检查是否有循环依赖的分类", zap.Int("id", item.ID), zap.Any("结果集", res))
					break
				}
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

func idInCategories(id int, items []entity.Category) bool {
	for _, item := range items {
		if item.ID == id {
			return true
		}
	}
	return false
}
