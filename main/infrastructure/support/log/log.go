package log

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
	"moss/domain/config"
	"moss/infrastructure/persistent/db"
	"moss/infrastructure/support/log/core"
	"runtime"
	"time"
)

var (
	Err    = zap.Error
	Any    = zap.Any
	String = zap.String
	Binary = zap.Binary
)

var (
	App     = core.New(config.Config.Log.App)
	SQL     = core.New(config.Config.Log.SQL)
	SlowSQL = core.New(config.Config.Log.SlowSQL)
	Visitor = core.New(config.Config.Log.Visitor)
	Spider  = core.New(config.Config.Log.Spider)
)

func init() {
	Init()
	db.DB.Logger = &GormLogger{}
}

func Init() {
	App.Init()
	SQL.Init()
	SlowSQL.Init()
	Visitor.Init()
	Spider.Init()
}

func Debug(msg string, fields ...zapcore.Field)  { App.Debug(msg, fields...) }
func Info(msg string, fields ...zapcore.Field)   { App.Info(msg, fields...) }
func Warn(msg string, fields ...zapcore.Field)   { App.Warn(msg, fields...) }
func Error(msg string, fields ...zapcore.Field)  { App.Error(msg, fields...) }
func Panic(msg string, fields ...zapcore.Field)  { App.Panic(msg, fields...) }
func DPanic(msg string, fields ...zapcore.Field) { App.DPanic(msg, fields...) }
func Fatal(msg string, fields ...zapcore.Field)  { App.Fatal(msg, fields...) }

func ErrorShortcut(msg string, err error) { App.ErrorShortcut(msg, err) }
func WarnShortcut(msg string, err error)  { App.WarnShortcut(msg, err) }

// GormLogger gorm日志驱动
type GormLogger struct{}

func (g *GormLogger) LogMode(level logger.LogLevel) logger.Interface             { return g }
func (g *GormLogger) Info(ctx context.Context, msg string, data ...interface{})  { SQL.Info(msg) }
func (g *GormLogger) Warn(ctx context.Context, msg string, data ...interface{})  { SQL.Warn(msg) }
func (g *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) { SQL.Error(msg) }
func (g *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if SQL.IsClosed() && SlowSQL.IsClosed() {
		return
	}
	_, file, line, _ := runtime.Caller(3)
	sql, rows := fc()
	Client.InvokePoolSQL(SqlData{
		File:      file,
		Line:      line,
		SQL:       sql,
		Rows:      rows,
		BeginTime: begin,
	})
}
