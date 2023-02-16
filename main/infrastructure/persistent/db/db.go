package db

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"os"
	"strings"
	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"moss/infrastructure/general/conf"
)

var DB *gorm.DB

func init() {
	if err := Init(); err != nil {
		panic(err)
	}
}

func Init() error {
	dial, err := newDial(conf.DB, conf.DSN)
	if err != nil {
		return err
	}
	db, err := gorm.Open(dial, &gorm.Config{
		PrepareStmt:                              true,  // 缓存预编译语句
		AllowGlobalUpdate:                        false, // 关闭无条件的全局更新（关闭可以防止全局更新删除）
		QueryFields:                              true,  // 查询 * 时，会自动填写所有字段名
		DisableForeignKeyConstraintWhenMigrating: true,  // 禁用外键约束
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func newDial(name conf.DbDriver, dsn string) (gorm.Dialector, error) {
	switch name {
	case conf.DbDriverSqlite:
		return sqlite.Open(dsn), nil
	case conf.DbDriverMysql:
		return mysql.Open(dsn), nil
	case conf.DbDriverPostgresql:
		return postgres.Open(dsn), nil
	default:
		return nil, fmt.Errorf("database type error: %s", name)
	}
}

func GetSize() int64 {
	switch conf.DB {
	case conf.DbDriverMysql:
		return getMysqlDbSize()
	case conf.DbDriverPostgresql:
		return getPostgresDbSize()
	case conf.DbDriverSqlite:
		return getSqliteDbSize()
	}
	return 0
}

func getSqliteDbSize() int64 {
	p := conf.DSN
	ps := strings.Split(conf.DSN, "?")
	if len(ps) > 0 {
		p = ps[0]
	}
	fi, err := os.Stat(p)
	if err != nil {
		return 0
	}
	return fi.Size()
}

func getMysqlDbSize() int64 {
	var r = struct {
		Size int64 `json:"size"`
	}{}
	DB.Raw("SELECT sum(DATA_LENGTH)+sum(INDEX_LENGTH) as size FROM information_schema.TABLES where TABLE_SCHEMA=?;", DB.Migrator().CurrentDatabase()).Scan(&r)
	return r.Size
}

func getPostgresDbSize() int64 {
	var r = struct {
		Size int64 `json:"size"`
	}{}
	DB.Raw(`select pg_database_size(?) as size`, DB.Migrator().CurrentDatabase()).Scan(&r)
	return r.Size
}
