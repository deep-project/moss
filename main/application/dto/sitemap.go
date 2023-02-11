package dto

import (
	"encoding/xml"
	"time"
)

type SitemapXML struct {
	XMLName xml.Name         `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 urlset"`
	Mobile  string           `xml:"xmlns:mobile,attr"`
	URL     []*SitemapXmlURL `xml:"url"`
}

func NewSitemapXML() *SitemapXML {
	return &SitemapXML{Mobile: "http://www.google.com/schemas/sitemap-mobile/1.0"}
}

func (s *SitemapXML) String() (string, error) {
	output, err := xml.Marshal(s)
	if err != nil {
		return "", err
	}
	return xml.Header + string(output), nil
}

type SitemapXmlURL struct {
	Loc        string           `xml:"loc"`
	Lastmod    string           `xml:"lastmod"`
	Changefreq string           `xml:"changefreq"`
	Priority   float64          `xml:"priority"`
	Mobile     SitemapURLMobile `xml:"mobile:mobile"`
}

type SitemapURLMobile struct {
	Type string `xml:"type,attr"`
}

func NewSitemapURL(loc string, t int64, changefreq string, Priority float64) *SitemapXmlURL {
	return &SitemapXmlURL{
		Loc:        loc,
		Lastmod:    time.Unix(t, 0).Format(time.RFC3339),
		Changefreq: changefreq,
		Priority:   Priority,
		Mobile:     SitemapURLMobile{Type: "pc,mobile"},
	}
}
