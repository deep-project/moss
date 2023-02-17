package entity

import (
	"moss/infrastructure/support/cache/core"
	"moss/infrastructure/utils/timex"
)

type Cache struct {
	Enable       bool          `json:"enable"`
	Options      *CacheOptions `json:"options"`
	ActiveDriver string        `json:"active_driver"`
	Driver       *core.Driver  `json:"driver"`
}

func NewCache() *Cache {
	return &Cache{
		ActiveDriver: core.BadgerDriverName,
		Driver:       core.NewDriver(),
		Options: &CacheOptions{
			Home:     &CacheOptionItem{Enable: true, TTL: timex.Duration{Number: 30, Unit: timex.DurationMinute}},
			Article:  &CacheOptionItem{Enable: true, TTL: timex.Duration{Number: 1, Unit: timex.DurationDay}},
			Category: &CacheOptionItem{Enable: true, TTL: timex.Duration{Number: 8, Unit: timex.DurationHour}},
			Tag:      &CacheOptionItem{Enable: true, TTL: timex.Duration{Number: 16, Unit: timex.DurationHour}},
			Sitemap:  &CacheOptionItem{Enable: false, TTL: timex.Duration{Number: 10, Unit: timex.DurationMinute}},
			Page:     &CacheOptionItem{Enable: false, TTL: timex.Duration{Number: 24, Unit: timex.DurationHour}},
		},
	}
}

func (*Cache) ConfigID() string {
	return "cache"
}

// CurrentDriver 当前的驱动
func (c *Cache) CurrentDriver() (core.Cache, error) {
	return c.Driver.Get(c.ActiveDriver)
}

func (c *Cache) GetOption(name string) *CacheOptionItem {
	switch name {
	case "home":
		return c.Options.Home
	case "article":
		return c.Options.Article
	case "category":
		return c.Options.Category
	case "tag":
		return c.Options.Tag
	case "sitemap":
		return c.Options.Sitemap
	case "page":
		return c.Options.Page
	}
	return nil
}

type CacheOptions struct {
	Home     *CacheOptionItem `json:"home"`
	Article  *CacheOptionItem `json:"article"`
	Category *CacheOptionItem `json:"category"`
	Tag      *CacheOptionItem `json:"tag"`
	Sitemap  *CacheOptionItem `json:"sitemap"`
	Page     *CacheOptionItem `json:"page"`
}

type CacheOptionItem struct {
	Enable bool           `json:"enable"`
	TTL    timex.Duration `json:"ttl"`
}
