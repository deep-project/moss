package drivers

import (
	"errors"
	"moss/infrastructure/persistent/storage"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/jlaffaye/ftp"
)

type Ftp struct {
	Host     string          `json:"host"`
	Port     string          `json:"port"`
	Name     string          `json:"name"`
	Password string          `json:"password"`
	Timeout  int             `json:"timeout"`
	Handle   *ftp.ServerConn `json:"-"`
	lock     sync.Mutex
}

func (f *Ftp) Init() (err error) {
	_ = f.Close()
	if f.Host == "" {
		return errors.New("host undefined")
	}
	if f.Port == "" {
		return errors.New("port undefined")
	}
	if f.Name == "" {
		return errors.New("name undefined")
	}
	if f.Password == "" {
		return errors.New("password undefined")
	}
	if f.Handle, err = ftp.Dial(f.Host+":"+f.Port,
		ftp.DialWithTimeout(5*time.Second),
		ftp.DialWithShutTimeout(time.Duration(f.Timeout)*time.Second),
	); err != nil {
		return
	}
	return f.Handle.Login(f.Name, f.Password)
}

func (f *Ftp) Close() error {
	if f.Handle == nil {
		return errors.New("handle uninitialized")
	}
	return f.Handle.Quit()
}

func (f *Ftp) Get(key string) (*storage.GetValue, error) {
	if f.Handle == nil {
		return nil, errors.New("handle uninitialized")
	}
	if !f.isAlive() {
		_ = f.Init() // 重新连接
	}
	f.lock.Lock()
	defer f.lock.Unlock()
	key, _ = f.formatKey(key)
	_ = f.Handle.ChangeDir("/") // 必须切换到根目录
	r, err := f.Handle.Retr(key)
	if err != nil {
		return nil, err
	}
	return storage.NewGetValue(r), nil
}

func (f *Ftp) Set(key string, val *storage.SetValue) error {
	if f.Handle == nil {
		return errors.New("handle uninitialized")
	}
	f.lock.Lock()
	defer f.lock.Unlock()
	if !f.isAlive() {
		_ = f.Init() // 重新连接
	}
	key, dir := f.formatKey(key)
	// 测试是否能切换到目标目录，否则则创建目录
	if !strings.HasPrefix(dir, "/") {
		dir = "/" + dir
	}
	if f.Handle.ChangeDir(dir) != nil {
		if err := f.createDirs(dir); err != nil {
			return err
		}
	}

	_, filename := filepath.Split(key)
	return f.Handle.Stor(filename, val.Reader)
}

func (f *Ftp) Delete(key string) error {
	if f.Handle == nil {
		return errors.New("handle uninitialized")
	}
	f.lock.Lock()
	defer f.lock.Unlock()
	if !f.isAlive() {
		_ = f.Init() // 重新连接
	}
	key, _ = f.formatKey(key)
	_ = f.Handle.ChangeDir("/") // 必须切换到根目录
	return f.Handle.Delete(key)
}

// 循环创建ftp的文件夹
func (f *Ftp) createDirs(dir string) error {
	dirs := strings.Split(dir, string(filepath.Separator))
	current := "/"
	for _, v := range dirs {
		current = current + v + "/"
		if f.Handle.ChangeDir(current) != nil {
			err := f.Handle.MakeDir(current) // 创建目录
			if err != nil {
				return err
			}
			err = f.Handle.ChangeDir(current) // 再次切换进目录
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *Ftp) formatKey(key string) (string, string) {
	dir, _ := filepath.Split(key)
	key = filepath.ToSlash(key)
	dir = filepath.ToSlash(dir)
	return key, dir
}

func (f *Ftp) isAlive() bool {
	if f.Handle == nil {
		return false
	}
	if err := f.Handle.NoOp(); err != nil {
		return false
	}
	return true
}
