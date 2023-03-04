package upload

import (
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/random"
	"github.com/sony/sonyflake"
	"moss/domain/config"
	"moss/infrastructure/persistent/storage"
	"moss/infrastructure/support/log"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func init() {
	if err := Init(); err != nil {
		log.Error("init upload error", log.Err(err))
	}
}

func Init() error {
	config.Config.Upload.Storage.CloseAll() // 先全部关闭
	d, err := config.Config.Upload.Storage.ActiveDriver()
	if err != nil {
		return err
	}
	return d.Init()
}

func Upload(key, defaultExt string, val *storage.SetValue) (res *Result, err error) {
	res = NewResult(key, defaultExt)
	store, err := config.Config.Upload.Storage.ActiveDriver()
	if err != nil {
		return
	}
	err = store.Set(res.FullPath, val)
	return
}

type Result struct {
	RawKey   string
	RawDir   string
	RawName  string
	Ext      string
	Dir      string
	Name     string
	FullPath string
	URL      string
}

func NewResult(key, defaultExt string) *Result {
	ext := path.Ext(key)
	r := &Result{
		RawKey:  key,
		RawDir:  path.Dir(key),
		RawName: strings.TrimSuffix(path.Base(key), ext),
		Ext:     ext,
	}
	if r.RawName == "" {
		r.RawName = r.snowflake()
	}
	if r.Ext == "" {
		r.Ext = defaultExt
	}
	if !strings.HasPrefix(r.Ext, ".") {
		r.Ext = "." + r.Ext
	}
	// dir
	switch config.Config.Upload.PathFormat {
	case "date":
		r.Dir = time.Now().Format("20060102")
	case "hashDate":
		r.Dir = r.md5(time.Now().Format("20060102"))
	case "hashName":
		s := r.md5(r.RawName)
		r.Dir = filepath.Join(s[0:2], s[2:4], s[4:6])
	default:
		r.Dir = r.RawDir
	}
	// name
	switch config.Config.Upload.NameFormat {
	case "md5":
		r.Name = r.md5(r.RawName)
	case "uuid":
		r.Name = r.uuid()
	case "snowflake":
		r.Name = r.snowflake()
	default:
		r.Name = r.RawName
	}
	// fullPath
	r.FullPath = filepath.Join(r.Dir, r.Name+r.Ext)
	r.FullPath = filepath.ToSlash(r.FullPath)
	r.FullPath = strings.TrimPrefix(r.FullPath, "/") // 去掉开头的斜杠，防止对象存储无法识别目录，而存成一个整体的Key
	// url
	r.URL = config.Config.Upload.GetDomain() + r.FullPath
	return r
}

func (r *Result) md5(val string) string {
	return cryptor.Md5String(val)
}

func (r *Result) uuid() string {
	uuid, _ := random.UUIdV4()
	return uuid
}

var snowflakeHandle = sonyflake.NewSonyflake(sonyflake.Settings{MachineID: func() (uint16, error) {
	return 1, nil
}})

func (r *Result) snowflake() string {
	id, _ := snowflakeHandle.NextID()
	return strconv.FormatInt(int64(id), 10)
}
