package service

import (
	"moss/domain/core/aggregate"
	"moss/domain/core/repository/context"
	"moss/domain/core/service"
	"moss/domain/core/utils"
)

// CategoryTree 获取所有的分类树
func CategoryTree() ([]aggregate.CategoryTree, error) {
	items, err := service.Category.List(context.NewContext(50000, "")) // 限制最大值
	if err != nil {
		return nil, err
	}
	return utils.MakeCategoryTree(utils.CategoryEntityListToCategoryTreeList(items), 0), nil
}
