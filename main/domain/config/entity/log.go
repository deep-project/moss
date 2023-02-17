package entity

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"moss/infrastructure/general/constant"
	"path/filepath"
)

type Log struct {
	App              *LogItem `json:"app"`
	SQL              *LogItem `json:"sql"`
	SlowSQL          *LogItem `json:"slow_sql"`
	Visitor          *LogItem `json:"visitor"`
	Spider           *LogItem `json:"spider"`
	Plugin           *LogItem `json:"plugin"`
	SlowSQLThreshold int      `json:"slow_sql_threshold"`
	SpiderFeature    []string `json:"spider_feature"`
	HttpPoolSize     int      `json:"http_pool_size"`
	SqlPoolSize      int      `json:"sql_pool_size"`
}

func NewLog() *Log {
	l := &Log{
		SlowSQLThreshold: 200,
		SpiderFeature: []string{
			// 知名
			"Googlebot", "bingBot", "Baiduspider", "msnbot", "YandexBot", "Yahoo! Slurp",
			// 国内
			"360Spider", "SogouSpider", "Sogou web spider", "Sogou Pic Spider", "DNSPod", "Bytespider",
			"AspiegelBot", "YisouSpider", "DuckDuckGo", "JikeSpider", "YoudaoBot", "JianKongBao",
			// 国外
			"AhrefsBot", "MJ12bot", "SemrushBot", "lycos", "Twitterbot", "DotBot", "MauiBot", "Rogerbot",
			"exabot", "Teoma", "DOCOMO", "Speedy Spider", "LinkpadBot",
			// 通用特征
			"slurp", "robozilla", "bot", "spider", "crawl",
		},
		App:          NewLogItem("app", zapcore.EncoderConfig{MessageKey: "msg", LevelKey: "level", TimeKey: "time", NameKey: "logger", CallerKey: "file", StacktraceKey: "stacktrace"}),
		SQL:          NewLogItem("sql", zapcore.EncoderConfig{MessageKey: "msg", LevelKey: "level", TimeKey: "time"}),
		SlowSQL:      NewLogItem("slow_sql", zapcore.EncoderConfig{MessageKey: "msg", LevelKey: "level", TimeKey: "time"}),
		Visitor:      NewLogItem("visitor", zapcore.EncoderConfig{TimeKey: "time"}),
		Spider:       NewLogItem("spider", zapcore.EncoderConfig{TimeKey: "time"}),
		Plugin:       &LogItem{ID: "plugin", Enable: true, MaxSize: 20, MaxAge: 5, MaxBackups: 1, Compress: true},
		HttpPoolSize: 1000,
		SqlPoolSize:  1000,
	}
	return l
}

func (*Log) ConfigID() string {
	return "log"
}

func (l *Log) Items() []*LogItem {
	return []*LogItem{l.App, l.SQL, l.SlowSQL, l.Visitor, l.Spider}
}

func (l *Log) GetByID(id string) *LogItem {
	for _, item := range l.Items() {
		if item.ID == id {
			return item
		}
	}
	return nil
}

type LogItem struct {
	ID         string                `json:"id"`
	Level      zapcore.Level         `json:"level"`
	Enable     bool                  `json:"enable"`
	FilePath   string                `json:"-"` // 安全考虑,暂不允许自定义
	MaxSize    int                   `json:"max_size"`
	MaxAge     int                   `json:"max_age"`
	MaxBackups int                   `json:"max_backups"`
	Compress   bool                  `json:"compress"`
	ZapConfig  zapcore.EncoderConfig `json:"-"`
}

func NewLogItem(id string, zapEncoder zapcore.EncoderConfig) *LogItem {
	return &LogItem{
		ID:         id,
		Enable:     true,
		FilePath:   filepath.Join(constant.LogDir, fmt.Sprintf("%s.log", id)),
		MaxSize:    100,
		MaxAge:     5,
		MaxBackups: 1,
		Compress:   true,
		ZapConfig:  zapConfig(zapEncoder),
	}
}

func zapConfig(opt zapcore.EncoderConfig) zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     opt.MessageKey,
		LevelKey:       opt.LevelKey,
		TimeKey:        opt.TimeKey,
		NameKey:        opt.NameKey,
		CallerKey:      opt.CallerKey,
		StacktraceKey:  opt.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochMillisTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
}
