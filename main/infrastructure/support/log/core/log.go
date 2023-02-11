package core

import (
	"github.com/gookit/color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"moss/domain/config/entity"
)

type Log struct {
	config *entity.LogItem
	client *zap.Logger
	writer *lumberjack.Logger
}

func New(conf *entity.LogItem) *Log {
	return &Log{config: conf}
}

func (l *Log) Init() {
	// 初始化之前
	// 先关闭 writer 和 client
	if l.writer != nil {
		_ = l.writer.Close()
	}
	if l.client != nil {
		l.client = nil
	}
	if l.config.Enable {
		l.initWriter()
		l.initClient()
	}
}

func (l *Log) initWriter() {
	l.writer = &lumberjack.Logger{
		Filename:   l.config.FilePath,
		MaxSize:    l.config.MaxSize,
		MaxAge:     l.config.MaxAge,
		MaxBackups: l.config.MaxBackups,
		Compress:   l.config.Compress,
	}
}

func (l *Log) initClient() {
	level := zapcore.InfoLevel
	skip := 1
	if l.config.ID == "app" {
		skip = 2
		level = l.config.Level
	}
	l.client = zap.New(
		zapcore.NewCore(zapcore.NewJSONEncoder(l.config.ZapConfig),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(l.writer)),
			zap.NewAtomicLevelAt(level),
		),
		zap.AddCallerSkip(skip), zap.AddCaller(),
	)
}

func (l *Log) IsClosed() bool {
	return !l.config.Enable || l.client == nil
}

// Debug 级别最低，多用于调试信息
func (l *Log) Debug(msg string, fields ...zapcore.Field) {
	if !l.IsClosed() {
		l.client.Debug(msg, fields...)
	}
}

// Info 一般等级，用来反馈系统的当前状态给最终用户的
func (l *Log) Info(msg string, fields ...zapcore.Field) {
	if !l.IsClosed() {
		l.client.Info(msg, fields...)
	}
}

// Warn 可修复的问题，系统可继续运行下去
func (l *Log) Warn(msg string, fields ...zapcore.Field) {
	if !l.IsClosed() {
		l.client.Warn(msg, fields...)
	}
}

// WarnShortcut Warn的快捷方法
func (l *Log) WarnShortcut(msg string, err error) {
	if !l.IsClosed() && err != nil {
		l.client.Warn(msg, zap.Error(err))
	}
}

// Error 可修复性，但无法确定系统会正常的工作下去
func (l *Log) Error(msg string, fields ...zapcore.Field) {
	if !l.IsClosed() {
		l.client.Error(msg, fields...)
	}
}

// ErrorShortcut Error的快捷方法
func (l *Log) ErrorShortcut(msg string, err error) {
	if !l.IsClosed() && err != nil {
		l.client.Error(msg, zap.Error(err))
	}
}

// Panic 恐慌级别，很严重，系统无法正常工作
// 会立即终止程序！
func (l *Log) Panic(msg string, fields ...zapcore.Field) {
	if !l.IsClosed() {
		l.client.Panic(msg, fields...)
	}
}

// DPanic 级别同 Panic，
// 区别是只有开发模式下才会终止程序
func (l *Log) DPanic(msg string, fields ...zapcore.Field) {
	if !l.IsClosed() {
		l.client.DPanic(msg, fields...)
	}
}

// Fatal 相当严重，会立即终止程序！
func (l *Log) Fatal(msg string, fields ...zapcore.Field) {
	if !l.IsClosed() {
		color.Red.Println(msg)
		l.client.Fatal(msg, fields...)
	}
}
