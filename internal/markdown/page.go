package markdown

import (
	"strings"
)

type Page struct {
	sbContent  strings.Builder
	Title      string
	ID         string
	tags       []string
	categories []string
	Assets     []PageAsset
}

type PageAsset struct {
	ContentURL   string
	DownloadPath string
}

func NewPage(title string) *Page {
	id := strings.Replace(strings.ToLower(title), " ", "-", -1)

	return &Page{
		sbContent: strings.Builder{},
		Title:     title,
		ID:        id,
		Assets:    make([]PageAsset, 0),
	}
}

func (p *Page) AddAsset(contentURL string, downloadPath string) {
	p.Assets = append(p.Assets, PageAsset{
		ContentURL:   contentURL,
		DownloadPath: downloadPath,
	})
}

func (p *Page) AddBlock(block string) {
	p.sbContent.WriteString("\n" + block)
}

func (p *Page) Build() string {
	return p.sbContent.String()
}
