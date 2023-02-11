package dto

type LogAPP struct {
	Level string         `json:"level"` // 级别
	File  string         `json:"file"`  // 文件
	Msg   string         `json:"msg"`   // 消息
	Data  map[string]any `json:"data"`  // 其他数据
}

type LogHTTP struct {
	Take      float64 `json:"take"`      // 耗时
	Status    int     `json:"status"`    // 状态码
	Depth     uint64  `json:"depth"`     // 访问深度
	IP        string  `json:"ip"`        // 访客ip
	Method    string  `json:"method"`    // 请求方法
	URL       string  `json:"url"`       // 访问URL
	Referer   string  `json:"referer"`   // 来路URL
	UserAgent string  `json:"userAgent"` // userAgent
	Headers   string  `json:"headers"`   // 全部请求头信息
}

type LogSQL struct {
	Path string  `json:"path"` // 文件路径
	SQL  string  `json:"sql"`  // sql语句
	Rows int64   `json:"rows"` // 查询条数
	Take float64 `json:"take"` // 耗时
}
