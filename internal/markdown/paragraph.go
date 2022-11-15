package markdown

import (
	"bytes"
	"fmt"
	"github.com/dstotijn/go-notion"
	"text/template"
)

type Paragraph struct {
	Text string
}

func newParagraph(block notion.ParagraphBlock) Paragraph {
	s := ""

	for _, text := range block.RichText {
		s += text.PlainText
	}

	return Paragraph{
		Text: s,
	}
}

func (p Page) AddParagraphToPage(block *notion.ParagraphBlock) error {
	md := newParagraph(*block)

	template, err := template.ParseFiles("blocks/ParagraphTemplate.md")
	if err != nil {
		return err
	}

	var mdBuffer bytes.Buffer
	err = template.Execute(&mdBuffer, md)
	if err != nil {
		return err
	}

	p.sbContent.WriteString(mdBuffer.String())
	fmt.Println(mdBuffer.String())

	return nil
}
