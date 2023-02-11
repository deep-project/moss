package service

import (
	"moss/domain/core/service"
)

// DeleteTag 删除标签
func DeleteTag(id int) error {
	if err := service.Tag.Delete(id); err != nil {
		return err
	}
	return service.Mapping.DeleteTag(id)
}

// BatchDeleteTag 批量删除
func BatchDeleteTag(ids []int) (err error) {
	for _, id := range ids {
		if err = DeleteTag(id); err != nil {
			return
		}
	}
	return
}
