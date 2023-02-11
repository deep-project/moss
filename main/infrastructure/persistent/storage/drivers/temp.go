package drivers

import (
	"moss/infrastructure/persistent/storage"
)

type Temp struct {
}

func (t *Temp) Init() error {
	return nil
}

func (t *Temp) Close() error {
	return nil
}

func (t *Temp) Get(key string) (*storage.GetValue, error) {
	return nil, nil
}

func (t *Temp) Set(key string, val *storage.SetValue) error {
	return nil
}

func (t *Temp) Delete(key string) error {
	return nil
}
