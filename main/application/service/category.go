package service

import (
	"moss/domain/core/aggregate"
	"moss/domain/core/entity"
	"moss/domain/core/repository/context"
	"moss/domain/core/service"
	"moss/domain/core/utils"
)

// CategoryTree 获取所有的分类树
func CategoryTree() ([]aggregate.CategoryTree, error) {
	items, err := service.Category.List(context.NewContext(1000, "")) // 限制最大值
	if err != nil {
		return nil, err
	}
	return utils.MakeCategoryTree(utils.CategoryEntityListToCategoryTreeList(items), 0), nil
}

// CategoryGetOrCreate 获取或创建类目
func CategoryGetOrCreate(name string) (*entity.Category, error) {
	res, err := service.Category.GetOrCreate(name)
	if err != nil { // 如果出现错误，有可能是并发请求造成的，再次尝试获取就可以得到正确的结果
		res, err = service.Category.GetOrCreate(name)
	}
	return res, err
}
