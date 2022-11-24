package markdown

import (
	"bytes"
	"github.com/dstotijn/go-notion"
)

type Code struct {
	Code     string
	Language string
}

func (p Page) newCode(block notion.CodeBlock) Code {
	s := ""

	for _, text := range block.RichText {
		s += text.PlainText
	}

	language := ""
	if block.Language != nil {
		language = *block.Language
	}

	return Code{
		Code:     s,
		Language: language,
	}
}

func (p *Page) AddCodeToPage(block *notion.CodeBlock) error {
	md := p.newCode(*block)

	template, err := p.FetchTemplate("CodeTemplate.md")
	if err != nil {
		return err
	}

	var mdBuffer bytes.Buffer
	err = template.Execute(&mdBuffer, md)
	if err != nil {
		return err
	}

	p.AddBlock(mdBuffer.String())

	return nil
}
