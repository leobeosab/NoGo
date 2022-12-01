package markdown

import (
	"bytes"
	"github.com/dstotijn/go-notion"
)

type Paragraph struct {
	Text string
}

func (p Page) newParagraph(block notion.ParagraphBlock) Paragraph {
	paragraphStr, err := p.RichTextArrToString(block.RichText)
	if err != nil {
		panic(err)
	}

	return Paragraph{
		Text: paragraphStr,
	}
}

func (p *Page) AddParagraphToPage(block *notion.ParagraphBlock) error {
	md := p.newParagraph(*block)

	t, err := p.FetchTemplate("ParagraphTemplate.md")
	if err != nil {
		return err
	}

	var buff bytes.Buffer
	err = t.Execute(&buff, md)
	if err != nil {
		return err
	}

	p.AddContent(buff.String())

	return nil
}
