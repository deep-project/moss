package drivers

import (
	"errors"
	"github.com/dgraph-io/badger/v3"
	"moss/infrastructure/utils"
	"moss/infrastructure/utils/timex"
	"sync"
	"time"
)

type Badger struct {
	Path                    string         `json:"path"`
	MemTableSize            int64          `json:"memTableSize"`            // 内存表大小（兆）
	BaseTableSize           int64          `json:"baseTableSize"`           // 数据表大小（兆）
	ValueLogFileSize        int64          `json:"valueLogFileSize"`        // 日志文件大小（兆）
	NumMemtables            int            `json:"numMemtables"`            // 内存表数量
	NumLevelZeroTables      int            `json:"numLevelZeroTables"`      // 零级表数量
	NumLevelZeroTablesStall int            `json:"numLevelZeroTablesStall"` // 零级表停滞数量
	NumCompactors           int            `json:"numCompactors"`           // 压缩工数量
	GcInterval              timex.Duration `json:"gcInterval"`              // 垃圾回收时间间隔
	GcDiscardRatio          float64        `json:"gcDiscardRatio"`          // 垃圾回收丢弃比例
	Handle                  *badger.DB     `json:"-"`
	onceGC                  sync.Once
}

func (b *Badger) Init() (err error) {
	_ = b.Close()
	if b.Path == "" {
		return errors.New("path undefined")
	}
	var opt = badger.DefaultOptions(b.Path)

	opt.MemTableSize = (b.MemTableSize / 2) << 20
	opt.BaseTableSize = (b.BaseTableSize / 2) << 20
	opt.ValueLogFileSize = (b.ValueLogFileSize / 2) << 20

	opt.NumMemtables = b.NumMemtables
	opt.NumLevelZeroTables = b.NumLevelZeroTables
	opt.NumLevelZeroTablesStall = b.NumLevelZeroTablesStall
	opt.NumCompactors = b.NumCompactors

	opt.Logger = nil // 不打印日志
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
		return
	}
	ratio := b.GcDiscardRatio
	if ratio <= 0 || ratio > 1 {
		ratio = 0.5
	}
	ticker := time.NewTicker(td)
	defer ticker.Stop()
	for range ticker.C {
	again:
		if b.Handle == nil || b.Handle.IsClosed() {
			return
		}
		err := b.Handle.RunValueLogGC(ratio)
		if err == nil {
			goto again
		} else {
			return
		}
	}
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
	if b.Handle == nil || b.Handle.IsClosed() {
		return errors.New("client uninitialized or is closed")
	}
	return nil
}

func (b *Badger) Size() (int64, error) {
	size, err := utils.DirSize(b.Path)
	return int64(size), err
}
