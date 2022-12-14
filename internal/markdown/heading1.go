package markdown

import (
	"bytes"
	"github.com/dstotijn/go-notion"
)

type heading1markdown struct {
	Text string
}

func (p PageBuilder) newHeading1Markdown(block notion.Heading1Block) heading1markdown {
	s, err := p.RichTextArrToString(block.RichText)
	if err != nil {
		panic(err)
	}

	return heading1markdown{
		Text: s,
	}
}

func (p *PageBuilder) AddHeading1ToPage(block *notion.Heading1Block) error {
	md := p.newHeading1Markdown(*block)

	template, err := p.FetchTemplate("Heading1Template.md")
	if err != nil {
		return err
	}

	var mdBuffer bytes.Buffer
	err = template.Execute(&mdBuffer, md)
	if err != nil {
		return err
	}

	p.AddContent(mdBuffer.String())

	return nil
}
