package markdown

import (
	"bytes"
)

type HeaderMarkdown struct {
	Title         string
	Categories    []string
	Tags          []string
	Date          string
	AssetPath     string
	CoverFileName string
	CoverAlt      string
	Description   string
}

// TODO refactor what's in the Page struct
func (p *Page) SetHeader(categories []string, tags []string) error {

	header := HeaderMarkdown{
		Title:         p.Title,
		AssetPath:     p.AssetURL,
		Categories:    categories,
		Tags:          tags,
		Date:          p.Metadata.PublishDate,
		CoverFileName: p.Metadata.CoverFileName,
		Description:   p.Metadata.Description,
		CoverAlt:      p.Metadata.CoverAlt,
	}

	temp, err := p.FetchTemplate("HeaderTemplate.md")
	if err != nil {
		return err
	}

	var mdBuffer bytes.Buffer
	err = temp.Execute(&mdBuffer, header)
	if err != nil {
		return err
	}

	// Prepent Header
	currentBuffer := p.sbContent.String()
	p.sbContent.Reset()
	p.sbContent.WriteString(mdBuffer.String())
	p.sbContent.WriteString(currentBuffer)

	return nil
}
