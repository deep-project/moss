package drivers

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"moss/infrastructure/persistent/storage"
	"strings"
)

type Oss struct {
	Endpoint        string      `json:"endpoint"`
	AccessKeyID     string      `json:"accessKeyID"`
	AccessKeySecret string      `json:"accessKeySecret"`
	Bucket          string      `json:"bucket"`
	Client          *oss.Client `json:"-"`
	ClientBucket    *oss.Bucket `json:"-"`
}

func (o *Oss) Init() (err error) {
	if o.Endpoint == "" {
		return errors.New("endpoint is undefined")
	}
	if o.AccessKeyID == "" {
		return errors.New("accessKeyID is undefined")
	}
	if o.AccessKeySecret == "" {
		return errors.New("accessKeySecret is undefined")
	}
	if o.Bucket == "" {
		return errors.New("bucket is undefined")
	}
	if o.Client, err = oss.New(o.Endpoint, o.AccessKeyID, o.AccessKeySecret); err != nil {
		return
	}
	o.ClientBucket, err = o.Client.Bucket(o.Bucket)
	return
}

func (o *Oss) isInit() error {
	if o.Client == nil {
		return errors.New("client uninitialized")
	}
	if o.Client == nil {
		return errors.New("bucket uninitialized")
	}
	return nil
}

func (o *Oss) Close() error {
	return nil
}

func (o *Oss) Get(key string) (*storage.GetValue, error) {
	if err := o.isInit(); err != nil {
		return nil, err
	}
	r, err := o.ClientBucket.GetObject(key)
	return storage.NewGetValue(r), err
}

func (o *Oss) Set(key string, val *storage.SetValue) (err error) {
	if err = o.isInit(); err != nil {
		return
	}
	return o.ClientBucket.PutObject(key, val.Reader, oss.ContentType(val.ContentType))
}

func (o *Oss) Delete(key string) error {
	if err := o.isInit(); err != nil {
		return err
	}
	return o.ClientBucket.DeleteObject(key)
}

// DeleteByPrefix from https://help.aliyun.com/document_detail/88644.html
func (o *Oss) DeleteByPrefix(p string) error {

	// 如果您需要删除所有前缀为src的文件，则prefix设置为src。设置为src后，所有前缀为src的非目录文件、src目录以及目录下的所有文件均会被删除。
	// 如果您仅需要删除src目录及目录下的所有文件，则prefix设置为src/。
	if !strings.HasSuffix(p, "/") {
		p = p + "/"
	}
	marker := oss.Marker("")
	prefix := oss.Prefix(p)
	for {
		lor, err := o.ClientBucket.ListObjects(marker, prefix)
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		objects := []string{}
		for _, object := range lor.Objects {
			objects = append(objects, object.Key)
		}
		// 将oss.DeleteObjectsQuiet设置为true，表示不返回删除结果。
		_, err = o.ClientBucket.DeleteObjects(objects, oss.DeleteObjectsQuiet(true))
		if err != nil {
			return err
		}
		prefix = oss.Prefix(lor.Prefix)
		marker = oss.Marker(lor.NextMarker)
		if !lor.IsTruncated {
			break
		}
	}
	return nil
}
