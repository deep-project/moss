package entity

import (
	"github.com/duke-git/lancet/v2/random"
	"strings"
)

type Site struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Title       string `json:"title"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
	Salt        string `json:"salt"`
}

func NewSite() *Site {
	return &Site{Name: "moss", Title: "Hello Moss", Salt: random.RandString(10)}
}

func (*Site) ConfigID() string {
	return "site"
}

func (s *Site) GetURL() (url string) {
	url = s.URL
	if url == "" {
		return
	}
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "//") {
		url = "//" + url
	}
	if strings.HasSuffix(url, "/") {
		url = strings.TrimSuffix(url, "/")
	}
	return
}
