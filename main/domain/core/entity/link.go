package entity

type Link struct {
	ID          int    `gorm:"type:int;size:32;primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"type:varchar(150);default:'';index"        json:"name"`
	URL         string `gorm:"type:varchar(250);default:''"              json:"url"`
	Logo        string `gorm:"type:varchar(250);default:''"              json:"logo"`
	CreateTime  int64  `gorm:"type:int;size:32"                          json:"create_time"`
	ExpireTime  int64  `gorm:"type:int;size:32;default:0;index"          json:"expire_time"`
	Note        string `gorm:"type:varchar(250);default:''"              json:"note"`
	Detect      bool   `gorm:"type:boolean; default:false; index"        json:"detect"`       // 自动检验开关
	DetectDelay int64  `gorm:"type:int; default:0;"                      json:"detect_delay"` // 自动检验延时 （分钟）
	Status      bool   `gorm:"type:boolean; default:false; index"        json:"status"`       // 状态 true:上链   false:下连
}
