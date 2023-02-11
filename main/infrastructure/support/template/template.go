package template

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/fsnotify/fsnotify"
	"github.com/valyala/bytebufferpool"
	"io"
	"io/fs"
	"moss/domain/config"
	"moss/infrastructure/general/constant"
	"moss/infrastructure/support/log"
	"moss/infrastructure/support/template/engine"
	"moss/infrastructure/utils/osx"
	"moss/resources"
	"os"
	"path/filepath"
	"strings"
)

var Client = new(Template)

func init() {
	if !fileutil.IsExist(constant.ThemesDir) {
		if err := InitThemeDir(); err != nil {
			log.Error("init dir failed", log.Err(err))
		}
	}
	if err := Client.InitWatcher(); err != nil {
		log.Error("init watcher failed", log.Err(err))
	}
	if err := InitTemplate(); err != nil {
		log.Error("init template failed", log.Err(err))
	}
}

func InitThemeDir() error {
	return osx.EmbedCopyToDir(resources.Themes, resources.ThemesDirName, constant.ThemesDir)
}

func InitTemplate() (err error) {
	themePath, err := CurrentThemePath()
	if err == nil {
		Client.InitEngine(newEngine(themePath, []string{"page", "template"}))
	}
	subPaths := []string{filepath.Join(themePath, "page"), filepath.Join(themePath, "template")}
	return Client.watcher.reset(subPaths...)
}

func CurrentThemePath() (string, error) {
	if config.Config.Theme.Current == "" {
		return "", errors.New("theme is undefined")
	}
	return filepath.Join(constant.ThemesDir, config.Config.Theme.Current), nil
}

func Render(template string, binds Binds) ([]byte, error) {
	return Client.Render(template, binds)
}

type Template struct {
	engine  Engine
	watcher watcher
}

func (t *Template) InitEngine(e Engine) {
	t.engine = e
	initEngineFn(t.engine)
}

func (t *Template) InitWatcher() (err error) {
	return t.watcher.Init(func(e fsnotify.Event) {
		if err := t.Reload(); err != nil {
			log.Warn("init watcher failed", log.Err(err))
		}
	})
}

func (t *Template) Render(template string, binds Binds) ([]byte, error) {
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	err := t.engine.Render(buf, template, binds.make())
	return buf.Bytes(), err
}

func (t *Template) Reload() error {
	if t.engine == nil {
		return errors.New("engine is undefined")
	}
	t.engine.Reload()
	return nil
}

type Binds struct {
	Page Page
	Data any
}

type Page struct {
	Name        string
	Path        string
	Title       string
	Description string
	Keywords    string
	PageNumber  int
}

func (b *Binds) make() map[string]any {
	return map[string]any{
		"Page": b.Page,
		"Data": b.Data,
	}
}

type Engine interface {
	Render(out io.Writer, template string, binds map[string]any, layout ...string) error
	AddFunc(name string, fn any)
	Reload()
}

func newEngine(themePath string, allowedDir []string) Engine {
	return engine.NewJet(themePath, skipLoad(allowedDir))
}

// skip some dir do not load
func skipLoad(allowedDir []string) func(string) bool {
	return func(path string) bool {
		path = filepath.ToSlash(path)
		if dir, _ := filepath.Split(path); dir == "" {
			return true // path is a root file
		}
		for _, val := range allowedDir {
			if strings.HasPrefix(path, filepath.ToSlash(val)) {
				return false
			}
		}
		return true
	}
}

type watcher struct {
	client *fsnotify.Watcher
}

func newWatcher() *watcher {
	return &watcher{}
}

func (w *watcher) Init(fn func(fsnotify.Event)) (err error) {
	w.client, err = osx.Watch([]string{}, fn)
	return
}

func (w *watcher) reset(paths ...string) error {
	if w.client == nil {
		return errors.New("watcher client is undefined")
	}
	if err := w.clear(); err != nil {
		return err
	}
	return w.add(paths...)
}

func (w *watcher) clear() (err error) {
	for _, d := range w.client.WatchList() {
		if err = w.client.Remove(d); err != nil {
			return
		}
	}
	return
}

// 由于fsnotify.Watcher不能监控到超过两层的目录
// 所以遍历所有目录层级。全部加入到watcher监控列表
func (w *watcher) add(paths ...string) (err error) {
	for _, path := range paths {
		if err = filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
			if d == nil || !d.IsDir() {
				return nil
			}
			return w.client.Add(p)
		}); err != nil {
			return
		}
	}
	return
}

func ThemeList() (list []Theme, err error) {
	directories, err := os.ReadDir(constant.ThemesDir)
	if err != nil {
		return
	}
	for _, dir := range directories {
		if dir.IsDir() {
			list = append(list, newTheme(dir.Name()))
		}
	}
	return
}

type Theme struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Author        string `json:"author"`
	Version       string `json:"version"`
	Description   string `json:"description"`
	Homepage      string `json:"homepage"`
	License       string `json:"license"`
	HasScreenshot bool   `json:"has_screenshot"`
}

func newTheme(dir string) Theme {
	t := Theme{ID: dir}
	t.HasScreenshot = osx.IsExist(filepath.Join(constant.ThemesDir, dir, "screenshot.png"))
	if b, err := os.ReadFile(filepath.Join(constant.ThemesDir, dir, "theme.json")); err == nil {
		_ = json.Unmarshal(b, &t)
	}
	return t
}

func ReadThemeScreenshot(dir string) string {
	b, err := os.ReadFile(filepath.Join(constant.ThemesDir, dir, "screenshot.png"))
	if err != nil {
		return ""
	}
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(b)
}
