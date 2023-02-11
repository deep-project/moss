package entity

import (
	"errors"
	"fmt"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"runtime"
	"time"
)

type Plugin struct {
	Entry       PluginEntry
	Info        *PluginInfo
	CronID      cron.EntryID
	Running     bool
	RunTime     time.Time
	RunError    error
	RunCount    int
	RunDuration time.Duration
	Log         *zap.Logger
}

func (p *Plugin) Load() error {
	return p.Entry.Load(p)
}

// Run 用于定时任务接口
func (p *Plugin) Run() {
	_ = p.RunWithError()
}

func (p *Plugin) RunWithError() error {
	if p.Running {
		return errors.New("plugin is running")
	}
	p.Running = true
	defer func() { p.Running = false }()
	p.RunTime = time.Now()
	defer p.Panic()
	p.RunError = p.Entry.Run(p)
	p.RunCount++
	p.RunDuration = time.Since(p.RunTime)
	return p.RunError
}

// Panic 拦截panic错误，防止程序崩溃
func (p *Plugin) Panic() {
	if r := recover(); r != nil {
		var buf = make([]byte, 10240)
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}
		buf = buf[:runtime.Stack(buf, false)]
		p.Log.DPanic("panic error", zap.Error(err), zap.String("buf", string(buf)))
		p.RunError = err
	}
}

type PluginEntry interface {
	Info() *PluginInfo      // 插件信息
	Load(ctx *Plugin) error // 插件装载
	Run(ctx *Plugin) error  // 插件执行
}

type PluginInfo struct {
	ID         string `json:"id"`
	About      string `json:"about"`
	RunEnable  bool   `json:"run_enable"`  // 启用执行任务
	CronEnable bool   `json:"cron_enable"` // 启用定时任务功能
	NoOptions  bool   `json:"no_options"`  // 没有配置项
	HideLogs   bool   `json:"hide_logs"`   // 是否隐藏日志记录
	PluginInfoPersistent
}

func (p *PluginInfo) SetCronStart(val bool) {
	p.CronStart = val
}

func (p *PluginInfo) SetCronExp(val string) {
	p.CronExp = val
}

type PluginInfoPersistent struct {
	CronStart bool   `json:"cron_start"` // 启动定时任务
	CronExp   string `json:"cron_exp"`   // 定时任务表达式
}
