package drivers

import (
	"errors"
	"github.com/dgraph-io/badger/v2"
	"github.com/dgraph-io/badger/v2/options"
	"moss/infrastructure/general/constant"
	"moss/infrastructure/utils"
	"moss/infrastructure/utils/osx"
	"moss/infrastructure/utils/timex"
	"sync"
	"time"
)

type Badger struct {
	Path string `json:"path"`

	// 值日志数据加载模式
	// 0:FileIO 从文件加载
	// 1:LoadToRAM 全部加载到内存
	// 2:MemoryMap 映射加载到内存 (默认)
	ValueLogLoadingMode options.FileLoadingMode `json:"valueLogLoadingMode"`
	TableLoadingMode    options.FileLoadingMode `json:"tableLoadingMode"` // LSM树的加载模式
	NumMemtables        int                     `json:"numMemtables"`     // 内存表数量
	MaxTableSize        int64                   `json:"maxTableSize"`     // 内存表大小（兆）
	ValueLogFileSize    int64                   `json:"valueLogFileSize"` // 日志文件大小（兆）
	NumCompactors       int                     `json:"numCompactors"`    // 压缩工数量
	Compression         options.CompressionType `json:"compression"`      // 压缩方式 0:none 1:snappy 2:zstd
	SyncWrites          bool                    `json:"syncWrites"`       // 同步写 关闭可以提高性能
	GcInterval          timex.Duration          `json:"gcInterval"`       // 垃圾回收时间间隔
	GcDiscardRatio      float64                 `json:"gcDiscardRatio"`   // 垃圾回收丢弃比例

	Handle *badger.DB `json:"-"`
	onceGC sync.Once
}

func NewBadger() *Badger {

	return &Badger{
		Path:                constant.CacheDir + "/badger",
		ValueLogLoadingMode: options.MemoryMap,
		TableLoadingMode:    options.MemoryMap,
		NumMemtables:        2,
		MaxTableSize:        16,
		ValueLogFileSize:    256, // 设置512在1G内存下会占用过高
		NumCompactors:       2,
		Compression:         1,
		SyncWrites:          false,
		GcInterval:          timex.Duration{Number: 5, Unit: timex.DurationMinute},
		GcDiscardRatio:      0.9,
	}
}

func (b *Badger) Init() (err error) {
	_ = b.Close()
	if b.Path == "" {
		return errors.New("path undefined")
	}

	if b.NumCompactors <= 1 {
		return errors.New("numCompactors must be > 1")
	}
	_ = osx.CreateDirIsNotExist(b.Path, 0755)
	var opt = badger.DefaultOptions(b.Path).
		WithTruncate(true).
		WithLogger(nil).
		WithLoadBloomsOnOpen(false). // 启动时延迟加载布隆过滤器
		WithZSTDCompressionLevel(7). // zstd压缩等级
		WithValueLogLoadingMode(b.ValueLogLoadingMode).
		WithNumMemtables(b.NumMemtables).
		WithNumLevelZeroTables(b.NumMemtables).
		WithNumLevelZeroTablesStall(b.NumMemtables * 2).
		WithMaxTableSize(b.MaxTableSize << 20).
		WithValueLogFileSize((b.ValueLogFileSize / 2) << 20).
		WithNumCompactors(b.NumCompactors).
		WithCompression(b.Compression).
		WithSyncWrites(b.SyncWrites)

	if b.Handle, err = badger.Open(opt); err != nil {
		return err
	}
	b.onceGC.Do(func() {
		go b.autoGC()
	})
	return
}

func (b *Badger) autoGC() {
	td := b.GcInterval.Duration()
	if td == 0 {
		b.GcInterval.Number = 5
		b.GcInterval.Unit = "minute"
		td = 5 * time.Minute
	}
	ticker := time.NewTicker(td)
	defer ticker.Stop()
	for range ticker.C {
	again:
		if err := b.RunValueLogGC(); err == nil {
			goto again
		}
	}
}

func (b *Badger) RunValueLogGC() error {
	if err := b.undefined(); err != nil {
		return err
	}
	if b.GcDiscardRatio <= 0 || b.GcDiscardRatio > 1 {
		b.GcDiscardRatio = 0.9
	}
	return b.Handle.RunValueLogGC(b.GcDiscardRatio)
}

func (b *Badger) Close() error {
	if err := b.undefined(); err != nil {
		return err
	}
	if err := b.Handle.Close(); err != nil {
		return err
	}
	b.Handle = nil
	return nil
}

func (b *Badger) Get(bucket, key string) (val []byte, err error) {
	if err := b.undefined(); err != nil {
		return nil, err
	}
	err = b.Handle.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(b.prefix(bucket) + key))
		if err != nil {
			return err
		}
		val, err = item.ValueCopy(nil)
		return err
	})
	return
}

func (b *Badger) Set(bucket, key string, val []byte, ttl time.Duration) error {
	if err := b.undefined(); err != nil {
		return err
	}
	return b.Handle.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(b.prefix(bucket)+key), val)
		if ttl > 0 {
			e.WithTTL(ttl)
		}
		return txn.SetEntry(e)
	})
}

func (b *Badger) Delete(bucket, key string) error {
	if err := b.undefined(); err != nil {
		return err
	}
	return b.Handle.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(b.prefix(bucket) + key))
	})
}

func (b *Badger) ClearBucket(bucket string) error {
	if err := b.undefined(); err != nil {
		return err
	}
	return b.Handle.DropPrefix([]byte(b.prefix(bucket)))
}

func (b *Badger) prefix(bucket string) string {
	return bucket + ":"
}

func (b *Badger) undefined() error {
	if b == nil || b.Handle == nil || b.Handle.IsClosed() {
		return errors.New("client uninitialized or is closed")
	}
	return nil
}

func (b *Badger) Size() (int64, error) {
	size, _ := utils.DirSize(b.Path)
	return int64(size), nil
}
