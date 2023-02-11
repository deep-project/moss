package service

import (
	"encoding/json"
	"errors"
	"github.com/robfig/cron/v3"
	"moss/domain/support/entity"
	"moss/domain/support/factory"
	"moss/domain/support/repository"
	"moss/domain/support/utils"
	"moss/infrastructure/general/message"
	"time"
)

var Plugin = NewPluginService()

type PluginService struct {
	Items []*entity.Plugin
	Cron  *cron.Cron
}

func NewPluginService() *PluginService {
	p := PluginService{Cron: cron.New(cron.WithSeconds())}
	p.Cron.Start()
	return &p
}

// Init 初始化插件
func (p *PluginService) Init(entry entity.PluginEntry) (err error) {
	item, err := factory.NewPlugin(entry)
	if err != nil {
		return
	}
	if p.Exists(item.Info.ID) {
		return message.ErrIdAlreadyExists
	}
	if err = p.Sync(item); err != nil {
		return
	}
	if err = item.Load(); err != nil {
		return
	}
	// 开启定时执行
	if item.Info.CronEnable {
		if item.Info.CronExp == "" {
			return errors.New("cron exp required")
		}
		if err = utils.CheckCronExp(item.Info.CronExp); err != nil {
			return err
		}
		if item.Info.CronStart {
			if err = p.StartCron(item); err != nil {
				return err
			}
		}
	} else {
		item.Info.CronStart = false // 如果未开启定时器，直接关闭定时器启动状态
	}
	p.Items = append(p.Items, item)
	return
}

// Exists 判断是否已存在在entries中
func (p *PluginService) Exists(id string) bool {
	for _, item := range p.Items {
		if item.Info.ID == id {
			return true
		}
	}
	return false
}

func (p *PluginService) Get(id string) (*entity.Plugin, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	for _, item := range p.Items {
		if item.Info.ID == id {
			return item, nil
		}
	}
	return nil, errors.New("plugin not found")
}

// Sync 同步插件, 先从数据库拉取,然后推送到数据库
func (p *PluginService) Sync(item *entity.Plugin) error {
	if err := p.Pull(item); err != nil {
		return err
	}
	return p.Push(item)
}

// Pull 从存储拉取插件数据
func (p *PluginService) Pull(item *entity.Plugin) error {
	if err := p.PullInfo(item); err != nil {
		return err
	}
	return p.PullOptions(item)
}
func (p *PluginService) PullInfo(item *entity.Plugin) error {
	info, err := repository.Plugin.GetInfo(item.Info.ID)
	if err != nil {
		return err
	}
	return p.mergeInfo(info, item)
}
func (p *PluginService) PullOptions(item *entity.Plugin) error {
	options, err := repository.Plugin.GetOptions(item.Info.ID)
	if err != nil {
		return err
	}
	return p.mergeOptions(options, item)
}

// Push 推送数据到数据库
func (p *PluginService) Push(item *entity.Plugin) error {
	if err := p.PushInfo(item); err != nil {
		return err
	}
	return p.PushOptions(item)
}
func (p *PluginService) PushInfo(item *entity.Plugin) error {
	info, err := json.Marshal(item.Info.PluginInfoPersistent)
	if err != nil {
		return err
	}
	return p.SaveInfo(item.Info.ID, info)
}
func (p *PluginService) PushOptions(item *entity.Plugin) error {
	options, err := json.Marshal(item.Entry)
	if err != nil {
		return err
	}
	return p.SaveOptions(item.Info.ID, options)
}

func (p *PluginService) mergeInfo(info []byte, item *entity.Plugin) error {
	if len(info) == 0 {
		return nil
	}
	return json.Unmarshal(info, &item.Info.PluginInfoPersistent)
}
func (p *PluginService) mergeOptions(options []byte, item *entity.Plugin) error {
	if len(options) == 0 {
		return nil
	}
	return json.Unmarshal(options, item.Entry)
}

func (p *PluginService) SaveInfo(id string, info []byte) error {
	return repository.Plugin.SaveInfo(id, info)
}
func (p *PluginService) SaveOptions(id string, options []byte) error {
	return repository.Plugin.SaveOptions(id, options)
}

func (p *PluginService) Run(id string) error {
	item, err := p.Get(id)
	if err != nil {
		return err
	}
	if !item.Info.RunEnable {
		return errors.New("plugin run is disabled")
	}
	return item.RunWithError()
}

func (p *PluginService) GetOptions(id string) (entity.PluginEntry, error) {
	item, err := p.Get(id)
	if err != nil {
		return nil, err
	}
	return item.Entry, nil
}

func (p *PluginService) UpdateOptions(id string, options []byte) (err error) {
	item, err := p.Get(id)
	if err != nil {
		return err
	}
	if err = p.SaveOptions(id, options); err != nil {
		return
	}
	return p.mergeOptions(options, item)
}

func (p *PluginService) UpdateCronStart(id string, val bool) error {
	item, err := p.Get(id)
	if err != nil {
		return err
	}
	if !item.Info.CronEnable {
		return errors.New("cron is disabled")
	}
	item.Info.SetCronStart(val)
	if err = p.PushInfo(item); err != nil {
		return err
	}
	if val {
		return p.RestartCron(item)
	}
	p.StopCron(item)
	return nil
}

func (p *PluginService) UpdateCronExp(id string, val string) error {
	item, err := p.Get(id)
	if err != nil {
		return err
	}
	if err = utils.CheckCronExp(val); err != nil {
		return err
	}
	item.Info.SetCronExp(val)
	if err = p.PushInfo(item); err != nil {
		return err
	}
	if item.Info.CronStart {
		return p.RestartCron(item) // 重启
	}
	return nil
}

// RestartCron 重启cron
func (p *PluginService) RestartCron(item *entity.Plugin) error {
	p.StopCron(item)
	return p.StartCron(item)
}

func (p *PluginService) StartCron(item *entity.Plugin) (err error) {
	item.CronID, err = p.Cron.AddJob(item.Info.CronExp, item)
	return
}

func (p *PluginService) StopCron(item *entity.Plugin) {
	p.Cron.Remove(item.CronID)
}

// NextRunTime 下次运行时间
func (p *PluginService) NextRunTime(cronID cron.EntryID) (res time.Time) {
	entry := p.Cron.Entry(cronID)
	if entry.Valid() {
		return entry.Next
	}
	return
}
