package service

import (
	"moss/domain/config"
)

func CacheSize() (_ int64, err error) {
	d, err := config.Config.Cache.CurrentDriver()
	if err != nil {
		return
	}
	return d.Size()
}
