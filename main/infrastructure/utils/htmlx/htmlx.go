package htmlx

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

// GetTextFromHTML 获取html的纯文本
func GetTextFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	return doc.Text()
}
