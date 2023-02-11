package entity

type Sitemap struct {
	Article  *SitemapOption `json:"article"`
	Category *SitemapOption `json:"category"`
	Tag      *SitemapOption `json:"tag"`
}

type SitemapOption struct {
	Limit      int     `json:"limit"`
	ChangeFreq string  `json:"change_freq"`
	Priority   float64 `json:"priority"`
	InHours    int     `json:"in_hours"`
}

func (*Sitemap) ConfigID() string {
	return "sitemap"
}

func NewSitemap() *Sitemap {
	return &Sitemap{
		Article:  &SitemapOption{Limit: 1000, ChangeFreq: "monthly", Priority: 0.5},
		Category: &SitemapOption{Limit: 1000, ChangeFreq: "daily", Priority: 0.5},
		Tag:      &SitemapOption{Limit: 1000, ChangeFreq: "weekly", Priority: 0.5},
	}
}
