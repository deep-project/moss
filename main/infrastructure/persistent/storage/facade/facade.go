package facade

import (
	"errors"
	"moss/infrastructure/persistent/storage"
	"moss/infrastructure/persistent/storage/drivers"
)

type Storage struct {
	Active string  `json:"active"`
	Driver *Driver `json:"driver"`
}

func NewStorage() *Storage {
	return &Storage{
		Active: "local",
		Driver: NewDriver(),
	}
}

func (s *Storage) ActiveDriver() (storage.Storage, error) {
	return s.Driver.Get(s.Active)
}

func (s *Storage) CloseAll() {
	for _, item := range s.Driver.Items() {
		_ = item.Close()
	}
}

type Driver struct {
	Local *drivers.Local `json:"local"`
	Ftp   *drivers.Ftp   `json:"ftp"`
	B2    *drivers.B2    `json:"b2"`
	Cos   *drivers.Cos   `json:"cos"`
	Oss   *drivers.Oss   `json:"oss"`
	S3    *drivers.S3    `json:"s3"`
}

func NewDriver() *Driver {
	return &Driver{
		Local: &drivers.Local{},
		Ftp:   &drivers.Ftp{Port: "21"},
		B2:    &drivers.B2{},
		Cos:   &drivers.Cos{},
		Oss:   &drivers.Oss{},
		S3:    &drivers.S3{},
	}
}

func (d *Driver) Items() []storage.Storage {
	return []storage.Storage{
		d.Local,
		d.Ftp,
		d.B2,
		d.Cos,
		d.Oss,
		d.S3,
	}
}

func (d *Driver) Get(id string) (storage.Storage, error) {
	if id == "" {
		return nil, errors.New("id undefined")
	}
	switch id {
	case "local":
		return d.Local, nil
	case "ftp":
		return d.Ftp, nil
	case "b2":
		return d.B2, nil
	case "cos":
		return d.Cos, nil
	case "oss":
		return d.Oss, nil
	case "s3":
		return d.S3, nil
	}
	return nil, errors.New("driver not found")
}
