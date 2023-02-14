package drivers

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"strconv"
	"strings"
	"time"
)

type Redis struct {
	Addr     string          `json:"addr"`
	Password string          `json:"password"`
	DB       int             `json:"db"`
	Handle   *redis.Client   `json:"-"`
	Ctx      context.Context `json:"-"`
}

func (r *Redis) Init() error {
	_ = r.Close()
	r.Handle = redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB,
	})
	r.Ctx = context.Background()
	_, err := r.Handle.Ping(context.Background()).Result()
	return err
}

func (r *Redis) Close() error {
	if err := r.undefined(); err != nil {
		return err
	}
	return r.Handle.Close()
}

func (r *Redis) Get(bucket, key string) ([]byte, error) {
	if err := r.undefined(); err != nil {
		return nil, err
	}
	val, err := r.Handle.Get(r.Ctx, r.prefix(bucket)+key).Bytes()
	return val, err
}

func (r *Redis) Set(bucket, key string, val []byte, ttl time.Duration) (err error) {
	if err := r.undefined(); err != nil {
		return err
	}
	return r.Handle.Set(r.Ctx, r.prefix(bucket)+key, val, ttl).Err()
}

func (r *Redis) Delete(bucket, key string) error {
	if err := r.undefined(); err != nil {
		return err
	}
	return r.Handle.Del(r.Ctx, r.prefix(bucket)+key).Err()
}

func (r *Redis) ClearBucket(bucket string) error {
	if err := r.undefined(); err != nil {
		return err
	}
	iter := r.Handle.Scan(r.Ctx, 0, r.prefix(bucket)+"*", 0).Iterator()
	for iter.Next(r.Ctx) {
		if err := r.Handle.Del(r.Ctx, iter.Val()).Err(); err != nil {
			return err
		}
	}
	return iter.Err()
}

func (r *Redis) prefix(bucket string) string {
	return bucket + ":"
}

func (r *Redis) undefined() error {
	if r.Handle == nil {
		return errors.New("client uninitialized or is closed")
	}
	return nil
}

func (r *Redis) Size() (_ int64, err error) {
	info, err := r.Handle.Info(context.Background(), "Memory").Result()
	if err != nil {
		return
	}
	arr := strings.Split(info, "\r\n")
	var rss string
	for _, v := range arr {
		if strings.HasPrefix(v, "used_memory_rss") {
			rss = strings.TrimPrefix(v, "used_memory_rss:")
			break
		}
	}
	return strconv.ParseInt(rss, 10, 0)
}
