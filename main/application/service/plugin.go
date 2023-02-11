package service

import (
	"moss/application/dto"
	"moss/application/mapper"
	"moss/domain/support/entity"
	"moss/domain/support/factory"
	"moss/domain/support/service"
	"moss/infrastructure/support/log"
)

func PluginInit(items ...entity.PluginEntry) {
	for _, item := range items {
		if err := service.Plugin.Init(item); err != nil {
			log.Error("plugin loaded failed", log.Any("info", item.Info()), log.Err(err))
		} else {
			log.Debug("plugin loaded successfully", log.Any("info", item.Info()))
		}
	}
}

func PluginList() []dto.PluginList {
	return mapper.PluginItemsToPluginInfoList(service.Plugin.Items)
}

func PluginLogList(id string, page, limit int) (any, error) {
	filePath := factory.GetPluginLogFilePath(id)
	return logRead(filePath, page, limit)
}
