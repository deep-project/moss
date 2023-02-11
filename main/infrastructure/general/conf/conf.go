package conf

import (
	"github.com/spf13/viper"
	"moss/infrastructure/general/command"
	"strings"
)

type DbDriver string

const (
	fieldAddr = "addr"
	fieldDB   = "db"
	fieldDSN  = "dsn"

	DbDriverSqlite     DbDriver = "sqlite"
	DbDriverMysql      DbDriver = "mysql"
	DbDriverPostgresql DbDriver = "postgresql"
)

var (
	Addr string   // 监听地址
	DB   DbDriver // 数据库类型
	DSN  string   // 链接数据库DSN
)

func init() {
	p := viper.New()

	// get by file
	p.SetConfigFile(command.ConfFilePath)
	p.SetDefault(fieldAddr, command.Addr)
	p.SetDefault(fieldDB, "sqlite")
	p.SetDefault(fieldDSN, "./moss.db?_pragma=journal_mode(WAL)") // sqlite 默认使用 WAL模式

	_ = p.ReadInConfig()
	_ = p.WriteConfig()

	// get by env
	p.SetEnvPrefix("moss")
	_ = p.BindEnv(fieldAddr)
	_ = p.BindEnv(fieldDB)
	_ = p.BindEnv(fieldDSN)

	if command.Addr != "" {
		p.Set(fieldAddr, command.Addr)
	}

	Addr = p.GetString(fieldAddr)
	DB = FormatDbDriver(p.GetString(fieldDB))
	DSN = p.GetString(fieldDSN)
}

func FormatDbDriver(val string) DbDriver {
	switch strings.ToLower(val) {
	case "sqlite", "sqlite3":
		return DbDriverSqlite
	case "mysql", "mariadb", "maria":
		return DbDriverMysql
	case "pgsql", "postgres", "postgresql":
		return DbDriverPostgresql
	default:
		return DbDriver(val)
	}
}
