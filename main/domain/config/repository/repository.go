package repository

import (
	"fmt"
	"moss/infrastructure/persistent/db"
)

type Config struct {
	ID   string `gorm:"primaryKey;type:varchar(100);"`
	Data string `gorm:"type:string;"`
}

func init() {
	if err := MigrateTable(); err != nil {
		fmt.Println("migrate config table error: ", err.Error())
	}
}

func MigrateTable() error {
	return db.DB.AutoMigrate(&Config{})
}

func Save(id string, data []byte) error {
	return db.DB.Save(&Config{ID: id, Data: string(data)}).Error
}

func Get(id string) (data []byte, err error) {
	var str string
	err = db.DB.Model(&Config{}).Where("id = ?", id).Pluck("data", &str).Error
	return []byte(str), err
}
