package config

import (
	"moss/domain/config/aggregate"
	"moss/domain/config/service"
)

var Config = aggregate.NewConfig()

func init() {
	for _, item := range Config.Items() {
		if err := service.Sync(item); err != nil {
			panic(err)
		}
	}
}
