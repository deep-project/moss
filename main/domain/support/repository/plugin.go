package repository

import (
	"fmt"
	"gorm.io/gorm/clause"
	"moss/infrastructure/persistent/db"
)

var Plugin = new(PluginRepository)

func init() {
	if err := db.DB.AutoMigrate(&PluginTable{}); err != nil {
		fmt.Println("migrate plugin table error: ", err.Error())
	}
}

type PluginTable struct {
	ID      string `gorm:"primaryKey;type:varchar(100);" json:"id"`
	Info    string `gorm:"type:string;"                  json:"info"`
	Options string `gorm:"type:string;"                  json:"options"`
}

func (PluginTable) TableName() string {
	return "plugin"
}

type PluginRepository struct{}

func (*PluginRepository) SaveOptions(id string, options []byte) error {
	return db.DB.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}, DoUpdates: clause.AssignmentColumns([]string{"options"})}).
		Create(&PluginTable{ID: id, Options: string(options)}).Error
}

func (*PluginRepository) SaveInfo(id string, info []byte) error {
	return db.DB.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}, DoUpdates: clause.AssignmentColumns([]string{"info"})}).
		Create(&PluginTable{ID: id, Info: string(info)}).Error
}

func (*PluginRepository) GetInfo(id string) ([]byte, error) {
	var info string
	err := db.DB.Model(&PluginTable{}).Where("id = ?", id).Pluck("info", &info).Error
	return []byte(info), err
}

func (*PluginRepository) GetOptions(id string) ([]byte, error) {
	var options string
	err := db.DB.Model(&PluginTable{}).Where("id = ?", id).Pluck("options", &options).Error
	return []byte(options), err
}
