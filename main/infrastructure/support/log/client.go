package log

import (
	"github.com/panjf2000/ants/v2"
	"go.uber.org/zap"
	"moss/domain/config"
	"strconv"
	"strings"
	"time"
)

var Client = newLog()

type log struct {
	httpPool *ants.PoolWithFunc
	sqlPool  *ants.PoolWithFunc
}

func newLog() *log {
	l := &log{}
	l.initPool()
	return l
}

func (l *log) initPool() {
	var err error
	if l.httpPool, err = ants.NewPoolWithFunc(1000, l.poolWriteHTTP, ants.WithNonblocking(true)); err != nil {
		Warn("http log pool initialization error", Err(err))
	}
	if l.sqlPool, err = ants.NewPoolWithFunc(1000, l.poolWriteSQL, ants.WithNonblocking(true)); err != nil {
		Warn("sql log pool initialization error", Err(err))
	}
}

func (l *log) poolWriteHTTP(val any) {
	if entry, ok := val.(HttpData); ok {
		l.HTTP(entry)
	}
}

func (l *log) poolWriteSQL(val any) {
	if entry, ok := val.(SqlData); ok {
		l.SQL(entry)
	}
}

func (l *log) InvokePoolHTTP(entry HttpData) {
	if Visitor.IsClosed() && Spider.IsClosed() {
		return
	}
	if l.httpPool != nil {
		if err := l.httpPool.Invoke(entry); err != nil {
			Warn("invoke http pool error", Err(err))
		}
	}
}

func (l *log) InvokePoolSQL(entry SqlData) {
	if SQL.IsClosed() && SlowSQL.IsClosed() {
		return
	}
	if l.sqlPool != nil {
		if err := l.sqlPool.Invoke(entry); err != nil {
			Warn("invoke sql pool error", Err(err))
		}
	}
}

type HttpData struct {
	RequestTime time.Time // 请求时间
	Status      int       // 状态码
	Depth       uint64    // 访问深度
	IP          string    // 访客ip
	Method      string    // 请求方法
	URL         string    // 访问URL
	Referer     string    // 来路URL
	UserAgent   string    // userAgent
	Headers     string    // 全部请求头信息
	Path        string    // 访问路径
}

func (l *log) HTTP(entry HttpData) {
	client := Visitor
	if l.isSpider(entry.UserAgent) {
		client = Spider
	}
	if client.IsClosed() {
		return
	}
	// 如果是后台路径,则不打印 headers
	if strings.HasPrefix(entry.Path, config.Config.Router.GetAdminPath()) {
		entry.Headers = ""
	}
	client.Info("",
		zap.Float64("take", float64(time.Since(entry.RequestTime).Nanoseconds())/1e6), // 毫秒
		zap.Int("status", entry.Status),
		zap.Uint64("depth", entry.Depth),
		zap.String("ip", entry.IP),
		zap.String("method", entry.Method),
		zap.String("url", entry.URL),
		zap.String("referer", entry.Referer),
		zap.String("userAgent", entry.UserAgent),
		zap.String("headers", entry.Headers),
		//zap.String("region", ip2region.Region(ip)), // 访客位置
	)
}

func (l *log) isSpider(ua string) bool {
	if len(ua) == 0 {
		return true
	}
	for _, v := range config.Config.Log.SpiderFeature {
		if strings.Contains(ua, v) {
			return true
		}
	}
	return false
}

type SqlData struct {
	File      string
	Line      int
	SQL       string
	Rows      int64     // 查询条数
	BeginTime time.Time // 开始查询时间
}

func (l *log) SQL(entry SqlData) {
	take := float64(time.Since(entry.BeginTime).Nanoseconds()) / 1e6 // 毫秒
	path := entry.File + ":" + strconv.Itoa(entry.Line)
	client := SQL
	if take >= float64(config.Config.Log.SlowSQLThreshold) {
		client = SlowSQL
	}
	if client.IsClosed() {
		return
	}
	client.Info("",
		zap.String("file", path),
		zap.String("sql", entry.SQL),
		zap.Int64("rows", entry.Rows),
		zap.Float64("take", take),
	)
}
