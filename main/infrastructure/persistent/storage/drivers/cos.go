package drivers

import (
	"context"
	"errors"
	"github.com/tencentyun/cos-go-sdk-v5"
	"moss/infrastructure/persistent/storage"
	"net/http"
	"net/url"
	"strings"
)

type Cos struct {
	BucketURL string      `json:"bucketURL"`
	SecretID  string      `json:"secretID"`
	SecretKey string      `json:"secretKey"`
	Handle    *cos.Client `json:"-"`
}

func (c *Cos) Init() error {
	if c.BucketURL == "" {
		return errors.New("region undefined")
	}
	if c.SecretID == "" {
		return errors.New("secretID undefined")
	}
	if c.SecretKey == "" {
		return errors.New("secretKey undefined")
	}
	u, err := url.Parse(c.BucketURL)
	if err != nil {
		return err
	}
	c.Handle = cos.NewClient(&cos.BaseURL{BucketURL: u}, &http.Client{
		Transport: &cos.AuthorizationTransport{SecretID: c.SecretID, SecretKey: c.SecretKey},
	})
	return nil
}

func (c *Cos) Close() error {
	return nil
}

func (c *Cos) Get(key string) (*storage.GetValue, error) {
	if c.Handle == nil {
		return nil, errors.New("handle undefined")
	}
	resp, err := c.Handle.Object.Get(context.Background(), key, nil)
	return storage.NewGetValue(resp.Body), err
}

func (c *Cos) Set(key string, val *storage.SetValue) error {
	if c.Handle == nil {
		return errors.New("handle undefined")
	}
	_, err := c.Handle.Object.Put(context.Background(), key, val.Reader,
		&cos.ObjectPutOptions{ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{ContentType: val.ContentType}})
	return err
}

func (c *Cos) Delete(key string) error {
	if c.Handle == nil {
		return errors.New("handle undefined")
	}
	_, err := c.Handle.Object.Delete(context.Background(), key, nil)
	return err
}

// DeleteByPrefix from https://cloud.tencent.com/document/product/436/65648
func (c *Cos) DeleteByPrefix(prefix string) error {

	if !strings.HasSuffix(prefix, "/") {
		prefix = prefix + "/"
	}
	var marker string
	opt := &cos.BucketGetOptions{
		Prefix:  prefix,
		MaxKeys: 1000,
	}
	isTruncated := true
	for isTruncated {
		opt.Marker = marker
		v, _, err := c.Handle.Bucket.Get(context.Background(), opt)
		if err != nil {
			return err
		}
		for _, content := range v.Contents {
			_, err = c.Handle.Object.Delete(context.Background(), content.Key)
			if err != nil {
				return err
			}
		}
		isTruncated = v.IsTruncated
		marker = v.NextMarker
	}
	return nil
}
