package aggregate

import (
	"moss/domain/config/entity"
	"moss/domain/config/service"
	"moss/infrastructure/general/message"
)

type Config struct {
	Site     *entity.Site
	Admin    *entity.Admin
	Router   *entity.Router
	Upload   *entity.Upload
	Cache    *entity.Cache
	Theme    *entity.Theme
	Template *entity.Template
	Log      *entity.Log
	API      *entity.API
	Sitemap  *entity.Sitemap
	More     *entity.More
}

func NewConfig() *Config {
	return &Config{
		Site:     entity.NewSite(),
		Admin:    entity.NewAdmin(),
		Router:   entity.NewRouter(),
		Upload:   entity.NewUpload(),
		Cache:    entity.NewCache(),
		Theme:    entity.NewTheme(),
		Template: entity.NewTemplate(),
		Log:      entity.NewLog(),
		API:      entity.NewAPI(),
		Sitemap:  entity.NewSitemap(),
		More:     entity.NewMore(),
	}
}

func (c *Config) Items() []service.Config {
	return []service.Config{
		c.Site,
		c.Admin,
		c.Router,
		c.Upload,
		c.Cache,
		c.Theme,
		c.Template,
		c.Log,
		c.API,
		c.Sitemap,
		c.More,
	}
}

// Get 根据id获取内部配置项
func (c *Config) Get(id string) (service.Config, error) {
	if id == "" {
		return nil, message.ErrIdRequired
	}
	for _, v := range c.Items() {
		if v.ConfigID() == id {
			return v, nil
		}
	}
	return nil, message.ErrNotFound
}

// Save 根据ID保存配置
func (c *Config) Save(id string, data []byte) error {
	item, err := c.Get(id)
	if err != nil {
		return err
	}
	return service.Save(item, data)
}
