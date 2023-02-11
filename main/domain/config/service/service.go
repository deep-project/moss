package service

import (
	"encoding/json"
	"moss/domain/config/repository"
)

type Config interface {
	ConfigID() string
}

// Sync 同步配置, 进行一次拉取并推送操作
func Sync(item Config) error {
	if err := Pull(item); err != nil {
		return err
	}
	return Push(item)
}

// Pull 拉取配置，把数据库中的配置拉取到内存
func Pull(item Config) error {
	data, err := repository.Get(item.ConfigID())
	if err != nil {
		return err
	}
	return Merge(data, item)
}

func Push(item Config) error {
	data, err := json.Marshal(item)
	if err != nil {
		return err
	}
	return repository.Save(item.ConfigID(), data)
}

// Save 保存配置
func Save(item Config, data []byte) error {
	if err := repository.Save(item.ConfigID(), data); err != nil {
		return err
	}
	return Merge(data, item)
}

// Merge 把数据合并到配置地址
func Merge(data []byte, item Config) error {
	if len(data) == 0 {
		return nil
	}
	return json.Unmarshal(data, item)
}
