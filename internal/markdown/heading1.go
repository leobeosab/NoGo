package markdown

import (
	"bytes"
	"fmt"
	"github.com/dstotijn/go-notion"
	"text/template"
)

type heading1markdown struct {
	Text string
}

func newHeading1Markdown(block notion.Heading1Block) heading1markdown {
	s, err := RichTextArrToString(block.RichText)
	if err != nil {
		panic(err)
	}

	return heading1markdown{
		Text: s,
	}
}

func (p Page) AddHeading1ToPage(block *notion.Heading1Block) error {
	md := newHeading1Markdown(*block)

	template, err := template.ParseFiles("blocks/Heading1Template.md")
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
