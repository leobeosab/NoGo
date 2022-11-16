package markdown

import (
	"bytes"
	"github.com/dstotijn/go-notion"
	"text/template"
)

type Paragraph struct {
	Text string
}

func newParagraph(block notion.ParagraphBlock) Paragraph {
	paragraphStr, err := RichTextArrToString(block.RichText)
	if err != nil {
		panic(err)
	}

	return Paragraph{
		Text: paragraphStr,
	}
}

func (p *Page) AddParagraphToPage(block *notion.ParagraphBlock) error {
	md := newParagraph(*block)

	t, err := template.ParseFiles("blocks/ParagraphTemplate.md")
	if err != nil {
		return err
	}

	var buff bytes.Buffer
	err = t.Execute(&buff, md)
	if err != nil {
		return err
	}

	p.AddBlock(buff.String())

	return nil
}
