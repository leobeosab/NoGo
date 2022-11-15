package markdown

import (
	"strings"
)

type Page struct {
	sbContent strings.Builder
	Title string
	tags []string
	categories []string
}

func NewPage(title string) *Page {
	return &Page {
		sbContent: strings.Builder{},
		Title: title,
	}
}

