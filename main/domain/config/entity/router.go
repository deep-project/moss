package entity

import (
	"errors"
	"github.com/duke-git/lancet/v2/random"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/gofiber/fiber/v2"
	"moss/infrastructure/general/constant"
	"strings"
)

type Router struct {
	AdminPath        string       `json:"admin_path"`
	SitemapPath      string       `json:"sitemap_path"`
	ArticleRule      string       `json:"article_rule"`
	CategoryRule     string       `json:"category_rule"`
	CategoryPageRule string       `json:"category_page_rule"`
	TagRule          string       `json:"tag_rule"`
	TagPageRule      string       `json:"tag_page_rule"`
	Options          fiber.Config `json:"options"`
	ProxyHeader      []string     `json:"proxy_header"`   // 代理标头 （当请求来自CDN或代理时，为保证正确获取客户端ip，可以指定代理标头）
	CompressLevel    int          `json:"compress_level"` // gzip压缩等级 （-1不压缩、0平衡、1速度最佳、2压缩最好（慢））
	MinifyCode       bool         `json:"minify_code"`
	ETag             bool         `json:"etag"`
	PprofSecret      string       `json:"prof_secret"`
}

const (
	RouterDefaultSitemapPath      = "/sitemap"
	RouterDefaultArticleRule      = "/article/{slug}"
	RouterDefaultCategoryRule     = "/category/{slug}"
	RouterDefaultCategoryPageRule = "/category/{slug}/{page}"
	RouterDefaultTagRule          = "/tag/{slug}"
	RouterDefaultTagPageRule      = "/tag/{slug}/{page}"
)

func NewRouter() *Router {
	return &Router{
		AdminPath:        constant.DefaultAdminPath,
		SitemapPath:      RouterDefaultSitemapPath,
		ArticleRule:      RouterDefaultArticleRule,
		CategoryRule:     RouterDefaultCategoryRule,
		CategoryPageRule: RouterDefaultCategoryPageRule,
		TagRule:          RouterDefaultTagRule,
		TagPageRule:      RouterDefaultTagPageRule,
		Options: fiber.Config{
			ReadBufferSize: 10240,
		},
		CompressLevel: 0,
		MinifyCode:    true,
		ETag:          true,
		ProxyHeader:   []string{},
		PprofSecret:   random.RandString(10),
	}
}

func (*Router) ConfigID() string {
	return "router"
}

func (r *Router) GetOptions() fiber.Config {
	r.Options.AppName = strutil.UpperFirst(constant.AppName) + " v" + constant.AppVersion
	r.Options.UnescapePath = true // 必须开启, 中文等特殊字符的param才能自动转码，否则要自己 url.QueryUnescape()
	r.Options.Prefork = false
	r.Options.GETOnly = false
	// 使用sonic或go-json等第三方解析会加大内存消耗
	// 而json解析多用在后台api接口，性能要求不高，自带的json解析足矣
	//r.Options.JSONEncoder = json.Marshal
	//r.Options.JSONDecoder = json.Unmarshal
	return r.Options
}

func (r *Router) UpdateAdminPath(path string) error {
	if path == "" {
		return errors.New("path is required")
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	r.AdminPath = path
	return nil
}

func (r *Router) GetAdminPath() (dir string) {
	return r.formatPath(r.AdminPath, constant.DefaultAdminPath)
}

func (r *Router) GetSitemapPath() (dir string) {
	return r.formatPath(r.SitemapPath, RouterDefaultSitemapPath)
}

func (r *Router) formatPath(path, def string) string {
	if path == "" || path == "/" {
		path = def
	}
	// add prefix
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	// remove suffix
	if strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}
	return path
}

func (r *Router) GetArticleRule() string {
	return r.formatRule(r.ArticleRule, RouterDefaultArticleRule)
}
func (r *Router) GetCategoryRule() string {
	return r.formatRule(r.CategoryRule, RouterDefaultCategoryRule)
}
func (r *Router) GetCategoryPageRule() string {
	return r.formatPageRule(r.CategoryPageRule, RouterDefaultCategoryPageRule)
}
func (r *Router) GetTagRule() string {
	return r.formatRule(r.TagRule, RouterDefaultTagRule)
}
func (r *Router) GetTagPageRule() string {
	return r.formatPageRule(r.TagPageRule, RouterDefaultTagPageRule)
}

func (r *Router) formatRule(rule, def string) string {
	if !strings.HasPrefix(rule, "/") {
		rule = "/" + rule
	}
	if !strings.Contains(rule, "{slug}") {
		rule = def
	}
	return strings.ReplaceAll(rule, "{slug}", ":slug")
}

func (r *Router) formatPageRule(rule, def string) string {
	if !strings.HasPrefix(rule, "/") {
		rule = "/" + rule
	}
	if !strings.Contains(rule, "{slug}") || !strings.Contains(rule, "{page}") {
		rule = def
	}
	rule = strings.ReplaceAll(rule, "{slug}", ":slug")
	rule = strings.ReplaceAll(rule, "{page}", ":page<int>")
	return rule
}
