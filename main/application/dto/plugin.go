package dto

import (
	"moss/domain/support/entity"
)

type PluginList struct {
	entity.PluginInfo
	RunTime     int64  `json:"run_time"` // 时间戳
	RunError    string `json:"run_error"`
	RunCount    int    `json:"run_count"`
	RunDuration int64  `json:"run_duration"` // 毫秒
	NextRunTime int64  `json:"next_run_time"`
}
