package repository

import (
	"moss/infrastructure/support/log"
)

func init() {
	MigrateTable()
}

func MigrateTable() {

	if err := Article.MigrateTable(); err != nil {
		log.Error("migrate article table error", log.Err(err))
	}

	if err := Category.MigrateTable(); err != nil {
		log.Error("migrate category table error", log.Err(err))
	}

	if err := Tag.MigrateTable(); err != nil {
		log.Error("migrate tag table error", log.Err(err))
	}

	if err := Mapping.MigrateTable(); err != nil {
		log.Error("migrate mapping table error", log.Err(err))
	}

	if err := Link.MigrateTable(); err != nil {
		log.Error("migrate link table error", log.Err(err))
	}

	if err := Store.MigrateTable(); err != nil {
		log.Error("migrate store table error", log.Err(err))
	}
}
