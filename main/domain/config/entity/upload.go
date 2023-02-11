package entity

import (
	"moss/infrastructure/general/constant"
	"moss/infrastructure/persistent/storage/facade"
	"net/url"
	"strings"
)

type Upload struct {
	Domain     string          `json:"domain"`
	PathFormat string          `json:"path_format"`
	NameFormat string          `json:"name_format"`
	Storage    *facade.Storage `json:"storage"`
}

func NewUpload() *Upload {
	store := facade.NewStorage()
	store.Driver.Local.Path = constant.UploadDir
	return &Upload{Domain: constant.UploadDomain, PathFormat: "date", NameFormat: "snowflake", Storage: store}
}

func (*Upload) ConfigID() string {
	return "upload"
}

// GetDomain domain 必须以斜杠结尾
func (u *Upload) GetDomain() string {
	if u.Domain == "" {
		return "/"
	}
	if !strings.HasSuffix(u.Domain, "/") {
		return u.Domain + "/"
	}
	return u.Domain
}

// ContainsDomain 判断url是否是当前的domain
func (u *Upload) ContainsDomain(uri string) bool {
	if !u.DomainIsUrl() {
		return false
	}
	p, err := url.Parse(u.Domain)
	if err != nil {
		return false
	}
	return strings.Contains(uri, p.Host)
}

// DomainIsUrl 判断domain是否是URL
func (u *Upload) DomainIsUrl() bool {
	return strings.HasPrefix(u.Domain, "http") || strings.HasPrefix(u.Domain, "//")
}
