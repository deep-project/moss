package entity

import (
	"moss/domain/config"
	"moss/domain/core/vo"
	"strings"
	"time"
)

type Article struct {
	ArticleBase
	ArticleDetail
}

type ArticleBase struct {
	ID          int    `gorm:"type:int;size:32;primaryKey;autoIncrement" json:"id"`
	Slug        string `gorm:"type:varchar(150);uniqueIndex;not null"    json:"slug"`
	Title       string `gorm:"type:varchar(250);default:'';index"        json:"title"`
	CreateTime  int64  `gorm:"type:int;size:32;index"                    json:"create_time"`
	CategoryID  int    `gorm:"type:int;size:32;default:0;index"          json:"category_id"`
	Views       int    `gorm:"type:int;size:32;default:0;index"          json:"views"`
	Thumbnail   string `gorm:"type:varchar(250);default:''"              json:"thumbnail"`
	Description string `gorm:"type:varchar(250);default:''"              json:"description"`
}

func (ArticleBase) TableName() string {
	return "article"
}

func (a *ArticleBase) FullURL() string {
	return config.Config.Site.GetURL() + a.URL()
}

func (a *ArticleBase) URL() string {
	return strings.Replace(config.Config.Router.GetArticleRule(), ":slug", a.Slug, 1)
}

func (a *ArticleBase) CreateTimeFormat(layouts ...string) string {
	if a.CreateTime == 0 {
		return ""
	}
	var layout = "2006-01-02 15:04:05"
	if len(layouts) > 0 {
		layout = layouts[0]
	}
	return time.Unix(a.CreateTime, 0).Format(layout)
}

type ArticleDetail struct {
	ArticleID int        `gorm:"type:int;size:32;primaryKey"   json:"article_id"`
	Keywords  string     `gorm:"type:varchar(250);default:''"  json:"keywords"`
	Content   string     `gorm:"type:string"                   json:"content"`
	Source    string     `gorm:"type:varchar(250)"             json:"source"`
	Extends   vo.Extends `gorm:"type:string"                   json:"extends"`
}
