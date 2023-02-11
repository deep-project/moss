package drivers

import (
	"context"
	"errors"
	"io"
	"moss/infrastructure/persistent/storage"

	"github.com/kurin/blazer/b2"
)

type B2 struct {
	KeyID        string     `json:"keyID"`
	AppKey       string     `json:"appKey"`
	Bucket       string     `json:"bucket"`
	Handle       *b2.Client `json:"-"`
	HandleBucket *b2.Bucket `json:"-"`
}

func (b *B2) Init() (err error) {
	if b.KeyID == "" {
		return errors.New("keyID undefined")
	}
	if b.AppKey == "" {
		return errors.New("appKey undefined")
	}
	if b.Bucket == "" {
		return errors.New("bucket undefined")
	}
	if b.Handle, err = b2.NewClient(context.Background(), b.KeyID, b.AppKey); err != nil {
		return
	}
	b.HandleBucket, err = b.Handle.Bucket(context.Background(), b.Bucket)
	return nil
}

func (b *B2) isInit() error {
	if b.Handle == nil {
		return errors.New("handle uninitialized")
	}
	if b.HandleBucket == nil {
		return errors.New("bucket uninitialized")
	}
	return nil
}

func (b *B2) Close() error {
	return nil
}

func (b *B2) Get(key string) (_ *storage.GetValue, err error) {
	if err = b.isInit(); err != nil {
		return
	}
	r := b.HandleBucket.Object(key).NewReader(context.Background())
	return storage.NewGetValue(r), nil
}

func (b *B2) Set(key string, val *storage.SetValue) (err error) {
	if err = b.isInit(); err != nil {
		return
	}
	obj := b.HandleBucket.Object(key)
	w := obj.NewWriter(context.Background(), b2.WithAttrsOption(&b2.Attrs{ContentType: val.ContentType}))
	defer w.Close()
	_, err = io.Copy(w, val.Reader)
	return
}

func (b *B2) Delete(key string) (err error) {
	if err = b.isInit(); err != nil {
		return
	}
	return b.HandleBucket.Object(key).Delete(context.Background())
}
