package mapper

import (
	"moss/application/dto"
	"moss/domain/support/entity"
	"moss/domain/support/service"
)

func PluginItemsToPluginInfoList(items []*entity.Plugin) (res []dto.PluginList) {
	for _, item := range items {
		res = append(res, dto.PluginList{
			PluginInfo:  *item.Info,
			RunTime:     item.RunTime.Unix(),
			RunError:    item.RunError,
			RunCount:    item.RunCount,
			RunDuration: item.RunDuration.Milliseconds(),
			NextRunTime: service.Plugin.NextRunTime(item.CronID).Unix(),
		})
	}
	return
}
