package entity

import (
	"moss/domain/config"
	"strconv"
	"strings"
	"time"
)

type Category struct {
	ID          int    `gorm:"type:int;size:32;primaryKey;autoIncrement" json:"id"`
	ParentID    int    `gorm:"type:int;size:32;default:0;index"          json:"parent_id"`
	Slug        string `gorm:"type:varchar(150);uniqueIndex;not null"    json:"slug"`
	Name        string `gorm:"type:varchar(150);default:'';index"        json:"name"`
	CreateTime  int64  `gorm:"type:int;size:32"                          json:"create_time"`
	Title       string `gorm:"type:varchar(250);default:''"              json:"title"`
	Keywords    string `gorm:"type:varchar(250);default:''"              json:"keywords"`
	Description string `gorm:"type:varchar(250);default:''"              json:"description"`
}

func (c *Category) FullURL() string {
	return config.Config.Site.GetURL() + c.URL()
}

func (c *Category) URL() string {
	return strings.Replace(config.Config.Router.GetCategoryRule(), ":slug", c.Slug, 1)
}

func (c *Category) PageURL(page int) string {
	if page < 2 {
		return c.URL()
	}
	res := strings.Replace(config.Config.Router.GetCategoryPageRule(), ":slug", c.Slug, 1)
	return strings.Replace(res, ":page<int>", strconv.Itoa(page), 1)
}

func (c *Category) FullPageURL(page int) string {
	return config.Config.Site.GetURL() + c.PageURL(page)
}

func (c *Category) CreateTimeFormat(layouts ...string) string {
	if c.CreateTime == 0 {
		return ""
	}
	var layout = "2006-01-02 15:04:05"
	if len(layouts) > 0 {
		layout = layouts[0]
	}
	return time.Unix(c.CreateTime, 0).Format(layout)
}
