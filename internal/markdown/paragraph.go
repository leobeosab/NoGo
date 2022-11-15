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
	paragraphStr, err := RichTextArrToString(block.RichText)
	if err != nil {
		panic(err)
	}

	return Paragraph{
		Text: paragraphStr,
	}
}

func (p Page) AddParagraphToPage(block *notion.ParagraphBlock) error {
	md := newParagraph(*block)

	t, err := template.ParseFiles("blocks/ParagraphTemplate.md")
	if err != nil {
		return err
	}

	var mdBuffer bytes.Buffer
	err = t.Execute(&mdBuffer, md)
	if err != nil {
		return err
	}

	p.sbContent.WriteString(mdBuffer.String())
	fmt.Println(mdBuffer.String())

	return nil
}
