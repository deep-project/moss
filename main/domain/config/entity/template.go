package entity

type Template struct {
	Head             string             `json:"head"`
	Footer           string             `json:"footer"`
	RobotsTxt        string             `json:"robots_txt"`  // robots.txt
	AdsTxt           string             `json:"ads_txt"`     // ads.txt
	Logo             string             `json:"logo"`        // logo.png
	FaviconIco       string             `json:"favicon_ico"` // favicon.ico
	Carousel         []TemplateCarousel `json:"carousel"`
	EnableCarousel   bool               `json:"enable_carousel"`
	Menu             *TemplateMenu      `json:"menu"`
	IndexList        *TemplateList      `json:"index_list"`
	GlobalList       *TemplateList      `json:"global_list"`
	CategoryPageList *TemplateList      `json:"category_page_list"`
	TagPageList      *TemplateList      `json:"tag_page_list"`
	TagCloud         *TemplateTagCloud  `json:"tag_cloud"`
}

func (*Template) ConfigID() string {
	return "template"
}

type TemplateCarousel struct {
	Image string `json:"image"`
	Link  string `json:"link"`
	Title string `json:"title"`
}

type TemplateMenu struct {
	Select []int  `json:"select"`
	Limit  int    `json:"limit"`
	Order  string `json:"order"`
}

type TemplateTagCloud struct {
	Limit  int    `json:"limit"`
	Order  string `json:"order"`
	Select []int  `json:"select"`
}

type TemplateList struct {
	Limit        int    `json:"limit"`
	Order        string `json:"order"`
	CategoryIds  []int  `json:"category_ids"`
	MaxPage      int    `json:"max_page"`
	DisableCount bool   `json:"disable_count"`
}

func NewTemplate() *Template {
	return &Template{
		Carousel:         []TemplateCarousel{},
		RobotsTxt:        "User-agent: *\nDisallow:",
		Menu:             &TemplateMenu{Select: []int{}, Limit: 40},
		IndexList:        &TemplateList{Limit: 20, Order: "id desc", CategoryIds: []int{}},
		GlobalList:       &TemplateList{Limit: 20, Order: "views desc", CategoryIds: []int{}},
		CategoryPageList: &TemplateList{Limit: 20, Order: "id desc"},
		TagPageList:      &TemplateList{Limit: 20},
		TagCloud:         &TemplateTagCloud{Limit: 15, Order: "id asc", Select: []int{}},
		Footer:           `©2023 Powered by <a href="https://github.com/deep-project/moss"><strong>Moss</strong></a>`,
	}
}
