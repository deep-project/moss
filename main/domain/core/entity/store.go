package entity

import (
	"moss/domain/core/vo"
)

type Store struct {
	// store字段
	ID              int   `gorm:"type:int;size:32;primaryKey;autoIncrement" json:"id"`
	StoreCreateTime int64 `gorm:"type:int;size:32;index"                    json:"store_create_time"`

	// 可索引字段
	Slug       string `gorm:"type:varchar(150);default:'';index"  json:"slug"`
	Title      string `gorm:"type:varchar(250);default:'';index"  json:"title"`
	CategoryID int    `gorm:"type:int; size:32;default:0; index"  json:"category_id"`

	// 其他字段
	Thumbnail    string         `gorm:"type:varchar(250);default:''"  json:"thumbnail"`
	Views        int            `gorm:"type:int;size:32;default:0"    json:"views"`
	Description  string         `gorm:"type:varchar(250);default:''"  json:"description"`
	Keywords     string         `gorm:"type:varchar(250);default:''"  json:"keywords"`
	Content      string         `gorm:"type:string"                   json:"content"`
	Extends      vo.Extends     `gorm:"type:string"                   json:"extends"`
	Tags         vo.StringArray `gorm:"type:text;"                    json:"tags"`          // 直接创建tags
	CategoryName string         `gorm:"type:varchar(250);default:''"  json:"category_name"` // 直接通过分类名创建，优先级小于category_id
	CreateTime   int64          `gorm:"type:int;size:32"              json:"create_time"`
}
