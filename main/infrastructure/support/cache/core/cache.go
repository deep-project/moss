package core

import (
	"errors"
	"moss/infrastructure/general/constant"
	"moss/infrastructure/support/cache/drivers"
	"moss/infrastructure/utils/timex"
	"time"
)

type Cache interface {
	Init() error
	Close() error
	Get(bucket, key string) ([]byte, error)
	Set(bucket, key string, val []byte, ttl time.Duration) error
	Delete(bucket, key string) error
	ClearBucket(bucket string) error
	Size() (int64, error)
}

const (
	BadgerDriverName    = "badger"
	RedisDriverName     = "redis"
	MemcachedDriverName = "memcached"
)

type Driver struct {
	Badger    *drivers.Badger    `json:"badger"`
	Redis     *drivers.Redis     `json:"redis"`
	Memcached *drivers.Memcached `json:"memcached"`
}

func NewDriver() *Driver {
	return &Driver{
		Badger: &drivers.Badger{
			Path:                    constant.CacheDir,
			MemTableSize:            128,
			BaseTableSize:           4,
			ValueLogFileSize:        512,
			NumMemtables:            5,
			NumLevelZeroTables:      5,
			NumLevelZeroTablesStall: 10,
			NumCompactors:           4,
			GcInterval:              timex.Duration{Number: 1, Unit: timex.DurationHour},
			GcDiscardRatio:          0.5,
		},
		Redis:     &drivers.Redis{Addr: "127.0.0.1:6379"},
		Memcached: &drivers.Memcached{Addr: "127.0.0.1:11211"},
	}
}

func (d *Driver) Items() []Cache {
	return []Cache{
		d.Badger,
		d.Redis,
		d.Memcached,
	}
}

func (d *Driver) Get(id string) (Cache, error) {
	if id == "" {
		return nil, errors.New("driver undefined")
	}
	switch id {
	case BadgerDriverName:
		return d.Badger, nil
	case RedisDriverName:
		return d.Redis, nil
	case MemcachedDriverName:
		return d.Memcached, nil
	}
	return nil, errors.New("driver not found")
}
