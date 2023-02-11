package entity

type MappingArticleTag struct {
	ID        int `gorm:"type:int;size:32;primaryKey;autoIncrement"                  json:"id"`
	ArticleID int `gorm:"type:int;size:32;uniqueIndex:idx_mapping_article_tag"       json:"article_id"`
	TagID     int `gorm:"type:int;size:32;uniqueIndex:idx_mapping_article_tag;index" json:"tag_id"`
}
