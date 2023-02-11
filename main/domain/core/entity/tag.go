package entity

import (
	"moss/domain/config"
	"strconv"
	"strings"
	"time"
)

type Tag struct {
	ID          int    `gorm:"type:int;size:32;primaryKey;autoIncrement" json:"id"`
	Slug        string `gorm:"type:varchar(150);uniqueIndex;not null"    json:"slug"`
	Name        string `gorm:"type:varchar(150);uniqueIndex;not null"    json:"name"`
	CreateTime  int64  `gorm:"type:int;size:32"                          json:"create_time"`
	Title       string `gorm:"type:varchar(250);default:''"              json:"title"`
	Keywords    string `gorm:"type:varchar(250);default:''"              json:"keywords"`
	Description string `gorm:"type:varchar(250);default:''"              json:"description"`
}

func (t *Tag) FullURL() string {
	return config.Config.Site.GetURL() + t.URL()
}

func (t *Tag) URL() string {
	return strings.Replace(config.Config.Router.GetTagRule(), ":slug", t.Slug, 1)
}

func (t *Tag) PageURL(page int) string {
	if page < 2 {
		return t.URL()
	}
	res := strings.Replace(config.Config.Router.GetTagPageRule(), ":slug", t.Slug, 1)
	return strings.Replace(res, ":page<int>", strconv.Itoa(page), 1)
}

func (t *Tag) FullPageURL(page int) string {
	return config.Config.Site.GetURL() + t.PageURL(page)
}

func (t *Tag) CreateTimeFormat(layouts ...string) string {
	if t.CreateTime == 0 {
		return ""
	}
	var layout = "2006-01-02 15:04:05"
	if len(layouts) > 0 {
		layout = layouts[0]
	}
	return time.Unix(t.CreateTime, 0).Format(layout)
}
