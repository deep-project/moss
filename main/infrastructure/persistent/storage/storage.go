package storage

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

type Storage interface {
	Init() error
	Close() error
	Get(key string) (*GetValue, error)
	Set(key string, val *SetValue) error
	Delete(key string) error
}

type SetValue struct {
	Reader      io.ReadSeeker
	ContentType string
}

func NewSetValue(r io.ReadSeeker) *SetValue {
	return &SetValue{Reader: r}
}

func NewSetValueBytes(val []byte) *SetValue {
	return &SetValue{Reader: bytes.NewReader(val)}
}

func NewSetValueString(val string) *SetValue {
	return &SetValue{Reader: strings.NewReader(val)}
}

func NewSetValueStruct(val any) (*SetValue, error) {
	b, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	return NewSetValueBytes(b), nil
}

func (v *SetValue) Bytes() (b []byte, err error) {
	return io.ReadAll(v.Reader)
}

func (v *SetValue) String() (s string, err error) {
	b, err := v.Bytes()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

type GetValue struct {
	Reader io.ReadCloser
}

func NewGetValue(r io.ReadCloser) *GetValue {
	return &GetValue{Reader: r}
}

func NewGetValueBytes(val []byte) *GetValue {
	return &GetValue{Reader: io.NopCloser(bytes.NewBuffer(val))}
}

func NewGetValueNopCloser(r io.Reader) *GetValue {
	return &GetValue{Reader: io.NopCloser(r)}
}

func (v *GetValue) Bytes() (b []byte, err error) {
	b, err = io.ReadAll(v.Reader)
	defer v.Reader.Close()
	return
}

func (v *GetValue) String() (_ string, err error) {
	b, err := v.Bytes()
	if err != nil {
		return
	}
	return string(b), nil
}

func (v *GetValue) Unmarshal(ptr any) (err error) {
	b, err := v.Bytes()
	if err != nil {
		return
	}
	return json.Unmarshal(b, ptr)
}
